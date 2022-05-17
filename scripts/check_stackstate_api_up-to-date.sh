#!/bin/sh
# This script checks whether the committed code in 
# `/generated/stackstate_api` is up to date
set -e
cd "$(dirname "$0")" && cd ..

# Check if you don't have uncommited changes.
if [ -n "$(git status --porcelain)" ]; then
  echo "ERROR: You've got uncommitted changes in tracked files. Please make sure you're on a clean commit.";
  exit 2
fi

./scripts/generate_stackstate_api.sh

# Check whether there are any changes to source files. 
# If so, this means some code did not get committed.
git add -A
if [[ $( git diff --cached --name-only ) ]]; then
  printf -- "There were changes in the working directory!\n"
  printf -- "This means some cached files should have been committed but were not\n"
  git diff --cached --name-only
  exit 1
else
  printf -- "No changes detected on working directory!\n\n"
fi
