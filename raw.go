package cursor

import (
	"fmt"
)

// RawClear is clear screen CLS
func RawClear() string {
	return "\x1b[2J"
}

// RawMoveUp is cursor move up
func RawMoveUp(x uint8) string {
	return fmt.Sprintf("\x1b[%dA", x)
}

// RawMoveDown is cursor move down
func RawMoveDown(x uint8) string {
	return fmt.Sprintf("\x1b[%dB", x)
}

// RawMoveRight is cursor move right
func RawMoveRight(x uint8) string {
	return fmt.Sprintf("\x1b[%dC", x)
}

// RawMoveLeft is cursor move left
func RawMoveLeft(x uint8) string {
	return fmt.Sprintf("\x1b[%dD", x)
}

// RawMoveTo is cursor move to position
func RawMoveTo(x, y uint8) string {
	return fmt.Sprintf("\x1b[%d;%dH", x, y)
}

// RawHide is cursor hide
func RawHide() string {
	return "\x1b[?25l"
}

// RawHide is cursor show
func RawShow() string {
	return "\x1b[?25h"
}