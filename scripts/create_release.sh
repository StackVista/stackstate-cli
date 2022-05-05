#!/usr/bin/env bash
set -e
cd "$(dirname "$0")" && cd ..

PROG=`basename "$0"`

NAT='0|[1-9][0-9]*'
ALPHANUM='[0-9]*[A-Za-z-][0-9A-Za-z-]*'
IDENT="$NAT|$ALPHANUM"
FIELD='[0-9A-Za-z-]'

SEMVER_REGEX="\
^[vV]?\
($NAT)\\.($NAT)\\.($NAT)\
(\\-(${IDENT})(\\.(${IDENT}))*)?\
(\\+${FIELD}(\\.${FIELD})*)?$"

USAGE="\
Usage:
  $PROG --version <version>

Arguments:
  <version>  A version must match the following regular expression:
             \"${SEMVER_REGEX}\"
             In English:
             -- The version must match X.Y.Z[-PRERELEASE][+BUILD]
                where X, Y and Z are non-negative integers.
             -- PRERELEASE is a dot separated sequence of non-negative integers and/or
                identifiers composed of alphanumeric characters and hyphens (with
                at least one non-digit). Numeric identifiers must not have leading
                zeros. A hyphen (\"-\") introduces this optional part.
             -- BUILD is a dot separated sequence of identifiers composed of alphanumeric
                characters and hyphens. A plus (\"+\") introduces this optional part.

  <other_version>  See <version> definition.

Options:
  -h, --help             Print this help message.

See also:
  https://semver.org"

function help {
  echo -e "$USAGE" >&2
}

function error {
  echo -e "Error: $1" >&2
}

# Parse arguments
OPTIND=1         # Reset in case getopts has been used previously in the shell.
while getopts "h?v:" opt; do
  case "$opt" in
    h|\?)
      help
      exit 0
      ;;
    v) 
      VERSION=$OPTARG
      ;;
    \? ) 
      error "invalid Option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done
shift $((OPTIND-1))
[ "${1:-}" = "--" ] && shift

if [ ! -n "$VERSION" ]; then
  error "missing version, use the -v flag"
  help
  exit 1
fi

if [[ ! "$VERSION" =~ $SEMVER_REGEX ]]; then
  error "version $VERSION does not match the semver scheme 'X.Y.Z(-PRERELEASE)(+BUILD)'. See help for more information."
  exit 2
fi

# Check if you don't have uncommited changes.
if [ -n "$(git status --porcelain)" ]; then
  error "you've got uncommitted changes in tracked files. Please make sure you're on a clean commit.";
  exit 2
fi

# Check if you're not ahead of the origin.
git remote update 2>&1 > /dev/null
if git status -uno | grep -q "Your branch is ahead of"; then
  error "your branch is ahead of the origin. Please make sure the commit you want to release is already on the origin.";
  exit 2
fi

# Ready to release!
git tag "v$VERSION" -m "v$VERSION"
git push origin "v$VERSION"