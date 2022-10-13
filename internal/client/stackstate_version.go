package client

import (
	"fmt"

	"github.com/blang/semver/v4"
	"github.com/rs/zerolog/log"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

func VersionToString(version stackstate_api.ServerVersion) string {
	dev := ""
	if version.IsDev {
		dev = "-dev"
	}
	return fmt.Sprintf("%d.%d.%d+%s-%s%s",
		version.Major,
		version.Minor,
		version.Patch,
		version.Diff,
		version.Commit,
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
		return common.NewAPIVersionError(fmt.Errorf("Incompatible API version: got '%s', but need '%s'-compatible version", curr.String(), parsed.String()))
	}

	return nil
}
