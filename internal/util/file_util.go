package util

import (
	"errors"
	"os"
)

func DoesFileExist(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err == nil { //nolint:gocritic
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}
