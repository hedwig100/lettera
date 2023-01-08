package paper

import (
	"errors"
	"os"
)

func isExist(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
