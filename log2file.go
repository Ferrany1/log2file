package log2file

import (
	"github.com/Ferrany1/log2file/src/directory"
	"io/fs"
	"log"
	"os"
	"path"
	"reflect"
	"time"
)

// Standard names for log files
var standOptions = &fileOptions{
	fiPath:  filesPath{mainFilePath: "./", backupFilePath: "./"},
	fiNames: fileNames{logMain: "log_1", logBackup: "log_2", logExtension: "log"},
	router:  logRouter{port: 40013}}

// LogFileOptions struct
type fileOptions struct {
	fiPath  filesPath
	fiNames fileNames
	router  logRouter
}

type filesPath struct {
	mainFilePath   string
	backupFilePath string
}

// LogFileOptions Names params
type fileNames struct {
	logMain      string
	logBackup    string
	logExtension string
}

type logRouter struct {
	port       int
	workStatus bool
}

// Gets new LogFileOptions element
func GetOptions() *fileOptions {
	return standOptions
}

// Changes filesPath
func (c *fileOptions) ChangeOptionsPath(newMainPath, newBackupPath string) {
	c.fiPath = filesPath{mainFilePath: newMainPath, backupFilePath: newBackupPath}
}

// Changes fileNames for (logMain, logBackup, logExtension string) LogFileOptions element
func (c *fileOptions) ChangeOptionsNames(newMainFileName, newBackupFileName, newExtensionFileName string) {
	c.fiNames = fileNames{logMain: newMainFileName, logBackup: newBackupFileName, logExtension: newExtensionFileName}
}

// Changes router options
func (c *fileOptions) ChangeOptionsRouter(newPort int, updateRouterStatus bool) {
	c.router = logRouter{port: newPort, workStatus: updateRouterStatus}
}

// Logger Instance creation that will write new logs to logMain file
func (c *fileOptions) NewLogger() (*log.Logger, error) {
	if reflect.ValueOf(*c).IsZero() {
		c = standOptions
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

	if c.router.workStatus {
		go router()
	}

	return log.New(f, "", log.LstdFlags), nil
}

// Creates main log file and renames previous if it was present in directory
func (c *fileOptions) createLogFile() (string, error) {
	info, err := os.Stat(c.fiPath.mainFilePath)
	if err != nil || !info.IsDir() {
		if err := os.MkdirAll(c.fiPath.mainFilePath, 0755); err != nil {
			log.Fatalf("%v", err)
		}
	}

	fi, _, err := directory.ReadDirectory(c.fiPath.mainFilePath)
	if err != nil {
		return "", err
	}

	var (
		mainPath   = path.Join(c.fiPath.mainFilePath, c.fiNames.logMain+"."+c.fiNames.logExtension)
		BackupPath = path.Join(c.fiPath.backupFilePath, c.fiNames.logBackup+"."+c.fiNames.logExtension)
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

	err = os.WriteFile(mainPath, []byte(time.Now().Format(time.RFC3339)+" Started log.\n"), 0644)
	if err != nil {
		return "", err
	}

	return mainPath, nil
}

// Check if mainLog file is in directory
func (c *fileOptions) findLogFile(fi []fs.DirEntry) (ok bool) {
	for _, f := range fi {
		if f.Name() == c.fiNames.logMain+"."+c.fiNames.logExtension {
			return true
		}
	}
	return
}
