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

# Main API
OUTPUT_DIR="generated/stackstate_api"

rm -rf "${OUTPUT_DIR}"

openapi-generator-cli generate -i "$OPENAPIPATH/spec/openapi.yaml" -g go  -c stackstate_openapi/openapi_generator_config.yaml -o "${OUTPUT_DIR}" -t stackstate_openapi/template

# we need to throw these files away, otherwise go gets upset
rm "${OUTPUT_DIR}"/go.mod
rm "${OUTPUT_DIR}"/go.sum

# format code and clear unused imports
goimports -w $OUTPUT_DIR

# Admin API
ADMIN_OUTPUT_DIR="generated/stackstate_admin_api"

rm -rf "${ADMIN_OUTPUT_DIR}"

openapi-generator-cli generate -i "$OPENAPIPATH/spec/openapi_admin.yaml" -g go  -c stackstate_openapi/openapi_generator_config.yaml -o "${ADMIN_OUTPUT_DIR}" -t stackstate_openapi/template

# we need to throw these files away, otherwise go gets upset
rm "${ADMIN_OUTPUT_DIR}"/go.mod
rm "${ADMIN_OUTPUT_DIR}"/go.sum

# format code and clear unused imports
goimports -w $ADMIN_OUTPUT_DIR