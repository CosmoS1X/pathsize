package code

import (
	"fmt"
	"os"
)

const threshold float64 = 1000

// fmtHuman formats the size in bytes to a human-readable string.
func fmtHuman(size float64, human bool) string {
	if !human {
		return fmt.Sprintf("%.0fB", size)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	unitIdx := 0

	for unitIdx < len(units)-1 && size >= threshold {
		size /= threshold
		unitIdx++
	}

	if unitIdx == 0 {
		return fmt.Sprintf("%.0f%s", size, units[unitIdx])
	}

	return fmt.Sprintf("%.1f%s", size, units[unitIdx])
}

// GetPathSize returns the size of a file or directory in bytes as a string.
func GetPathSize(path string, human bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		size := float64(info.Size())
		return fmtHuman(size, human), nil
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

	return fmtHuman(float64(size), human), nil
}
