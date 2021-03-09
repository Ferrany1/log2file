package example1

import (
	"github.com/Ferrany1/log2file"
	"log"
)

func ExampleEmptyOptions() {
	// Gets standard logfile Options
	li := log2file.StandardOptions()
	// Inits logfile in current dict
	logger, err := li.NewLogger()
	if err != nil {
		log.Fatalln(err)
	}

	// Writes log into main file

	logger.Println("test")
}
