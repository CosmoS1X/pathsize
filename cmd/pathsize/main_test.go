package main

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	appname := "pathsize"
	path := filepath.Join("..", "..", "testdata", "test.txt")
	directoryPath := filepath.Join("..", "..", "testdata")
	errMsg := "Error:"

	cases := []struct {
		name            string
		args            []string
		exicode         int
		wantStdoutText  string
		wantStderrText  string
		wantStdoutEmpty bool
		wantStderrEmpty bool
	}{
		{
			name:            "prints size and path for existing file",
			args:            []string{appname, path},
			exicode:         0,
			wantStdoutText:  "6B\t" + path + "\n",
			wantStderrEmpty: true,
		},
		{
			name:            "prints size and path for existing directory",
			args:            []string{appname, "-raH", directoryPath},
			exicode:         0,
			wantStdoutText:  "293.9KB\t" + directoryPath + "\n",
			wantStderrEmpty: true,
		},
		{
			name:            "returns error for nonexistent path",
			args:            []string{appname, "unknown"},
			exicode:         1,
			wantStdoutEmpty: true,
			wantStderrText:  errMsg,
		},
		{
			name:            "returns error if too many arguments are provided",
			args:            []string{appname, "path1", "path2"},
			exicode:         1,
			wantStdoutEmpty: true,
			wantStderrText:  errMsg,
		},
		{
			name:            "returns error if no path argument is provided",
			args:            []string{appname},
			exicode:         1,
			wantStdoutEmpty: true,
			wantStderrText:  errMsg,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			code := run(c.args, &stdout, &stderr)

			require.Equal(t, c.exicode, code, "expected exit code")
			if c.wantStdoutEmpty {
				require.Empty(t, stdout.String(), "expected empty stdout")
			} else {
				require.Equal(t, c.wantStdoutText, stdout.String(), "expected stdout")
			}

			if c.wantStderrEmpty {
				require.Empty(t, stderr.String(), "expected empty stderr")
			} else {
				require.Contains(t, stderr.String(), c.wantStderrText, "expected error message in stderr")
			}
		})
	}
}
