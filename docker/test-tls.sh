#!/usr/bin/env bash
set -euo pipefail

image=${1:?usage: docker/test-tls.sh IMAGE}
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT

# The registry is not a CLI API: its HTTP 404 proves the actual CLI completed TLS.
# Use a dummy token, never credentials. A timeout/network error cannot pass.
args=(agent list --url https://registry.suse.com --api-token fixture)
if timeout 30 docker run --rm -e NO_COLOR=1 -e TERM=dumb "$image" "${args[@]}" > "$work/trusted" 2>&1; then
  echo "Expected an API error from the non-API endpoint" >&2
  exit 1
fi
cat "$work/trusted"
grep -Fq '(404 Not Found)' "$work/trusted"
echo 'PASS actual CLI reaches HTTP through TLS using image CA certificates'

if timeout 30 docker run --rm -e NO_COLOR=1 -e TERM=dumb \
  -e SSL_CERT_FILE=/dev/null -e SSL_CERT_DIR=/no-certificates \
  "$image" "${args[@]}" > "$work/untrusted" 2>&1; then
  echo "Expected certificate verification to fail without CA certificates" >&2
  exit 1
fi
cat "$work/untrusted"
grep -Fq 'x509: certificate signed by unknown authority' "$work/untrusted"
echo 'PASS actual CLI rejects TLS when system CA certificates are disabled'
