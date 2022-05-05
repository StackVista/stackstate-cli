#!/bin/sh
set -e
cd "$(dirname "$0")" && cd ..

# Prints the latest version of the StackState CLI based on git semver git tags starting with a "v" and a number
git tag --list "v[0-9]*" --sort v:refname | tail -1
