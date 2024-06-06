package cmd

import (
	"io"
	"os"

	"golang.org/x/term"
)

// terminal stores the UNIX file descriptor and isTerminal check for the tty
type terminal struct {
	fd         uintptr
	isTerminal bool
}

// Output manages the cli output that is user facing with many functionalites
// to manage and control the TUI (Terminal User Interface)
type Output struct {
	terminal
	out io.Writer
}

// NewOutput intialises the output type with output stream as standard out
// and the output stream properties like file descriptor and the output is a terminal
func NewOutput() *Output {
	o := &Output{out: os.Stdout}
	o.fd, o.isTerminal = getTerminalInfo(o.out)

	return o
}

func getTerminalInfo(in io.Writer) (inFd uintptr, isTerminalIn bool) {
	if file, ok := in.(*os.File); ok {
		inFd = file.Fd()
		isTerminalIn = term.IsTerminal(int(inFd))
	}

	return inFd, isTerminalIn
}
