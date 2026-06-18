package code

import (
	"fmt"
	"os"
)

// GetPathSize returns the size of a file or directory in bytes as a string.
func GetPathSize(path string) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return fmt.Sprintf("%dB", info.Size()), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var size int64

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return "", err
		}

		size += info.Size()
	}

	return fmt.Sprintf("%dB", size), nil
}
