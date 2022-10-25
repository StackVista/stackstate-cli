package client

import (
	"fmt"

	"github.com/blang/semver/v4"
	"github.com/rs/zerolog/log"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
)

func VersionToString(version stackstate_api.ServerVersion) string {
	dev := ""
	if version.IsDev {
		dev = "-dev"
	}
	diff := ""
	if version.Diff != "" {
		diff = "+" + version.Diff
	}
	commit := ""
	if version.Commit != "" {
		commit = "-" + version.Commit
	}
	return fmt.Sprintf("%d.%d.%d%s%s%s",
		version.Major,
		version.Minor,
		version.Patch,
		diff,
		commit,
		dev,
	)
}

// Checks if the version of the connected instance is at least minVersion according to SemVer.
func CheckVersionCompatibility(version stackstate_api.ServerVersion, minVersion string) common.CLIError {
	currVersion := VersionToString(version)
	log.Info().
		Str("server-version", currVersion).
		Str("min-supported-version", minVersion).
		Msg("Checking API version compatibility")

	// NOTE If no version is specified, accepts anything.
	if minVersion == "" {
		return nil
	}

	curr, err := semver.Make(currVersion)
	if err != nil {
		return common.NewAPIVersionError(err)
	}

	parsed, err := semver.Make(minVersion)
	if err != nil {
		return common.NewAPIVersionError(err)
	}

	if parsed.Major != curr.Major || !curr.GTE(parsed) {
		return NewAPIVersionMismatchError(curr.String(), parsed.String())
	}

	return nil
}

func NewAPIVersionMismatchError(got string, expected string) common.CLIError {
	return common.NewAPIVersionError(fmt.Errorf("Incompatible API version: got '%s', but need '%s'-compatible version", got, expected))
}
