package log

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
)

var VERBOSE_OUTPUT = false
var logger = log.New(os.Stdout, "[IH] ", 0)

var Green = color.New(color.FgGreen).SprintFunc()
var Red = color.New(color.FgRed).SprintFunc()
var Yellow = color.New(color.FgYellow).SprintFunc()
var Cyan = color.New(color.FgCyan).SprintFunc()

func Positive(msg string) {
	logger.Print(Green(msg))
}

func Print(msg interface{}) {
	logger.Print(msg)
}

func Printf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	logger.Print(msg)
}

func Debug(msg interface{}) {
	if VERBOSE_OUTPUT {
		logger.Print(msg)
	}
}

func Debugf(format string, v ...interface{}) {
	if VERBOSE_OUTPUT {
		msg := fmt.Sprintf(format, v...)
		logger.Print(msg)
	}
}

func Error(msg string) error {
	logger.Print(Red(msg))
	return errors.New(msg)
}

func Errorf(format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	logger.Print(Red(msg))
	return errors.New(msg)
}

func CombinedStd(format string, v ...interface{}) {
	if VERBOSE_OUTPUT {
		msg := fmt.Sprintf(format, v...)
		fmt.Print(Cyan(msg))
	}
}
