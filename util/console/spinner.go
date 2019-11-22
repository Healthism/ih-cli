package console

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

type Loading struct {
	s *spinner.Spinner
}

const (
	SPINNER_PROGRESS_COLOR = "yellow"
	SPINNER_PREFIX_COLOR   = color.FgGreen
)

func ShowLoading(message string, prefix string) Loading {
	if !ENABLE_VERBOSE {
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)

		if prefix != "" {
			s.Prefix = color.New(SPINNER_PREFIX_COLOR).SprintFunc()(PRE_SPACE + prefix + PRE_SPACE)
		}
		s.Suffix = PRE_SPACE + message

		s.Color(SPINNER_PROGRESS_COLOR)
		s.Start()
		return Loading{s}
	}

	return Loading{}
}

func (l *Loading) HideLoading(err error) {
	if !ENABLE_VERBOSE {
		if err != nil {
			l.s.FinalMSG = fmt.Sprintf("%s☠︎%s\n", l.s.Prefix, color.New(color.FgRed).SprintFunc()(l.s.Suffix))
		} else {
			l.s.FinalMSG = fmt.Sprintf("%s⚐%s\n", l.s.Prefix, l.s.Suffix)
		}

		l.s.Stop()
	}
}
