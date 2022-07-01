package xdg

import (
	"os"

	home "github.com/mitchellh/go-homedir"
)

// XDG spec https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
// TODO do we want to replace this with `os.UserConfigDir()`? That will change behaviour on MacOS, but adds support for Windows standard location
// TODO add support for XDG_CONFIG_DIRS
// Another option is to use: https://github.com/adrg/xdg
func GetXDGConfigHome() (xdgConfigHome string, err error) {
	xdgConfigHome = os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		xdgConfigHome, err = home.Expand("~/.config")
	}
	return
}
