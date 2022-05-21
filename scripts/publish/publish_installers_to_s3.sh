#!/bin/sh
set -e && cd "$(dirname "$0")" && cd ../..

DEST="s3://cli-dl.stackstate.com/stackstate-cli/"
echo "Uploading installers to $DEST"
aws s3 cp scripts/publish/installers/install.sh $DEST
aws s3 cp scripts/publish/installers/uninstall.sh $DEST
aws s3 cp scripts/publish/installers/install.ps1 $DEST