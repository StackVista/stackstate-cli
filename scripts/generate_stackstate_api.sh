#!/bin/bash
set -e

cd "$(dirname "$0")"
cd ..

rm -rf generated/stackstate_api
nix develop -c openapi-generator-cli generate -i stackstate_openapi/spec/api.yaml -g go  -c stackstate_openapi/openapi_generator_config.yaml -o generated/stackstate_api -t stackstate_openapi/template

# we need to throw these files away, otherwise go gets upset
rm generated/stackstate_api/go.mod
rm generated/stackstate_api/go.sum