#!/bin/sh
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/stackstate_openapi/spec/api.yaml  \
    -g go \
    -o /internal/openapi
