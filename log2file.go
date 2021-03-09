package log2file

import (
	"github.com/Ferrany1/log2file/src/directory"
	"github.com/Ferrany1/log2file/src/logger"
	"github.com/Ferrany1/log2file/src/logger/jsonLogger"
	"github.com/Ferrany1/log2file/src/logger/standartLogger"
	"os"
)

// Standard names for log files
var standOptions = &options{
	fileOptions: directory.LogFileOptions{
		FilesPath:  directory.FilesPath{MainFilePath: "./", BackupFilePath: "./"},
		FilesNames: directory.FileNames{LogMain: "log_1", LogBackup: "log_2", LogExtension: "log"},
	},
	jsonFormat: false,
}

// LogFileOptions struct
type options struct {
	fileOptions directory.LogFileOptions
	jsonFormat  bool
}

// Returns new LogFileOptions element
func StandardOptions() *options {
	return standOptions
}

// Changes filesPath
func (o *options) ChangeOptionsPath(newMainPath, newBackupPath string) {
	o.fileOptions.FilesPath = directory.FilesPath{MainFilePath: newMainPath, BackupFilePath: newBackupPath}
}

// Changes fileNames for (logMain, logBackup, logExtension string) LogFileOptions element
func (o *options) ChangeOptionsNames(newMainFileName, newBackupFileName, newExtensionFileName string) {
	o.fileOptions.FilesNames = directory.FileNames{LogMain: newMainFileName, LogBackup: newBackupFileName, LogExtension: newExtensionFileName}
}

// Changes jsonFormat of logging
func (o *options) ChangeJsonFormat(jsonFormat bool) {
	o.jsonFormat = jsonFormat
}

// Logger Instance creation that will write new logs to logMain file
func (o *options) NewLogger() (logger.FileLogger, error) {
	lPath, err := o.fileOptions.CreateLogFile()
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(lPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	if o.jsonFormat {
		l := new(jsonLogger.JsonLogger)
		l.SetOutput(f)
		return l, nil
	} else {
		l := new(standartLogger.StandardLogger)
		l.SetOutput(f)
		return l, nil
	}
}
