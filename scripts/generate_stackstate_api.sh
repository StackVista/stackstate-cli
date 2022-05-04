#!/bin/bash
set -e

cd "$(dirname "$0")"
cd ..

OPENAPI_VERSION=`cat openapi_version`
CHECKOUT_DIR="stackstate_openapi/checkout"

rm -rf "$CHECKOUT_DIR"
git clone https://gitlab.com/stackvista/platform/stackstate-openapi.git "$CHECKOUT_DIR"
git -C "$CHECKOUT_DIR" checkout "$OPENAPI_VERSION"

rm -rf generated/stackstate_api
nix develop -c openapi-generator-cli generate -i "$CHECKOUT_DIR/spec/api.yaml" -g go  -c stackstate_openapi/openapi_generator_config.yaml -o generated/stackstate_api -t stackstate_openapi/template

# we need to throw these files away, otherwise go gets upset
rm generated/stackstate_api/go.mod
rm generated/stackstate_api/go.sum