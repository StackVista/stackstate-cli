#/bin/bash


VERSION=`curl https://dl.stackstate.com/stackstate-cli/LATEST_VERSION 2> /dev/null` 
DL="https://dl.stackstate.com/stackstate-cli/v$VERSION/stackstate-cli-full-$VERSION.darwin-x86_64.tar.gz"
echo "Downloading: $DL"
curl -fLo stackstate-cli.tar.gz $DL 

tar xzvf stackstate-cli.tar.gz --directory /usr/local/bin > dev/null 2>&1
rm stackstate-cli.tar.gz

sts cli-config save --url $URL --api-token $API_TOKEN