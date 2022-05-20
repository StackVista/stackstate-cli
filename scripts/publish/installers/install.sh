#!/usr/bin/env bash

GREEN='\033[0;32m'
RED='\033[0;31m'
NO_COLOR='\033[0m'

function error() {
  echo "${RED}[ERROR]${NO_COLOR} $1"
  exit 10
}

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  OS=linux
elif [[ "$OSTYPE" == "darwin"* ]]; then
  OS=darwin
else
  error "Unsupported operating system: $OSTYPE. Please checkout the CLI docs on docs.stackstate.com or contact StackState for support with your OS."
fi
ARCH=`uname -m`
if [[ "$ARCH" != "x86_64" && "$ARCH" != "arm64" ]]; then
  error "Unsupported architecture: $ARCH. Please checkout the CLI docs on docs.stackstate.com or contact StackState for support with your OS."
fi

VERSION=`curl https://dl.stackstate.com/stackstate-cli/LATEST_VERSION 2> /dev/null` 
DL="https://dl.stackstate.com/stackstate-cli/v$VERSION/stackstate-cli-full-$VERSION.$OS-$ARCH.tar.gz"
echo "Downloading and unpacking: $DL"
curl $DL | tar xz --directory /usr/local/bin

# configure the CLI if the STS_URL and STS_API_TOKEN are provided
if [[ -n "$STS_URL" && -n "$STS_API_TOKEN" ]]; then
  sts cli-config save --url $STS_URL --api-token $STS_API_TOKEN
fi

printf "Success! Type ${GREEN}sts${NO_COLOR} to get started!\n"