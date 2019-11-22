package console

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	PRE_SPACE      = "  "
	ENABLE_VERBOSE = false
	SprintRed      = color.New(color.FgRed).SprintFunc()
	SprintBlue     = color.New(color.FgBlue).SprintFunc()
	SprintYellow   = color.New(color.FgYellow).SprintFunc()
)

func Print(message interface{}) {
	fmt.Printf("%v\n", message)
}

func Printf(format string, messages ...interface{}) {
	message := fmt.Sprintf(format, messages...)
	fmt.Printf("%v\n", message)
}

func Verbose(message interface{}) {
	if ENABLE_VERBOSE {
		fmt.Printf("%v\n", SprintYellow(message))
	}
}

func Verbosef(format string, messages ...interface{}) {
	if ENABLE_VERBOSE {
		message := fmt.Sprintf(format, messages...)
		fmt.Printf("%v\n", SprintYellow(message))
	}
}

func Error(message string) {
	fmt.Printf("%v\n", SprintRed(message))
}

func Errorf(format string, messages ...interface{}) {
	message := fmt.Sprintf(format, messages...)
	fmt.Printf("%v\n", SprintRed(message))
}

func AddLine() {
	Print(SprintBlue("╺━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╸"))
}

func AddTable(messages []string) {
	Print(SprintBlue("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓"))
	for _, message := range messages {
		Printf("%s  %-83s%s", SprintBlue("┃"), message, SprintBlue("┃"))
	}
	Print(SprintBlue("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛"))
}
