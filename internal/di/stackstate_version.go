package di

import (
	"fmt"

	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func VersionToString(version stackstate_client.ServerVersion) string {
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
