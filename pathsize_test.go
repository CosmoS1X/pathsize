package pathsize

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testdataDir  = "testdata"
	testFile     = "test.txt"
	pngFile      = "0_YISbBYJg5hkJGcQd.png"
	hiddenDir    = "hidden"
	recursiveDir = "recursive"
)

func TestGet(t *testing.T) {
	cases := []struct {
		name      string
		path      string
		recursive bool
		human     bool
		all       bool
		want      string
		wantErr   bool
	}{
		{
			name: "returns size for a single file",
			path: filepath.Join(testdataDir, testFile),
			want: "6B",
		},
		{
			name:  "returns human-readable size for a single file",
			path:  filepath.Join(testdataDir, testFile),
			human: true,
			want:  "6B",
		},
		{
			name:  "returns human-readable size in KB for a file",
			path:  filepath.Join(testdataDir, pngFile),
			human: true,
			want:  "22.5KB",
		},
		{
			name: "returns size for a directory",
			path: testdataDir,
			want: "293830B",
		},
		{
			name:  "returns human-readable size for a directory",
			path:  testdataDir,
			human: true,
			want:  "293.8KB",
		},
		{
			name: "ignores hidden files",
			path: filepath.Join(testdataDir, hiddenDir),
			want: "0B",
		},
		{
			name: "includes hidden files",
			path: filepath.Join(testdataDir, hiddenDir),
			all:  true,
			want: "11B",
		},
		{
			name:      "returns recursive size while ignoring hidden files",
			path:      filepath.Join(testdataDir, recursiveDir),
			recursive: true,
			want:      "55B",
		},
		{
			name:      "returns recursive size including hidden files",
			path:      filepath.Join(testdataDir, recursiveDir),
			recursive: true,
			all:       true,
			want:      "89B",
		},
		{
			name:    "nonexistent path",
			path:    "unknown",
			wantErr: true,
		},
		{
			name: "unreadable directory",
			path: func() string {
				dir := t.TempDir()
				require.NoError(t, os.Chmod(dir, 0o000))
				return dir
			}(),
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got, err := Get(c.path, c.recursive, c.human, c.all)
			if c.wantErr {
				require.Error(t, err)
				require.Empty(t, got)
				return
			}

			require.NoError(t, err)
			require.Equal(t, c.want, got)
		})
	}
}
