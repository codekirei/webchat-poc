package utils

import (
	"os"
)

func EnsureDir(dirPath string) error {
	info, err := os.Stat(dirPath)
	if err == nil && info.IsDir() {
		return nil
	}

	err = os.MkdirAll(dirPath, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}
