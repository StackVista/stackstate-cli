package term

import (
	"os"
	"strings"
)

func SupportsColor() bool {
	if strings.ToLower(os.Getenv("TERM")) == "dumb" {
		return false
	}
	// Implementation of https://no-color.org/
	if _, noColorEnvExists := os.LookupEnv("NO_COLOR"); noColorEnvExists {
		return false
	}

	return true
}
