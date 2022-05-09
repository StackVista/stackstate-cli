#!/bin/sh
set -e && cd "$(dirname "$0")" && cd ..

golangci-lint run -v
