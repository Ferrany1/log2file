package standartLogger

import (
	"fmt"
	"log"
)

type StandardLogger struct {
	log.Logger
}

func (s *StandardLogger) FormatError(messageText string) (err error) {
	return fmt.Errorf("not implemented for this Logger")
}
