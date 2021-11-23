#!/bin/bash
docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/spec/api.yaml \
  -g go \
  -o /local/generated
