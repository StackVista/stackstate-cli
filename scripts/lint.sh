#!/bin/bash
set -e
cd "$(dirname "$0")" && cd ..

nix develop -c golangci-lint run -v
