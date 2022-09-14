#!/bin/sh

set -e && cd "$(dirname "$0")" && cd ..

#CHECK PATH IS DEFINED
if [ -z "${1}" ];
then
  echo "Missing path to local stackstate-openapi repository";
  display_help;
  exit;
fi
OPENAPIPATH="$1"

echo "Stackstate OpenAPI repository path: $OPENAPIPATH"

OUTPUT_DIR="generated/stackstate_api"

rm -rf "${OUTPUT_DIR}"

openapi-generator-cli generate -i "$OPENAPIPATH/spec/openapi.yaml" -g go  -c stackstate_openapi/openapi_generator_config.yaml -o "${OUTPUT_DIR}" -t stackstate_openapi/template

# we need to throw these files away, otherwise go gets upset
rm "${OUTPUT_DIR}"/go.mod
rm "${OUTPUT_DIR}"/go.sum