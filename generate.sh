#!/bin/sh

# Throw everything away except for the mock files. 
# You ask WHY?
# The OpenAPI request structs are non-exported members. This means they can not be
# intercepted in a mock that is not within the same directory. The long term solution
# would be to generate the mocks along with the rest of the code, but that may be
# something for later
find internal/stackstate_client -type f ! -name 'mock_*' -delete

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/stackstate_openapi/spec/api.yaml  \
    -g go \
    -c /local/openapi_generator_config.yaml \
    -o /local/internal/stackstate_client

# we need to throw these files away, otherwise go gets upset
rm internal/stackstate_client/go.mod
rm internal/stackstate_client/go.sum