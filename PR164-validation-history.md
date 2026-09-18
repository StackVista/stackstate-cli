Use digest-pinned SUSE BCI Nano 15.7 for the static SUSE Observability CLI. Nano supplies tzdata 2026c and CA certificates, so Vancouver keeps the corrected UTC-07 winter time without embedding a second timezone database in the binary.

Removed the embedded `time/tzdata` import and all workaround-only Go 1.27.1/linter 2.13.2/configuration/`reflect.Pointer` changes. Go 1.25.13, linter 2.9.0 and application source now match the main-branch baseline. Release builds already use `CGO_ENABLED=0`.

Retained the 35 meaningful actual-CLI timezone/JSON assertions and existing native amd64/arm64 candidate workflow, since normal CI does not build and scan both container architectures. Added a small actual-CLI TLS check with a missing-CA negative control. Removed redundant scanner metadata files; raw scan reports retain scanner provenance. Inventories, source/binary/image identities, UNKNOWN-inclusive vulnerability scans and separate secret scans remain. There is no new workflow or production publishing path.

Signed Nano candidate: `d15b40d4add021e0867b97624c450f9728e0d7ee`. Base index: `sha256:794d075f0ffa66f6fc1f7073b4957c4c797f29f04384fe547b34f865c7afeebb` (amd64 `95e1de2733c016b004bcdda59a507c09c9843e914bba08748c31c7f7dc3c5325`, arm64 `e1bb2f96b26dc1faa11f9fcbb8cad5bb07b58ef3a0d2c90c7be8c4f661d844d8`).

Local amd64 validation: all 35 timezone assertions and both TLS controls passed; actionlint and offline Zizmor for the candidate workflow passed. Changed-head validation passed:

