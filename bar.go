package pbar

import (
	"strings"
)

type BarRepeat string

func (b *BarRepeat) BarFormat(info *Info, offset, length, total int) string {
	mask := string(*b)
	end := offset + length
	if m := total / len(mask); m > 0 {
		mask = strings.Repeat(mask, m+1)
		*b = BarRepeat(mask)
	}
	return mask[offset:end]
}

type BarSidesRepeat string

func (b *BarSidesRepeat) BarFormat(info *Info, offset, length, total int) string {
	mask := string(*b)
	end := offset + length
	if m := (total - len(mask)) / 2; m > 0 {
		mask = strings.Repeat(mask[:1], m+1) + mask
		mask = mask + strings.Repeat(mask[len(mask)-1:], m+1)
		*b = BarSidesRepeat(mask)
	}
	return mask[offset:end]
}
