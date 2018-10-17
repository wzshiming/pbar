package pbar

import (
	"strings"
)

// Bar Mark progress is a bar
type Bar interface {
	BarFormat(info *Info, offset, length, total int) string
}

// BaseBar basic implementation bar
type BaseBar struct {
	Filler string
	Left   Mark
	Mid    Mark
	Right  Mark
}

// BarFormat returns bar string
func (b *BaseBar) BarFormat(info *Info, offset, length, total int) string {
	mask := []rune(b.Filler)
	end := offset + length
	if m := total / len(mask); m > 0 {
		mask = []rune(strings.Repeat(string(mask), m+1))
	}

	if b.Left != nil {
		ss := []rune(b.Left.MarkFormat(info))
		copy(mask[:len(ss)], ss)
	}

	if b.Right != nil {
		ss := []rune(b.Right.MarkFormat(info))
		copy(mask[len(mask)-len(ss)-1:], ss)
	}

	if b.Mid != nil {
		ss := []rune(b.Mid.MarkFormat(info))
		beg := len(mask)/2 - len(ss)/2
		copy(mask[beg:beg+len(ss)], ss)
	}

	return string(mask[offset:end])
}
