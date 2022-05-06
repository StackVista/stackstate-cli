#!/bin/sh
set -e && cd "$(dirname "$0")" && cd ..

mkdir -p dist
LATEST_VERSION=`./scripts/print_latest_version.sh`
echo $LATEST_VERSION > dist/LATEST_VERSION

DEST="s3://cli-dl.stackstate.com/stackstate-cli/"
echo "Uploading LATEST_VERSION=$LATEST_VERSION to ${DEST}LATEST_VERSION"
aws s3 cp dist/LATEST_VERSION s3://cli-dl.stackstate.com/stackstate-cli/