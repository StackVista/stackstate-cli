# Container checks

The SUSE Observability CLI uses digest-pinned BCI Nano, which supplies timezone
data and CA certificates for the static release binary. The image's tzdata 2026c
includes the DLA-4569-1 correction: Vancouver stays at UTC-07 in winter 2026
(abbreviated MST). There is no embedded Go timezone database or special toolchain
requirement for timezone data; refresh the base digest to update its packages.

Build with the toolchain selected by `go.mod` and the release Dockerfile:

```bash
mkdir -p /tmp/cli-build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /tmp/cli-build/sts .
docker build --platform linux/amd64 -f docker/Dockerfile.goreleaser \
  -t cli-timezones /tmp/cli-build
bash docker/test-timezones.sh cli-timezones amd64
bash docker/test-tls.sh cli-timezones
```

Use `arm64` on an arm64 runner (or with local emulation). The timezone fixture
invokes the actual CLI against a loopback API with external networking disabled.
Its 35 assertions cover both agent timestamp columns, license/service-token
expiry dates and unchanged JSON epochs: UTC, New York winter/summer, Kathmandu,
and Vancouver historical winter, November 2026 and January 2027. Expectations
are literals independent of the fixture's timezone database. The helper is copied
into a disposable container, never into the built/scanned image.

The TLS check uses the actual CLI against a public HTTPS endpoint and checks that
it reaches an HTTP response. A negative control disables the system CA paths and
must fail certificate verification. This check requires internet access.

The existing `container-timezones.yml` runs on both native architectures and
retains binary/image identities, runtime output, inventories and raw scan reports.
Vulnerability scans include UNKNOWN with no exceptions or VEX filtering; secrets
are scanned separately. Empty inventories fail validation. This workflow builds
candidates only; normal CI supplies unit tests, lint, license and release-config
checks.
