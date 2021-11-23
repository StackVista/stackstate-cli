#!/bin/sh
rm -rf internal/openapi
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/stackstate_openapi/spec/api.yaml  \
    -g go \
    -c /local/openapi_generator_config.yaml \
    -o /local/internal/stackstate_client
rm internal/stackstate_client/go.mod
rm internal/stackstate_client/go.sum