package cursor

import (
	"fmt"
)

// RawSaveCursorPosition is save the cursor position
func RawSaveCursorPosition() string {
	return "\x1b7"
}

// RawRestoreCursorPosition is restore the cursor position
func RawRestoreCursorPosition() string {
	return "\x1b8"
}

// RawShowAlternateScreen is show alternate screen
func RawShowAlternateScreen() string {
	return "\x1b[?1049h"
}

// RawHideAlternateScreen is hide alternate screen
func RawHideAlternateScreen() string {
	return "\x1b[?1049l"
}

// RawClear is clear the screen
func RawClear() string {
	return "\x1b[2J"
}

// RawClearLine is clear from the cursor to the end of the line
func RawClearLine() string {
	return "\x1b[K"
}

// RawMoveUp is cursor move up
func RawMoveUp(x uint64) string {
	return fmt.Sprintf("\x1b[%dA", x)
}

// RawMoveDown is cursor move down
func RawMoveDown(x uint64) string {
	return fmt.Sprintf("\x1b[%dB", x)
}

// RawMoveRight is cursor move right
func RawMoveRight(x uint64) string {
	return fmt.Sprintf("\x1b[%dC", x)
}

// RawMoveLeft is cursor move left
func RawMoveLeft(x uint64) string {
	return fmt.Sprintf("\x1b[%dD", x)
}

// RawMoveTo is cursor move to position
func RawMoveTo(x, y uint64) string {
	return fmt.Sprintf("\x1b[%d;%dH", y, x)
}

// RawHide is cursor hide
func RawHide() string {
	return "\x1b[?25l"
}

// RawShow is cursor show
func RawShow() string {
	return "\x1b[?25h"
}
