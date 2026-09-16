#!/usr/bin/env bash
set -euo pipefail

# Run against a prebuilt image; the helper never enters the scanned image layers.
image=${1:?usage: docker/test-timezones.sh IMAGE ARCH}
arch=${2:?usage: docker/test-timezones.sh IMAGE ARCH}
work=$(mktemp -d)
container=
cleanup() {
  if [[ -n "$container" ]]; then docker rm -f "$container" >/dev/null; fi
  rm -rf "$work"
}
trap cleanup EXIT
CGO_ENABLED=0 GOOS=linux GOARCH="$arch" go build -o "$work/fixture" ./docker/testdata/timezones.go
container=$(docker create --network=none --platform "linux/$arch" --entrypoint /fixture "$image")
docker cp "$work/fixture" "$container:/fixture"
docker start -a "$container"
test "$(docker inspect --format '{{.State.ExitCode}}' "$container")" = 0
