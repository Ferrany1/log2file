package example1

import (
	"github.com/Ferrany1/log2file"
	"log"
)

func ExampleEmptyOptions() {
	// Gets standard logfile Options
	li := log2file.GetOptions()
	// Inits logfile in current dict
	logger, err := li.Logger()
	if err != nil {
		log.Println(err)
	}
	// Writes log into main file
	logger.Println("test")
}
