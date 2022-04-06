#/bin/bash
./osslsigncode sign -certs test_cert.pem  -key test_key.key  \
    -n "Your Application" -i http://www.yourwebsite.com/ \
    -in ../dist/stackstate-cli-full_windows_amd64/sts.exe -out sts_signed.exe
