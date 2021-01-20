package log2file

import (
	"github.com/Ferrany1/log2file/src/directory"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"time"
)

// Standard names for log files
var (
	standConfig = &fileOptions{fiNames: fileNames{logMain: "log_1", logBackup: "log_2", logExtension: "log"}}
)

// LogFileOptions struct
type fileOptions struct {
	fiNames fileNames
}

// LogFileConfig Names params
type fileNames struct {
	logMain      string
	logBackup    string
	logExtension string
}

// Creates new LogFileConfig element
func NewOptions() *fileOptions {
	return new(fileOptions)
}

// Changes fileNames for LogFileConfig element
func (c *fileOptions) ChangeConfigNames(nMainFile, nBackupFile, nExtensionFile string) *fileOptions {
	c.fiNames = fileNames{logMain: nMainFile, logBackup: nBackupFile, logExtension: nExtensionFile}
	return c
}

// Logger Instance creation that will write new logs to logMain file
func (c *fileOptions) Logger() (*log.Logger, error) {
	if reflect.ValueOf(*c).IsZero() {
		c = standConfig
	}

	lPath, err := c.createLogFile()
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(lPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return log.New(f, "", log.LstdFlags), nil
}

// Creates main log file and renames previous if it was present in directory
func (c *fileOptions) createLogFile() (string, error) {
	fi, dir, err := directory.ReadCurrentDirectory()
	if err != nil {
		return "", err
	}

	var (
		mainPath   = path.Join(dir, c.fiNames.logMain+"."+c.fiNames.logExtension)
		BackupPath = path.Join(dir, c.fiNames.logBackup+"."+c.fiNames.logExtension)
	)

	if c.findLogFile(fi) {
		err = os.Rename(mainPath, BackupPath)
		if err != nil {
			return "", err
		}
	}
	_, err = os.Create(mainPath)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(mainPath, []byte(time.Now().Format(time.RFC3339)+" Started log.\n"), 0644)
	if err != nil {
		return "", err
	}

	return mainPath, nil
}

// Check if mainLog file is in directory
func (c *fileOptions) findLogFile(fi []os.FileInfo) (ok bool) {
	for _, f := range fi {
		if f.Name() == c.fiNames.logMain+"."+c.fiNames.logExtension {
			return true
		}
	}
	return
}
