#!/bin/sh
set -e && cd "$(dirname "$0")" && cd ..

# Make sure we've got all tags
git remote update 2>&1 > /dev/null

# Prints the latest version of the StackState CLI based on git semver git tags starting with a "v" and a number
TAG_NAME=`git tag --list "v[0-9]*" --sort v:refname | tail -1`
VERSION=${TAG_NAME:1}
echo $VERSION
