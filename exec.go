package cursor

import (
	"os"

	_ "github.com/wzshiming/winseq"
)

// Clear is clear screen CLS
func Clear() {
	os.Stdout.WriteString(RawClear())
}

// MoveUp is cursor move up
func MoveUp(x uint8) {
	os.Stdout.WriteString(RawMoveUp(x))
}

// MoveDown is cursor move down
func MoveDown(x uint8) {
	os.Stdout.WriteString(RawMoveDown(x))
}

// MoveRight is cursor move right
func MoveRight(x uint8) {
	os.Stdout.WriteString(RawMoveRight(x))
}

// MoveLeft is cursor move left
func MoveLeft(x uint8) {
	os.Stdout.WriteString(RawMoveLeft(x))
}

// MoveTo is cursor move to position
func MoveTo(x, y uint8) {
	os.Stdout.WriteString(RawMoveTo(x, y))
}

// Hide is cursor hide
func Hide() {
	os.Stdout.WriteString(RawHide())
}

// Hide is cursor show
func Show() {
	os.Stdout.WriteString(RawShow())
}