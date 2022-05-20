#/bin/bash

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  OS=linux
elif [[ "$OSTYPE" == "darwin"* ]]; then
  OS=darwin
else
  echo "OS $OSTYPE not supported. Please checkout the CLI docs on doc.stackstate.com or contact StackState support for help with your OS."
  exit 1
fi

VERSION=`curl https://dl.stackstate.com/stackstate-cli/LATEST_VERSION 2> /dev/null` 
ARCH=`uname -m`
DL="https://dl.stackstate.com/stackstate-cli/v$VERSION/stackstate-cli-full-$VERSION.$OS-$ARCH.tar.gz"
echo "Downloading and unpacking: $DL"
curl $DL | tar xz --directory /usr/local/bin

# configure the CLI if the STS_URL And STS_API_TOKEN are given
if [[ -n "$STS_URL" && -n "$STS_API_TOKEN" ]]; then
  sts cli-config save --url $STS_URL --api-token $STS_API_TOKEN
fi

GREEN='\033[0;32m'
NO_COLOR='\033[0m'
printf "Success! Type '${GREEN}sts${NO_COLOR}' to get started!\n"
