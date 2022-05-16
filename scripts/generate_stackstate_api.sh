#!/bin/sh

set -e && cd "$(dirname "$0")" && cd ..

OPENAPI_VERSION=$(cat openapi_version)
CHECKOUT_DIR="stackstate_openapi/checkout"
OUTPUT_DIR="generated/stackstate_api"

rm -rf "$CHECKOUT_DIR"

# In gitlab we authenticate with the job token when cloning
if [[ -z "${CI_JOB_TOKEN}" ]]; then
  git clone https://gitlab.com/stackvista/platform/stackstate-openapi.git "$CHECKOUT_DIR"
else
  git clone "https://gitlab-ci-token:${CI_JOB_TOKEN}@gitlab.com/stackvista/platform/stackstate-openapi.git" "$CHECKOUT_DIR"
fi

git -C "$CHECKOUT_DIR" checkout "$OPENAPI_VERSION"

rm -rf "${OUTPUT_DIR}"
openapi-generator-cli generate -i "$CHECKOUT_DIR/spec/openapi.yaml" -g go  -c stackstate_openapi/openapi_generator_config.yaml -o "${OUTPUT_DIR}" -t stackstate_openapi/template

# we need to throw these files away, otherwise go gets upset
rm "${OUTPUT_DIR}"/go.mod
rm "${OUTPUT_DIR}"/go.sum
