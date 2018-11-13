package pbar

import (
	"strings"
)

// MarkBar basic implementation progress bar mark
type MarkBar struct {
	Width          int
	Numer, Denom   *MarkInput
	PaddingBarForm *Bar
	MidMarkForm    *Marks
	PendingBarForm *Bar
}

// MarkFormat returns mark string
func (p *MarkBar) MarkFormat(info *Info) string {
	numer, _ := p.Numer.Int64()
	denom, _ := p.Denom.Int64()
	cur := int(numer)
	tol := int(denom)
	if cur > tol {
		cur = tol
	}

	if p.Width > 0 {
		width := p.Width
		if tol != 0 {
			cur *= width
			cur /= tol
			tol = width
		} else {
			tol = width
			cur = 0
		}
	}

	mid := ""
	padding := ""
	pending := ""

	if p.MidMarkForm != nil {
		mid = p.MidMarkForm.MarkFormat(info)
	}

	if p.PaddingBarForm != nil {
		padding = p.PaddingBarForm.BarFormat(info, 0, cur, tol)
	}

	if p.PendingBarForm != nil {
		pending = p.PendingBarForm.BarFormat(info, cur, tol-cur, tol)
	}

	info.Refresh++
	return strings.Join([]string{padding, mid, pending}, "")
}

// Bar basic implementation bar
type Bar struct {
	Filler string
	Left   *Marks
	Mid    *Marks
	Right  *Marks
}

// BarFormat returns bar string
func (b *Bar) BarFormat(info *Info, offset, length, total int) string {
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
