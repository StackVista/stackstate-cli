#!/bin/sh
set -e && cd "$(dirname "$0")" && cd ..



if output=$(git status --porcelain) && [ -z "$output" ]; then
  echo "You've got uncommitted changes in tracked files. Please make sure you're on a clean commit.";
  exit 1
fi

