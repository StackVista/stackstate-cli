#!/bin/sh
set -e && cd "$(dirname "$0")" && cd ..

checkForVariable() {
  if [[ -z ${!1+set} ]]; then
    echo "Error: Define $1 environment variable"
    exit 1
  fi
}

checkForVariable AWS_DEFAULT_REGION
checkForVariable AWS_ACCESS_KEY_ID
checkForVariable AWS_SECRET_ACCESS_KEY
checkForVariable docker_user
checkForVariable docker_password

# Login to docker hub, so Goreleaser can push its docker images there.
echo "$docker_password" | docker login -u "$docker_user" --password-stdin docker.io

goreleaser release
scripts/publish/publish_latest_version_to_s3.sh
scripts/publish/publish_installers_to_s3.sh
