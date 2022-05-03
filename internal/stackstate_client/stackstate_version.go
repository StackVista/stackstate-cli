package stackstate_client

import (
	"fmt"

	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
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
