package fileloader

import (
	"os"
	"path/filepath"
)

// EnsureDirExist will create it if it does not exist
//
// /foo/file.txt -> /foo
//
// /foo -> /
//
// /foo/ -> /foo
func EnsureDirExist(value string) error {

	dir := filepath.Dir(value)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0700)
	}

	return nil
}

func CheckFile(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	return !info.IsDir() && !os.IsNotExist(err)
}
