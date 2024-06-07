package terminal

import (
	"fmt"
	"strings"
)

const (
	// ESC Escape character.
	ESC = '\x1b'
	// CSI Control Sequence Introducer.
	CSI = string(ESC) + "["
	// OSC Operating System Command.
	OSC = string(ESC) + "]"
)

// Sequence definitions.
const (
	// Cursor positioning
	CursorUpSeq              = "%dA"
	CursorDownSeq            = "%dB"
	CursorForwardSeq         = "%dC"
	CursorBackSeq            = "%dD"
	CursorNextLineSeq        = "%dE"
	CursorPreviousLineSeq    = "%dF"
	CursorPositionSeq        = "%d;%dH"
	EraseDisplaySeq          = "%dJ"
	EraseLineSeq             = "%dK"
	SaveCursorPositionSeq    = "s"
	RestoreCursorPositionSeq = "u"
	ChangeScrollingRegionSeq = "%d;%dr"
	InsertLineSeq            = "%dL"
	DeleteLineSeq            = "%dM"

	// Explicit values ersasing lines
	EraseLineRightSeq  = "0K"
	EraseLineLeftSeq   = "1K"
	EraseEntireLineSeq = "2K"

	// Screen
	RestoreScreenSeq = "?47l"
	SaveScreenSeq    = "?47h"
	AltScreenSeq     = "?1049h"
	ExitAltScreenSeq = "?1049l"

	// Bracketed paste.
	// https://en.wikipedia.org/wiki/Bracketed-paste
	EnableBracketedPasteSeq  = "?2004h"
	DisableBracketedPasteSeq = "?2004l"

	SetWindowTitleSeq = "2;%s"
	ShowCursorSeq     = "?25h"
	HideCursorSeq     = "?25l"
)

const (
	moveCursorUp = iota + 1
	clearScreen
)

// Reset the terminal to its default style, removing any active styles.
func (o *Output) Reset() {
	fmt.Fprint(o.out, CSI+"0"+"m")
}

// RestoreScreen restores a previously saved screen state.
func (o *Output) RestoreScreen() {
	fmt.Fprint(o.out, CSI+RestoreScreenSeq)
}

// SaveScreen saves the screen state.
func (o *Output) SaveScreen() {
	fmt.Fprint(o.out, CSI+SaveScreenSeq)
}

// AltScreen switches to the alternate screen buffer. The former view can be
// restored with ExitAltScreen().
func (o *Output) AltScreen() {
	fmt.Fprint(o.out, CSI+AltScreenSeq)
}

// ExitAltScreen exits the alternate screen buffer and returns to the former
// terminal view.
func (o *Output) ExitAltScreen() {
	fmt.Fprint(o.out, CSI+ExitAltScreenSeq)
}

// ClearScreen clears the visible portion of the terminal.
func (o *Output) ClearScreen() {
	fmt.Fprintf(o.out, CSI+EraseDisplaySeq, clearScreen)
	o.MoveCursor(1, 1)
}

// MoveCursor moves the cursor to a given position.
func (o *Output) MoveCursor(row, column int) {
	fmt.Fprintf(o.out, CSI+CursorPositionSeq, row, column)
}

// HideCursor hides the cursor.
func (o *Output) HideCursor() {
	fmt.Fprint(o.out, CSI+HideCursorSeq)
}

// ShowCursor shows the cursor.
func (o *Output) ShowCursor() {
	fmt.Fprint(o.out, CSI+ShowCursorSeq)
}

// SaveCursorPosition saves the cursor position.
func (o *Output) SaveCursorPosition() {
	fmt.Fprint(o.out, CSI+SaveCursorPositionSeq)
}

// RestoreCursorPosition restores a saved cursor position.
func (o *Output) RestoreCursorPosition() {
	fmt.Fprint(o.out, CSI+RestoreCursorPositionSeq)
}

// CursorUp moves the cursor up a given number of lines.
func (o *Output) CursorUp(n int) {
	fmt.Fprintf(o.out, CSI+CursorUpSeq, n)
}

// CursorDown moves the cursor down a given number of lines.
func (o *Output) CursorDown(n int) {
	fmt.Fprintf(o.out, CSI+CursorDownSeq, n)
}

// CursorForward moves the cursor up a given number of lines.
func (o *Output) CursorForward(n int) {
	fmt.Fprintf(o.out, CSI+CursorForwardSeq, n)
}

// CursorBack moves the cursor backwards a given number of cells.
func (o *Output) CursorBack(n int) {
	fmt.Fprintf(o.out, CSI+CursorBackSeq, n)
}

// CursorNextLine moves the cursor down a given number of lines and places it at
// the beginning of the line.
func (o *Output) CursorNextLine(n int) {
	fmt.Fprintf(o.out, CSI+CursorNextLineSeq, n)
}

// CursorPrevLine moves the cursor up a given number of lines and places it at
// the beginning of the line.
func (o *Output) CursorPrevLine(n int) {
	fmt.Fprintf(o.out, CSI+CursorPreviousLineSeq, n)
}

// ClearLine clears the current line.
func (o *Output) ClearLine() {
	fmt.Fprint(o.out, CSI+EraseEntireLineSeq)
}

// ClearLineLeft clears the line to the left of the cursor.
func (o *Output) ClearLineLeft() {
	fmt.Fprint(o.out, CSI+EraseLineLeftSeq)
}

// ClearLineRight clears the line to the right of the cursor.
func (o *Output) ClearLineRight() {
	fmt.Fprint(o.out, CSI+EraseLineRightSeq)
}

// ClearLines clears a given number of lines.
func (o *Output) ClearLines(n int) {
	clearLine := fmt.Sprintf(CSI+EraseLineSeq, clearScreen)
	cursorUp := fmt.Sprintf(CSI+CursorUpSeq, moveCursorUp)

	fmt.Fprint(o.out, clearLine+strings.Repeat(cursorUp+clearLine, n))
}

// ChangeScrollingRegion sets the scrolling region of the terminal.
func (o *Output) ChangeScrollingRegion(top, bottom int) {
	fmt.Fprintf(o.out, CSI+ChangeScrollingRegionSeq, top, bottom)
}

// InsertLines inserts the given number of lines at the top of the scrollable
// region, pushing lines below down.
func (o *Output) InsertLines(n int) {
	fmt.Fprintf(o.out, CSI+InsertLineSeq, n)
}

// DeleteLines deletes the given number of lines, pulling any lines in
// the scrollable region below up.
func (o *Output) DeleteLines(n int) {
	fmt.Fprintf(o.out, CSI+DeleteLineSeq, n)
}

// SetWindowTitle sets the terminal window title.
func (o *Output) SetWindowTitle(title string) {
	fmt.Fprintf(o.out, OSC+SetWindowTitleSeq, title)
}

// EnableBracketedPaste enables bracketed paste.
func (o *Output) EnableBracketedPaste() {
	fmt.Fprintf(o.out, CSI+EnableBracketedPasteSeq)
}

// DisableBracketedPaste disables bracketed paste.
func (o *Output) DisableBracketedPaste() {
	fmt.Fprintf(o.out, CSI+DisableBracketedPasteSeq)
}
