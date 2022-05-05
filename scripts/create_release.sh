#!/bin/sh
set -e && cd "$(dirname "$0")" && cd ..

echo "Updating git from remote"
git remote update

if [ -n "$(git status --porcelain)" ]; then
  echo "You've got uncommitted changes in tracked files. Please make sure you're on a clean commit.";
  exit 1
fi

