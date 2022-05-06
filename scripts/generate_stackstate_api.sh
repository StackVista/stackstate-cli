#!/bin/bash
set -e

cd "$(dirname "$0")"
cd ..

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
#    openapitools/openapi-generator-cli@sha256:890596803c279fa18a7d7206f5a4b775dbec0d72515521820fc866635915d99c generate \

# we need to throw these files away, otherwise go gets upset
rm generated/stackstate_api/go.mod
rm generated/stackstate_api/go.sum
