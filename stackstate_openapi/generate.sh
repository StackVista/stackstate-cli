#!/bin/bash
set -e

cd "$(dirname "$0")"
cd ..

rm -rf generated/stackstate_api
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v5.3.0 generate \
    -i /local/stackstate_openapi/spec/api.yaml  \
    -g go \
    -c /local/stackstate_openapi/openapi_generator_config.yaml \
    -o /local/generated/stackstate_api \
    -t /local/stackstate_openapi/template

# we need to throw these files away, otherwise go gets upset
rm generated/stackstate_api/go.mod
rm generated/stackstate_api/go.sum