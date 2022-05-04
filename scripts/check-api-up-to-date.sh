#!/bin/bash
# this script checks whether the committed api is up to date

set -e

cd "$(dirname "$0")"
cd ..

./scripts/generate_stackstate_api.sh

# This script will check whether there are any changes to source files. If so, this means some code did not get committed.

git add -A

# Check if there are changes now leftover
if [[ $( git diff --cached --name-only ) ]]; then
  printf -- "There were changes in the working directory!\n"
  printf -- "This means some cached files should have been committed but were not\n"
  git diff --cached --name-only
  exit 1
else
  printf -- "No changes detected on working directory!\n\n"
fi
