package example1

import (
	"github.com/Ferrany1/log2file"
	"log"
)

func ExampleEmptyConfig() {
	// Inits new logfile config
	li := log2file.NewOptions()
	// Inits logfile in current dict
	logger, err := li.Logger()
	if err != nil {
		log.Println(err)
	}
	// Writes log into main file
	logger.Println("test")
}
