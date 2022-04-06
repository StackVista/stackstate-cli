# My adventure with codesigning

## Windows machine for testing

I spun up a Window 2019 server on AWS Ec2. 

* Login with RDP:
Install Microsoft Remote Desktop: https://apps.apple.com/us/app/microsoft-remote-desktop/id1295203466?mt=12
Open: `windows-signing/CLI-windows-test-machine.rdp`

## Creating a self-signed certificate

Followed the steps in: 

https://docs.microsoft.com/en-us/windows/msix/package/create-certificate-package-signing

Somehow I got an error at some point that it couldn't find the certificate with that thumbnail. I don't know what the problem was, but I had some help from https://stackoverflow.com/a/44766033/1860591 

I used `SuperS3cret!` as a password. 

This then gave me the file PKCS12 file `test.pfx`, which I send to myself via wetransfer.com.

## Building osslsigncode

I followed the Mac OSx instructions for building a version of osslsigncode:

https://github.com/mtrojnar/osslsigncode

It worked without hiccups. I have included the binary that I've built on my x86 Mac under `windows-signing/osslsigncode`.

## Signing a windows binary

1. I create a binary by running `gorelease release --rm-dist --skip-validate`
2. To run `osslsigncode` I needed a key and certificate in PEM format. I ran the commands found on:

https://trustzone.com/knowledge-base/split-pfx-file-into-pem-key-files-openss-windows-linux/

```
openssl pkcs12 -in test.pfx -out test_key.key -nocerts -nodes
openssl pkcs12 -in test.pfx -out test_cert.pem -nokeys -clcerts
```
3. I then executed osslsigncode. See  `windows-signing/sign.sh`.
4. Out comes a `sts_signed.exe`
5. I tested this on the windows machine and it worked. The publisher is also exactly as indicated `CN=Contoso Software, O=Contoso Corporation, C=US`
6. I did not timestamp the file. That needs to happen though

