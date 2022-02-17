package testutils

import (
	"path/filepath"
	"testing"
)

func TempDir(t *testing.T, path ...string) string {
	return filepath.Join(t.TempDir(), "dfctl-test", t.Name(), filepath.Join(path...))
}
