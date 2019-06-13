package log

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
)

var Verbose = false
var logger = log.New(os.Stdout, "[IH] ", 0)

var red = color.New(color.FgRed).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()

func Print(msg interface{}) {
	logger.Print(msg)
}

func Printf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	logger.Print(msg)
}

func Debug(msg interface{}) {
	if Verbose {
		fmt.Print(yellow(msg))
	}
}

func Debugf(format string, v ...interface{}) {
	if Verbose {
		msg := fmt.Sprintf(format, v...)
		fmt.Print(yellow(msg))
	}
}

func Error(msg string) error {
	logger.Print(red(msg))
	return errors.New(msg)
}

func Errorf(format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	logger.Print(red(msg))
	return errors.New(msg)
}
