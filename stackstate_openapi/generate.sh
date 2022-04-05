#!/bin/bash

cd "$(dirname "$0")"
cd ..

rm -rf internal/stackstate_client
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v5.3.0 generate \
    -i /local/stackstate_openapi/spec/api.yaml  \
    -g go \
    -c /local/stackstate_openapi/openapi_generator_config.yaml \
    -o /local/internal/stackstate_client \
    -t /local/stackstate_openapi/template

# we need to throw these files away, otherwise go gets upset
rm internal/stackstate_client/go.mod
rm internal/stackstate_client/go.sum