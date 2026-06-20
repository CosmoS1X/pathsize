package code

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize(t *testing.T) {
	t.Run("file", func(t *testing.T) {
		result, err := GetPathSize("testdata/test.txt", false, false)
		require.NoError(t, err)
		require.Equal(t, "6B", result)
	})

	t.Run("file with human-readable size", func(t *testing.T) {
		result, err := GetPathSize("testdata/test.txt", true, false)
		require.NoError(t, err)
		require.Equal(t, "6B", result)
		result, err = GetPathSize("testdata/0_YISbBYJg5hkJGcQd.png", true, false)
		require.NoError(t, err)
		require.Equal(t, "22.5KB", result)
	})

	t.Run("directory", func(t *testing.T) {
		result, err := GetPathSize("testdata", false, false)
		require.NoError(t, err)
		require.Equal(t, "293830B", result)
	})

	t.Run("directory with human-readable size", func(t *testing.T) {
		result, err := GetPathSize("testdata", true, false)
		require.NoError(t, err)
		require.Equal(t, "293.8KB", result)
	})

	t.Run("nonexistent path", func(t *testing.T) {
		result, err := GetPathSize("unknown", false, false)
		require.ErrorIs(t, err, os.ErrNotExist)
		require.Empty(t, result)
	})

	t.Run("unreadable directory", func(t *testing.T) {
		dir := t.TempDir()
		require.NoError(t, os.Chmod(dir, 0o000))

		result, err := GetPathSize(dir, false, false)
		require.Error(t, err)
		require.Empty(t, result)
	})

	t.Run("directory with hidden files", func(t *testing.T) {
		result, err := GetPathSize("testdata/hidden", false, false)
		require.NoError(t, err)
		require.Equal(t, "0B", result)

		result, err = GetPathSize("testdata/hidden", false, true)
		require.NoError(t, err)
		require.Equal(t, "11B", result)
	})
}
