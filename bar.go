package pbar

import (
	"fmt"
	"strings"

	"github.com/wzshiming/ctc"
)

type BaseBar struct {
	Filler string
	Left   Mark
	Mid    Mark
	Right  Mark
}

func (b *BaseBar) BarFormat(info *BaseInfo, offset, length, total int) string {
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

type BarRepeatNegative string

func (b *BarRepeatNegative) BarFormat(info *BaseInfo, offset, length, total int) string {
	mask := string(*b)
	end := offset + length
	if m := total / len(mask); m > 0 {
		mask = strings.Repeat(mask, m+1)
		*b = BarRepeatNegative(mask)
	}
	return fmt.Sprintf("%s%s%s", ctc.Negative, mask[offset:end], ctc.Reset)
}

type BarRepeat string

func (b *BarRepeat) BarFormat(info *BaseInfo, offset, length, total int) string {
	mask := string(*b)
	end := offset + length
	if m := total / len(mask); m > 0 {
		mask = strings.Repeat(mask, m+1)
		*b = BarRepeat(mask)
	}
	return mask[offset:end]
}

type BarSidesRepeatNegative string

func (b *BarSidesRepeatNegative) BarFormat(info *BaseInfo, offset, length, total int) string {
	mask := string(*b)
	end := offset + length
	if m := (total - len(mask)) / 2; m > 0 {
		mask = strings.Repeat(mask[:1], m+1) + mask
		mask = mask + strings.Repeat(mask[len(mask)-1:], m+1)
		*b = BarSidesRepeatNegative(mask)
	}
	return fmt.Sprintf("%s%s%s", ctc.Negative, mask[offset:end], ctc.Reset)
}

type BarSidesRepeat string

func (b *BarSidesRepeat) BarFormat(info *BaseInfo, offset, length, total int) string {
	mask := string(*b)
	end := offset + length
	if m := (total - len(mask)) / 2; m > 0 {
		mask = strings.Repeat(mask[:1], m+1) + mask
		mask = mask + strings.Repeat(mask[len(mask)-1:], m+1)
		*b = BarSidesRepeat(mask)
	}
	return mask[offset:end]
}