- [Normal CI 35330138280](https://github.com/StackVista/stackstate-cli/actions/runs/35330138280): tests, linter, license scan and GoReleaser check passed; publication skipped.
- [Native candidate workflow 35330138263](https://github.com/StackVista/stackstate-cli/actions/runs/35330138263): both amd64 and arm64 passed all 35 timezone/JSON assertions and both TLS controls, using release-style Go 1.25.13 binaries from this exact signed head.
- Downloaded raw reports identify SLES 15.7 and seven OS packages in both Trivy and Grype inventories, including `timezone` 2026c and `ca-certificates-mozilla-prebuilt` 2.84. Trivy also inventories 98 Go binary packages. Zero Trivy/Grype vulnerability findings and zero secrets on both architectures; no findings suppressed.
- [amd64 artifact 10540712513](https://github.com/StackVista/stackstate-cli/actions/runs/35330138263/artifacts/10540712513): binary SHA256 `569e0741656e5a7a0765937f4c6876c2f2492ff8279ebab937c9b1dc815c498e`; image ID `sha256:b76ed90b31d6cd31426e8807a88ef5f7aac906ad557c533979a0ab2138643270`.
- [arm64 artifact 10540697547](https://github.com/StackVista/stackstate-cli/actions/runs/35330138263/artifacts/10540697547): binary SHA256 `05be29a2cb1abe3a7d8c4391776276aa5d3feb12978ed788b3542f610bb719e8`; image ID `sha256:820eca337fe827d6739eb301faa16957a3333f0e2e71bdeeba8b343fb20c76d2`.

 TLS validation expects the SUSE registry's HTTP 404 (it is not a CLI API), then proves certificate verification fails with CA paths disabled. It requires registry availability.

The [original independent reviewer](https://omnigent.tooling.stackstate.io/c/4a0f96a272e75dde97385eb0952413ae) accepted signed Nano head `d15b40d4add021e0867b97624c450f9728e0d7ee` (completion `d1af9a4c7d7543e9a5fcb8c990cdb4fb`). Live head and successful normal/native CI were verified unchanged. Ready for human review at Louis’s explicit request; auto-merge remains disabled. Tracks https://github.com/StackVista/cve-reporter/issues/65. Human merge/release, published-image assessment and delivery evidence remain outstanding.

<details><summary>Preserved prior direction, signed candidates and evidence</summary>

**Withdrawn from the merge-ready queue at Louis's request; Nano rework is pending.** Reuse this PR and the original correction implementer. Replace Micro plus embedded timezone data with a digest-pinned BCI Nano approach, removing Go/lint/other changes needed solely for that workaround unless a separate current requirement justifies them. Prefer simple product behavior and existing validation paths; keep focused meaningful timezone regressions.

Validate the actual release-style static CLI on both architectures: named zones including Vancouver, fractional offsets and DST, unchanged JSON, TLS trust, usable OS/package inventories, vulnerability scans and separate secret scans. Operator Nano base inspection is a lead, not actual CLI validation; missing inventories cannot establish a clean result. The final changed head must be reviewed by the same independent reviewer before another human merge handoff.

Tracking: https://github.com/StackVista/cve-reporter/issues/65. Preserve signed head `37637ae44e2b1509689073d815275fdcea74ccd3` and all prior evidence below. Its previous acceptance does not cover the Nano change. No merge, release or deployment is authorized.

<details><summary>Preserved Micro/embedded-timezone candidate and evidence; superseded implementation direction</summary>

Preserve named `TZ` behavior in the SUSE Observability CLI's digest-pinned BCI micro image by embedding Go timezone data. Go 1.27.1 includes tzdata 2026c and the DLA-4569-1 correction: at epoch `1793604600000`, Vancouver displays `2026-11-02 00:30:00 MST` (UTC-07), and license/service-token expiry displays `2026-11-02`. Go 1.25.13 incorrectly displayed the previous day.

The existing native amd64/arm64 fixture now has 35 assertions: the original UTC, New York and Kathmandu cases, plus November 2026, January 2027 and historical Vancouver winter dates. JSON epochs remain unchanged. System timezone data retains precedence. No additional CI gate was added.

The toolchain update needs golangci-lint 2.13.2. Its goconst configuration preserves v2.9.0 coverage (composite literal checks were added later); `reflect.Ptr` is replaced by its equivalent `reflect.Pointer` name for the new vet check. Existing BCI package and Grype inventory corrections remain intact.

Signed candidate: `37637ae44e2b1509689073d815275fdcea74ccd3`. Local validation: 35 amd64 container assertions, full Go tests, final lint (zero issues), focused util tests and actionlint pass. Offline Zizmor findings match the baseline exactly: 32 existing findings in `ci.yml`, none in `container-timezones.yml`. [Normal CI 35321211106](https://github.com/StackVista/stackstate-cli/actions/runs/35321211106) passed tests, lint, license and GoReleaser checks; release publication was skipped. [Native candidate CI 35321211324](https://github.com/StackVista/stackstate-cli/actions/runs/35321211324) passed on both architectures. Downloaded evidence confirms this exact source head and Go 1.27.1, 35 fixture assertions per architecture, SLES 15.7, 21 RPMs, zero Trivy/Grype vulnerability matches (UNKNOWN included), and zero secrets.

Raw reports, runtime output and source/binary/image identities: [amd64 artifact 10537022490](https://github.com/StackVista/stackstate-cli/actions/runs/35321211324/artifacts/10537022490), [arm64 artifact 10537286703](https://github.com/StackVista/stackstate-cli/actions/runs/35321211324/artifacts/10537286703). GitHub verifies the commit signature.

Upstream provenance: [Go 1.27.1 timezone data version](https://github.com/golang/go/blob/go1.27.1/lib/time/update.bash), [tzdata 2026c Vancouver rules](https://github.com/eggert/tz/blob/2026c/northamerica), and [DLA-4569-1](https://www.debian.org/lts/security/2026/dla-4569). The database's MST abbreviation represents the required UTC-07 offset.

Tracks https://github.com/StackVista/cve-reporter/issues/65. [Independent re-review](https://omnigent.tooling.stackstate.io/c/4a0f96a272e75dde97385eb0952413ae) accepted `37637ae44e2b1509689073d815275fdcea74ccd3`: prior P2 DLA-4569-1 is resolved and no required findings remain. Ready for human review; operator supplied the human handoff. System timezone precedence was confirmed by source inspection, not a separate runtime assertion. Human merge/release, actual published-image assessment and complete clean delivery remain outstanding. Optional timezone freshness monitoring is not a gate.

<details><summary>Retained prior candidate evidence (superseded head)</summary>

Preserve named `TZ` behavior after #162 by embedding `time/tzdata` in the SUSE Observability CLI. Agent timestamps and license/service-token expiry dates retain local dates in BCI micro without system zoneinfo. The embedded database follows the Go toolchain; system timezone data takes precedence.

Focused native amd64/arm64 container checks cover UTC, New York winter/summer, Kathmandu and unchanged JSON timestamps, with separate secret scans and UNKNOWN-inclusive vulnerability scans. Candidate evidence retains source/binary/image identities and inventories.

Recovery from [run 35117696829](https://github.com/StackVista/stackstate-cli/actions/runs/35117696829): Grype JSON has no top-level `artifacts`; inventory validation now uses its CycloneDX output. Both reports also contained six matches: CVE-2026-54369/54370/54371 against libacl1 and libattr1. Refreshing the BCI micro 15.7 multi-architecture digest supplies the reported fixed RPM versions. No findings are suppressed and the zero-match gate remains enforced.

Validation at signed head `b387844594ceeac3011cdda635a7daea8a288af5`: [native build/runtime/scans](https://github.com/StackVista/stackstate-cli/actions/runs/35120491312) passed on both architectures: Go 1.25.13, 20 passing fixture checks each, 21 RPMs, zero Trivy/Grype vulnerability findings and zero secrets. Raw reports and source/binary/image identities: [amd64 artifact](https://github.com/StackVista/stackstate-cli/actions/runs/35120491312/artifacts/10456079360), [arm64 artifact](https://github.com/StackVista/stackstate-cli/actions/runs/35120491312/artifacts/10457230229). [Existing CI](https://github.com/StackVista/stackstate-cli/actions/runs/35120491087): tests, lint, license and GoReleaser checks all passed. Gate controls reject missing/empty inventories and the retained vulnerable reports. Changed workflow passes offline Zizmor; full-repo findings remain in unchanged `ci.yml`.

Tracks https://github.com/StackVista/cve-reporter/issues/65. Changed-head independent review and human merge/release remain outstanding; this candidate does not close the delivery tracker.


</details>

</details>


</details>


