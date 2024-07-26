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

if [[ "${OSTYPE}" == "linux"* ]]; then
  OS=linux
elif [[ "${OSTYPE}" == "darwin"* ]]; then
  OS=darwin
else
  error "Unsupported operating system: ${OSTYPE}. Please checkout the CLI docs on docs.stackstate.com or contact StackState for support with your OS."
fi
ARCH=`uname -m`
if [[ "${ARCH}" != "x86_64" && "${ARCH}" != "arm64"  && "${ARCH}" != "aarch64" ]]; then
  error "Unsupported architecture: ${ARCH}. Please checkout the CLI docs on docs.stackstate.com or contact StackState for support with your OS."
fi

# binaries are only published for arm64
if [[ "${ARCH}" == "aarch64" ]]; then
  ARCH="arm64"
fi

TARGET_CLI_PATH="${STS_CLI_LOCATION:-/usr/local/bin}"

echo "Trying to install StackState CLI to ${TARGET_CLI_PATH}"

# Check if the destination directory exists
if [[ ! -d "${TARGET_CLI_PATH}" ]]; then
  error "The directory to install cli to does not exist: ${TARGET_CLI_PATH}."
fi

# Check if the user has writing permissions on the destination directory
if [[ ! -w "${TARGET_CLI_PATH}" ]]; then
  # Destination directory exists but user doesn't have writing permission.
  echo "Sudo will be used on the provided destination directory ${TARGET_CLI_PATH}."
  SUDO_REQUIRED="true"
fi

# Download and unpack the CLI to the target CLI path
if [[ -z "$STS_CLI_VERSION" ]]; then
  STS_CLI_VERSION=`curl https://dl.stackstate.com/stackstate-cli/LATEST_VERSION 2> /dev/null`
  # The LATEST_VERSION file contains the published tag name, strip the v prefix
  STS_CLI_VERSION=${STS_CLI_VERSION#v}
fi
DL="https://dl.stackstate.com/stackstate-cli/v${STS_CLI_VERSION}/stackstate-cli-${STS_CLI_VERSION}.${OS}-${ARCH}.tar.gz"
echo "Installing: $DL"

if [[ "$SUDO_REQUIRED" == "true" ]]; then
  echo "STS requires sudo permission to install."
  echo "Alternatively, you can provide a custom destination directory with STS_CLI_LOCATION="
  echo "Make sure that the provided 'STS_CLI_LOCATION' is in your OS PATH."

  # sudo password will be asked when executing the command.
  curl $DL | sudo tar xz --directory ${TARGET_CLI_PATH}
else
  curl $DL | tar xz --directory ${TARGET_CLI_PATH}
fi

# Verify that 'sts' works
${TARGET_CLI_PATH}/sts > /dev/null 2>&1

# Configure the CLI if config parameters have been set
if [[ -n "${STS_URL}" && -n "${STS_API_TOKEN}" ]]; then
  ${TARGET_CLI_PATH}/sts context save --url ${STS_URL} --api-token ${STS_API_TOKEN}
fi

if [ "$(whereis -q sts)" == "" ]; then
  error "Can not find 'sts' on the PATH or execute it. Consider adding the directory to your PATH: PATH=\"\$PATH:${TARGET_CLI_PATH}\""
fi

printf "Success! Type ${GREEN}sts${NO_COLOR} to get started!\n"
