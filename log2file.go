package log2file

import (
	"fmt"
	"github.com/Ferrany1/log2file/src/directory"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"os"
	"path"
	"reflect"
	"strconv"
	"time"
)

// Standard names for log files
var (
	standOptions = &fileOptions{fiNames: fileNames{logMain: "log_1", logBackup: "log_2", logExtension: "log"}, router: logRouter{port: 40013}}
)

// LogFileOptions struct
type fileOptions struct {
	fiNames fileNames
	router  logRouter
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

// Changes fileNames for (logMain, logBackup, logExtension string) LogFileOptions element
func (c *fileOptions) ChangeOptionsNames(nMainFile, nBackupFile, nExtensionFile string, port int, routerStatus bool) *fileOptions {
	c.fiNames = fileNames{logMain: nMainFile, logBackup: nBackupFile, logExtension: nExtensionFile}
	c.router = logRouter{port: port, workStatus: routerStatus}
	return c
}

// Logger Instance creation that will write new logs to logMain file
func (c *fileOptions) Logger() (*log.Logger, error) {
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

// Starts a router with paths to logfiles
func router() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	lg := r.Group("/logs")
	lg.GET("/log_m", getLog1)
	lg.GET("/log_b", getLog2)

	log.Println(r.Run(":" + strconv.Itoa(standOptions.router.port)))
}

// Handler for main log file
func getLog1(c *gin.Context) {
	c.File(fmt.Sprintf("./%s.%s", standOptions.fiNames.logMain, standOptions.fiNames.logExtension))
}

// Handler for backup log file
func getLog2(c *gin.Context) {
	c.File(fmt.Sprintf("./%s.%s", standOptions.fiNames.logBackup, standOptions.fiNames.logExtension))
}
