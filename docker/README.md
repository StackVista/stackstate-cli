# Container timezone checks

The SUSE Observability CLI embeds Go's timezone database so named `TZ` values
continue to work in the BCI micro image without system zoneinfo. System-provided
timezone data still takes precedence. Updating the Go toolchain updates the
embedded database; an OS package scan alone does not establish its freshness.
The minimum Go version in `go.mod` is 1.27.1, which embeds IANA tzdata 2026c.
This includes the 2026b correction required by DLA-4569-1: Vancouver remains at
UTC-07 after November 1, 2026 (abbreviated MST in this database). Go 1.25.13
still applied the old winter fallback and displayed the previous date near midnight.

Build the Linux CLI with the toolchain selected by `go.mod`, then use the same
Dockerfile as GoReleaser:

```bash
mkdir -p /tmp/cli-build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /tmp/cli-build/sts .
docker build --platform linux/amd64 -f docker/Dockerfile.goreleaser \
  -t cli-timezones /tmp/cli-build
bash docker/test-timezones.sh cli-timezones amd64
```

Use `arm64` on an arm64 runner (or with local emulation). The fixture runs an HTTP
server on container loopback, with external networking disabled, and invokes the
image's actual CLI. It checks both agent timestamp columns, license and service
token expiry dates, and raw JSON epoch milliseconds. Literal UTC, New York winter
and summer, and Kathmandu expectations detect UTC fallback, DST and fractional
offset regressions. Vancouver checks cover the reported November 2026 affected
date, January 2027, and historical winter time, including unchanged JSON epochs.
The helper is copied into a disposable container, never into
the built/scanned image.

`container-timezones.yml` runs this check on both native architectures and retains
source/binary/image identities, runtime output, package inventories and raw
Trivy/Grype reports. Vulnerability scans include UNKNOWN and apply no exceptions
or VEX filtering; secrets are scanned separately. It builds candidates only.
