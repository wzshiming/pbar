package cursor

import (
	"os"

	_ "github.com/wzshiming/winseq" // Use Unix like Sequences in Windows
)

// Clear is clear the screen
func Clear() {
	os.Stdout.WriteString(RawClear())
}

// ClearLine is is clear from the cursor to the end of the line
func ClearLine() {
	os.Stdout.WriteString(RawClearLine())
}

// MoveUp is cursor move up
func MoveUp(x uint64) {
	os.Stdout.WriteString(RawMoveUp(x))
}

// MoveDown is cursor move down
func MoveDown(x uint64) {
	os.Stdout.WriteString(RawMoveDown(x))
}

// MoveRight is cursor move right
func MoveRight(x uint64) {
	os.Stdout.WriteString(RawMoveRight(x))
}

// MoveLeft is cursor move left
func MoveLeft(x uint64) {
	os.Stdout.WriteString(RawMoveLeft(x))
}

// MoveTo is cursor move to position
func MoveTo(x, y uint64) {
	os.Stdout.WriteString(RawMoveTo(x, y))
}

// Hide is cursor hide
func Hide() {
	os.Stdout.WriteString(RawHide())
}

// Show is cursor show
func Show() {
	os.Stdout.WriteString(RawShow())
}
