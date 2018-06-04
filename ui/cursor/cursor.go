// Documentation:
// http://tldp.org/HOWTO/Bash-Prompt-HOWTO/x361.html
// https://www.student.cs.uwaterloo.ca/~cs452/terminal.html
// https://en.wikipedia.org/wiki/ANSI_escape_code#Unix-like_systems

// Package cursor provides a collection of funtions to manipulate the Bash cursor position
package cursor

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	move              = "[<1>;<2>H"
	up                = "[<1>A"
	down              = "[<1>B"
	forward           = "[<1>C"
	backward          = "[<1>D"
	nextLine          = "[<1>C"
	previousLine      = "[<1>D"
	save              = "7"
	restore           = "8"
	clearScreen       = "[2J"
	clearLineForward  = "[K"
	clearLineBackward = "[1K"
	clearLine         = "[2K"
	scrollUp          = "[<1>S"
	scrollDown        = "[<1>T"
)

// MoveTo moves the cursor to the specified location
func MoveTo(row, col int) string { return escape(move, row, col) }

// MoveVertical moves the cursor up or down
func MoveVertical(n int) string {
	if n > 0 {
		return escape(up, n)
	}
	return escape(down, n)
}

// MoveHorizontal moves the cursor left or right
func MoveHorizontal(n int) string {
	if n > 0 {
		return escape(forward, n)
	}
	return escape(backward, n)
}

// Save remembers the last cursor location
func Save() string { return escape(save) }

// Restore returns the cursor to the last saved location
func Restore() string { return escape(restore) }

// ClearScreen makes the screen blank
func ClearScreen() string { return escape(clearScreen) }

// ClearLineForward deletes everything from the cursor to the end of the line
func ClearLineForward() string { return escape(clearLineForward) }

// ClearLineBackward deletes everything from the cursor to the end of the line
func ClearLineBackward() string { return escape(clearLineBackward) }

// ClearLine deletes everything from the cursor to the end of the line
func ClearLine() string { return escape(clearLine) }

// ScrollUp Scrolls the whole page up
func ScrollUp(n int) string { return escape(scrollUp, n) }

// ScrollDown Scrolls the whole page down
func ScrollDown(n int) string { return escape(scrollDown, n) }

// escape generates the Bash escape sequence to manipulate the cursor
func escape(code string, args ...int) string {
	for i, arg := range args {
		k := fmt.Sprintf("<%v>", i+1)
		v := strconv.Itoa(arg)
		code = strings.Replace(code, k, v, -1)
	}
	return "\033" + code
}
