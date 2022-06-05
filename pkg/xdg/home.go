package xdg

import (
	"os"

	home "github.com/mitchellh/go-homedir"
)

// XDG spec https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
func GetXDGConfigHome() (xdgConfigHome string, err error) {
	xdgConfigHome = os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		xdgConfigHome, err = home.Expand("~/.config")
	}
	return
}
