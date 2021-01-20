package example2


import (
	"github.com/Ferrany1/log2file"
	"log"
)

func ExampleCustomConfig() {
	// Inits new logfile config
	li := log2file.NewOptions()
	// Changes log files names and extension
	li.ChangeConfigNames("log_main", "log_backup", "log")
	// Inits logfile in current dict
	logger, err := li.Logger()
	if err != nil {
		log.Println(err)
	}
	// Writes log into main file
	logger.Println("test")
}
