package example2

import (
	"github.com/Ferrany1/log2file"
	"log"
)

func ExampleCustomOptions() {
	// Gets standard logfile Options
	li := log2file.StandardOptions()
	// Changes log files names and extension
	li.ChangeOptionsNames("log_main", "log_backup", "log")
	li.ChangeOptionsPath("./testDir", "./testDir")
	li.ChangeJsonFormat(true)
	// Inits logfile in current dict
	logger, err := li.NewLogger()
	if err != nil {
		log.Fatalln(err)
	}
	// Writes log into main file

	logger.Println(logger.FormatError("test"))
}
