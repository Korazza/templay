package logger

import (
	"log"
	"os"
)

var (
	Response *log.Logger
	Info     *log.Logger
	Warning  *log.Logger
	Error    *log.Logger
)

func init() {
	Response = log.New(os.Stdout, "", 0)
	Info = log.New(os.Stdout, "Info: ", 0)
	Warning = log.New(os.Stdout, "Warning: ", 0)
	Error = log.New(os.Stderr, "Error: ", 0)
}
