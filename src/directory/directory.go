package directory

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type LogFileOptions struct {
	FilesPath  FilesPath
	FilesNames FileNames
}

type FilesPath struct {
	MainFilePath   string
	BackupFilePath string
}

// LogFileOptions Names params
type FileNames struct {
	LogMain      string
	LogBackup    string
	LogExtension string
}

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

// Creates main log file and renames previous if it was present in directory
func (o *LogFileOptions) CreateLogFile() (string, error) {
	info, err := os.Stat(o.FilesPath.MainFilePath)
	if err != nil || !info.IsDir() {
		if err := os.MkdirAll(o.FilesPath.MainFilePath, 0755); err != nil {
			log.Fatalf("%v", err)
		}
	}

	fi, _, err := ReadDirectory(o.FilesPath.MainFilePath)
	if err != nil {
		return "", err
	}

	var (
		mainPath   = path.Join(o.FilesPath.MainFilePath, o.FilesNames.LogMain+"."+o.FilesNames.LogExtension)
		BackupPath = path.Join(o.FilesPath.BackupFilePath, o.FilesNames.LogBackup+"."+o.FilesNames.LogExtension)
	)

	if o.findLogFile(fi) {
		err = os.Rename(mainPath, BackupPath)
		if err != nil {
			return "", err
		}
	}

	_, err = os.Create(mainPath)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(mainPath, []byte(time.Now().Format(time.RFC3339)+" Started log.\n"), 0644)
	if err != nil {
		return "", err
	}

	return mainPath, nil
}

// Check if mainLog file is in directory
func (o *LogFileOptions) findLogFile(fi []fs.DirEntry) (ok bool) {
	for _, f := range fi {
		if f.Name() == o.FilesNames.LogMain+"."+o.FilesNames.LogExtension {
			return true
		}
	}
	return
}
