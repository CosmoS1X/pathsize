package code

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize(t *testing.T) {
	t.Run("file", func(t *testing.T) {
		result, err := GetPathSize("testdata/test.txt")
		require.NoError(t, err)
		require.Equal(t, "6B", result)
	})

	t.Run("directory", func(t *testing.T) {
		result, err := GetPathSize("testdata")
		require.NoError(t, err)
		require.Equal(t, "293830B", result)
	})

	t.Run("nonexistent path", func(t *testing.T) {
		result, err := GetPathSize("unknown")
		require.ErrorIs(t, err, os.ErrNotExist)
		require.Empty(t, result)
	})

	t.Run("unreadable directory", func(t *testing.T) {
		dir := t.TempDir()
		require.NoError(t, os.Chmod(dir, 0o000))

		result, err := GetPathSize(dir)
		require.Error(t, err)
		require.Empty(t, result)
	})
}
