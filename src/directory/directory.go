package directory

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Reads current directory
func ReadDirectory(dirPath string) (fi []fs.DirEntry, dir string, err error) {
	dir, err = filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), dirPath))
	if err != nil {
		return fi, "", fmt.Errorf("failed to get dir path: %v", err)
	}
	// Checks if executed by go run
	if strings.Contains(dir, "go-build") {
		dir = dirPath
	}

	fi, err = os.ReadDir(dir)
	if err != nil {
		return nil, "", fmt.Errorf("failed read dir: %v", err)
	}

	return
}
