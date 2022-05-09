#!/usr/bin/env bash
set -e && cd "$(dirname "$0")" && cd ..

OPENAPI_VERSION=`cat openapi_version`
CHECKOUT_DIR="stackstate_openapi/checkout"

rm -rf "$CHECKOUT_DIR"

# In gitlab we authenticate with the job token when cloning
if [[ -z "${CI_JOB_TOKEN}" ]]; then
  git clone https://gitlab.com/stackvista/platform/stackstate-openapi.git "$CHECKOUT_DIR"
else
  git clone "https://gitlab-ci-token:${CI_JOB_TOKEN}@gitlab.com/stackvista/platform/stackstate-openapi.git" "$CHECKOUT_DIR"
fi

git -C "$CHECKOUT_DIR" checkout "$OPENAPI_VERSION"

rm -rf generated/stackstate_api
docker run --rm -v "${PWD}:/local" \
    --user $(id -u):$(id -g) \
    openapitools/openapi-generator-cli:v6.0.0-beta generate \
    -i /local/stackstate_openapi/spec/openapi.yaml  \
    -g go \
    -c /local/stackstate_openapi/openapi_generator_config.yaml \
    -o /local/generated/stackstate_api \
    -t /local/stackstate_openapi/template \
    $@

# we need to throw these files away, otherwise go gets upset
rm generated/stackstate_api/go.mod
rm generated/stackstate_api/go.sum
