package jsonLogger

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

const DatetimeFormatNoSeconds = "2006-01-02 15:04:02"

type logMessage struct {
	Type     string `json:"type"`
	Time     string `json:"time"`
	Message  string `json:"message"`
	Package  string `json:"package"`
	Function string `json:"function"`
}

type JsonLogger struct {
	log.Logger
}

func (l *JsonLogger) FormatError(messageText string) (err error) {
	var message = new(logMessage)
	message.Type = "error"
	message.Time = time.Now().Format(DatetimeFormatNoSeconds)
	message.Message = strings.TrimSpace(messageText)

	programCounter, _, _, _ := runtime.Caller(1)

	return message.formatError(programCounter)
}

func (m *logMessage) formatError(programCounter uintptr) (err error) {
	parts := strings.Split(runtime.FuncForPC(programCounter).Name(), ".")
	partsLen := len(parts)

	m.Function = strings.TrimSpace(parts[partsLen-1])

	if parts[partsLen-2][0] == '(' {
		m.Function = strings.TrimSpace(parts[partsLen-2] + "." + m.Function)
		m.Package = strings.TrimSpace(strings.Join(parts[0:partsLen-2], "."))
	} else {
		m.Package = strings.TrimSpace(strings.Join(parts[0:partsLen-1], "."))
	}

	return fmt.Errorf("%+v", *m)
}
