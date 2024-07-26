# To install:
# curl -o- https://dl.stackstate.com/stackstate-cli/install.sh |STS_API_URL="url" STS_API_TOKEN="token" bash
#-----------------------------------
# Parameters to script
#-----------------------------------
# STS_CLI_VERSION  - version of the CLI to install (empty means latest)
# STS_URL - url of the StackState instance to configure (empty means don't configure)
# STS_API_TOKEN - API-TOKEN of the StackState instance to configure (empty means don't configure)
# STS_CLI_LOCATION - Path you want to install CLI (empty means `/usr/local/bin`)
#-----------------------------------

#!/usr/bin/env bash
GREEN='\033[0;32m'
RED='\033[0;31m'
NO_COLOR='\033[0m'

function error() {
  printf "${RED}[ERROR]${NO_COLOR} $1\n"
  exit 10
}

if [[ "$OSTYPE" == "linux"* ]]; then
  OS=linux
elif [[ "$OSTYPE" == "darwin"* ]]; then
  OS=darwin
else
  error "Unsupported operating system: $OSTYPE. Please checkout the CLI docs on docs.stackstate.com or contact StackState for support with your OS."
fi
ARCH=`uname -m`
if [[ "$ARCH" != "x86_64" && "$ARCH" != "arm64"  && "$ARCH" != "aarch64" ]]; then
  error "Unsupported architecture: $ARCH. Please checkout the CLI docs on docs.stackstate.com or contact StackState for support with your OS."
fi

# binaries are only published for arm64
if [[ "$ARCH" == "aarch64" ]]; then
  ARCH="arm64"
fi

# Check if custom location was defined
if [[ -z "$STS_CLI_LOCATION" ]]; then
  # Use default installation location
  TARGET_CLI_PATH="/usr/local/bin"
  # check if the user has permissions to write on default location
  if [[ ! -w "$TARGET_CLI_PATH" ]]; then
    # user has writing permissions, so no need to use sudo
    echo "Cannot write to the current directory. Please either execute the script from a writeable directory or set STS_CLI_LOCATION to a different directory."
    exit 1
  else
    NO_SUDO=true
  fi
else
  # Check if the custom installation location is valid
  if [[ ! -d "$STS_CLI_LOCATION" ]]; then
    error "Provided directory does not exist: $STS_CLI_LOCATION."
  # Check if the user has writing permissions on custom location
  elif [[ ! -w "$STS_CLI_LOCATION" ]]; then
    # Location exists but user doesn't have writing permission.
    echo "Sudo will be used on the provided location $STS_CLI_LOCATION."
  else
    # Location exists and user has writing permission
    NO_SUDO=true
  fi
  # Set installation location as defined by the input
  TARGET_CLI_PATH="$STS_CLI_LOCATION"
fi

# Download and unpack the CLI to the target CLI path
if [[ -z "$STS_CLI_VERSION" ]]; then
  STS_CLI_VERSION=`curl https://dl.stackstate.com/stackstate-cli/LATEST_VERSION 2> /dev/null`
  # The LATEST_VERSION file contains the published tag name, strip the v prefix
  STS_CLI_VERSION=${STS_CLI_VERSION#v}
fi
DL="https://dl.stackstate.com/stackstate-cli/v${STS_CLI_VERSION}/stackstate-cli-${STS_CLI_VERSION}.$OS-$ARCH.tar.gz"
echo "Installing: $DL"

if [[ -z "$NO_SUDO" ]]; then
  # check if custom location was passed to avoid redundant printing
  if [[ -z "$STS_CLI_LOCATION" ]]; then
    echo "STS requires sudo permission to install."
    echo "Alternatively, you can provide a custom location with STS_CLI_LOCATION="
    echo "Make sure that the provided 'STS_CLI_LOCATION' is in your OS Path."
  fi

  # sudo password will be asked when executing the command.
  curl $DL | sudo tar xz --directory $TARGET_CLI_PATH
else
  curl $DL | tar xz --directory $TARGET_CLI_PATH
fi

# Verify that 'sts' works
sts > /dev/null 2>&1
if [ $? -ne 0 ]; then
  error "Can not find 'sts' on the path or execute it"
fi

# Configure the CLI if config parameters have been set
if [[ -n "$STS_URL" && -n "$STS_API_TOKEN" ]]; then
  sts context save --url $STS_URL --api-token $STS_API_TOKEN
fi

printf "Success! Type ${GREEN}sts${NO_COLOR} to get started!\n"
