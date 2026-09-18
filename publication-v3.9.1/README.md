# CLI v3.9.1 publication evidence — 2026-09-18

Authorized by Louis; signed tag v3.9.1 (tag object 0d458b9c1d4c8513356a2a9a9b66cedad619fbf2) targets EXACT reviewed merge 42d67de511f7e9b0a6ff4bc4cfa80cdb999cede9. Local SSH signature and GitHub verification passed. Only this release tag was pushed to trigger production publication; application source was not changed. Raw verification reports were later retained in a signed fast-forward of the existing evidence branch (no CI/release trigger).

Release workflow: https://github.com/StackVista/stackstate-cli/actions/runs/35349892019 — SUCCESS, including lint/tests/license/config and all publication steps.
Release: https://github.com/StackVista/stackstate-cli/releases/tag/v3.9.1
PR: https://github.com/StackVista/stackstate-cli/pull/164
Ticket (supervisor-owned): https://github.com/StackVista/cve-reporter/issues/65

Docker Hub stackstate/stackstate-cli2:3.9.1 and :latest both resolve to sha256:4b827fcfc4cf4a0bb54067001ea8daf2ebd2d2a3da9713239fa560fa4cb9734c at verification.
- linux/amd64 / :3.9.1-x86_64: sha256:28e1756b62c30953be914bd54f50363e88f353293cd3d67c5a458f5b05ffcdd3
- linux/arm64 / :3.9.1-arm64v8: sha256:587222286b6497aae78ed7f4a527b512a6923fd86fa8645b6093045d23cbe803

Both images' base layers exactly match the approved BCI Nano index sha256:794d075f0ffa66f6fc1f7073b4957c4c797f29f04384fe547b34f865c7afeebb and its architecture manifests. Both have SLES15.7, timezone2026c, ca-certificates-mozilla-prebuilt2.84 and seven RPMs in BOTH Trivy and Grype inventories. Go binary inventories are present. Trivy0.74.0 and checksum-verified Grype0.118.0, with refreshed 2026-09-18 databases, found ZERO vulnerabilities including UNKNOWN. Separate Trivy secret scans found ZERO secrets. Raw reports and scan commands are retained alongside this summary.

Both image binaries are byte-identical to their Linux release archives. All five archive binaries embed Go1.25.13, CGO_ENABLED=0, version3.9.1, exact revision42d67de and vcs.modified=false. Actual amd64 container version execution also reports3.9.1/42d67de. SHA256 checksums validate all five archives; GitHub and https://dl.stackstate.com/stackstate-cli/v3.9.1/ serve byte-identical archives and checksum file. Public LATEST_VERSION is v3.9.1; all three installer scripts equal release source. Downloads were verified through the correct public host; no bucket listing/IAM assumptions were used.

Runtime/TZ/TLS applicability: merge tree equals independently accepted28f541d; compared with tested d15b40d4 only five validation/documentation files were deleted. Application/dependency files, Go toolchain, CGO_ENABLED=0 and Nano base/Dockerfile are identical. Reuse accepted native AMD64/ARM64 35 timezone/JSON assertions and TLS positive/missing-CA controls from https://github.com/StackVista/stackstate-cli/actions/runs/35330138263 (artifacts10540712513/10540697547), preserved with signed source at https://github.com/StackVista/stackstate-cli/tree/evidence/cli164-nano-d15b40d . Release metadata and artifact hashes necessarily differ. Native ARM64 runtime tests were not rerun, nor were Windows/macOS executables run; their release identity/checksums were verified.

Limits: no cosign-discoverable signature/attestation artifacts found for index or either architecture. Existing build config explicitly uses --provenance=false and defines no image-signing step. Signed source tag is not image authentication; binary metadata/checksums and verified base layers are identity evidence, not a signed build provenance claim. Other registry attachment formats were not exhaustively searched. Fresh publication evidence awaits SAME reviewer4a0 focused delivery acceptance. No deployment, unrelated release, IAM, VEX or shared changes; other holds preserved. Latest aliases are mutable and these observations are point-in-time.
