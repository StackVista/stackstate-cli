#!/bin/sh
set -e && cd "$(dirname "$0")" && cd ..

if [ -z ${AWS_ACCESS_KEY_ID} ]; then echo "missing AWS_ACCESS_KEY_ID environment variable!"; exit 1; fi
goreleaser release
scripts/publish_latest_version_to_s3.sh
