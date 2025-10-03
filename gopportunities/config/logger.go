package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	writer  *io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "DEBUG: ", logger.Flags()),
		warning: log.New(writer, "DEBUG: ", logger.Flags()),
		error:   log.New(writer, "DEBUG: ", logger.Flags()),
		writer:  &writer,
	}
}

func (l Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

func (l Logger) Error(v ...interface{}) {
	l.error.Println(v...)
}

func (l Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l Logger) Errorf(format string, v ...interface{}) {
	l.error.Printf(format, v...)
}
