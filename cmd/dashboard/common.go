package dashboard

import (
	"fmt"
)

// ResolveDashboardIdOrUrn resolves ID or identifier to a string that can be used with the API
// Returns the resolved identifier string or an error if neither is provided
func ResolveDashboardIdOrUrn(id int64, identifier string) (string, error) {
	switch {
	case id != 0:
		return fmt.Sprintf("%d", id), nil
	case identifier != "":
		return identifier, nil
	default:
		return "", fmt.Errorf("either --id or --identifier must be provided")
	}
}
