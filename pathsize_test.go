package pathsize

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	cases := []struct {
		name      string
		path      string
		recursive bool
		human     bool
		all       bool
		want      string
	}{
		{
			name: "returns size for a single file",
			path: "testdata/test.txt",
			want: "6B",
		},
		{
			name:  "returns human-readable size for a single file",
			path:  "testdata/test.txt",
			human: true,
			want:  "6B",
		},
		{
			name:  "returns human-readable size in KB for a file",
			path:  "testdata/0_YISbBYJg5hkJGcQd.png",
			human: true,
			want:  "22.5KB",
		},
		{
			name: "returns size for a directory",
			path: "testdata",
			want: "293830B",
		},
		{
			name:  "returns human-readable size for a directory",
			path:  "testdata",
			human: true,
			want:  "293.8KB",
		},
		{
			name: "ignores hidden files",
			path: "testdata/hidden",
			want: "0B",
		},
		{
			name: "includes hidden files",
			path: "testdata/hidden",
			all:  true,
			want: "11B",
		},
		{
			name:      "returns recursive size while ignoring hidden files",
			path:      "testdata/recursive",
			recursive: true,
			want:      "55B",
		},
		{
			name:      "returns recursive size including hidden files",
			path:      "testdata/recursive",
			recursive: true,
			all:       true,
			want:      "89B",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Get(c.path, c.recursive, c.human, c.all)
			require.NoError(t, err)
			require.Equal(t, c.want, got)
		})
	}

	// Test error cases
	t.Run("nonexistent path", func(t *testing.T) {
		got, err := Get("unknown", false, false, false)
		require.ErrorIs(t, err, os.ErrNotExist)
		require.Empty(t, got)
	})

	t.Run("unreadable directory", func(t *testing.T) {
		dir := t.TempDir()
		require.NoError(t, os.Chmod(dir, 0o000))

		got, err := Get(dir, false, false, false)
		require.Error(t, err)
		require.Empty(t, got)
	})
}
