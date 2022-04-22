#!/bin/bash
set -e

cd "$(dirname "$0")"
cd ..

rm -rf internal/stackstate_client
docker run --rm -v "${PWD}:/local" \
    --user $(id -u):$(id -g) \
    openapitools/openapi-generator-cli@sha256:890596803c279fa18a7d7206f5a4b775dbec0d72515521820fc866635915d99c generate \
    -i /local/stackstate_openapi/spec/api.yaml  \
    -g go \
    -c /local/stackstate_openapi/openapi_generator_config.yaml \
    -o /local/internal/stackstate_client \
    -t /local/stackstate_openapi/template \
    $@

# we need to throw these files away, otherwise go gets upset
rm internal/stackstate_client/go.mod
rm internal/stackstate_client/go.sum
