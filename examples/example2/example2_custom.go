package example2

import (
	"github.com/Ferrany1/log2file"
	"log"
)

func ExampleCustomOptions() {
	// Gets standard logfile Options
	li := log2file.GetOptions()
	// Changes log files names and extension
	li.ChangeOptionsNames("log_main", "log_backup", "log", 8081)
	// Inits logfile in current dict
	logger, err := li.Logger()
	if err != nil {
		log.Println(err)
	}
	// Writes log into main file
	logger.Println("test")
}
