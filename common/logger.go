package common

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type MyLogger struct {
	zerolog.Logger
}

var Logger MyLogger

func NewLogger() MyLogger {
	// create output configuration
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	// Format level: fatal, error, debug, info, warn
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	// format error
	output.FormatErrFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s: ", i)
	}

	zlog := zerolog.New(output).With().Caller().Timestamp().Logger()
	Logger = MyLogger{zlog}
	return Logger
}

func (l *MyLogger) Info() *zerolog.Event {
	return l.Logger.Info()
}

func (l *MyLogger) Error() *zerolog.Event {
	return l.Logger.Error()
}

func (l *MyLogger) Debug() *zerolog.Event {
	return l.Logger.Debug()
}

func (l *MyLogger) Warn() *zerolog.Event {
	return l.Logger.Warn()
}

func (l *MyLogger) Fatal() *zerolog.Event {
	return l.Logger.Fatal()
}
