#!/bin/sh
set -e && cd "$(dirname "$0")" && cd ..

go test ./... -v
