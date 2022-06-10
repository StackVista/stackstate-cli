# To install:
# curl -o- https://dl.stackstate.com/stackstate-cli/install.sh |STS_API_URL="url" STS_API_TOKEN="token" bash
#-----------------------------------
# Parameters to script
#-----------------------------------
# STS_CLI_VERSION  - version of the CLI to install (empty means latest)
# STS_URL - url of the StackState instance to configure (empty means don't configure)
# STS_API_TOKEN - API-TOKEN of the StackState instance to configure (empty means don't configure)
#-----------------------------------

#!/usr/bin/env bash
GREEN='\033[0;32m'
RED='\033[0;31m'
NO_COLOR='\033[0m'

function error() {
  printf "${RED}[ERROR]${NO_COLOR} $1\n"
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

# Download and unpack the CLI to the target CLI path
TARGET_CLI_PATH=/usr/local/bin
if [[ -z "$STS_CLI_VERSION" ]]; then
  STS_CLI_VERSION=`curl https://dl.stackstate.com/stackstate-cli/LATEST_VERSION 2> /dev/null`
fi
DL="https://dl.stackstate.com/stackstate-cli/v${STS_CLI_VERSION}/stackstate-cli-${STS_CLI_VERSION}.$OS-$ARCH.tar.gz"
echo "Installing: $DL"
curl $DL | tar xz --directory $TARGET_CLI_PATH

# Verify that 'sts' works
sts > /dev/null 2>&1
if [ $? -ne 0 ]; then
  error "Can not find 'sts' on the path or execute it?"
fi

# Configure the CLI if config parameters have been set
if [[ -n "$STS_URL" && -n "$STS_API_TOKEN" ]]; then
  sts context save --url $STS_URL --api-token $STS_API_TOKEN
fi

printf "Success! Type ${GREEN}sts${NO_COLOR} to get started!\n"
