package directory

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Reads current directory
func ReadCurrentDirectory() (fi []os.FileInfo, dir string, err error) {
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return fi, "", errors.New(fmt.Sprintf("failed to get dir path. %s", err.Error()))
	}
	// Checks if executed by go run
	if strings.Contains(dir, "go-build") {
		dir = "."
	}

	fi, err = ioutil.ReadDir(dir)
	if err != nil {
		return nil, "", errors.New(fmt.Sprintf("failed read dir. %s", err.Error()))
	}

	return
}