package pbar

import (
	"strings"
)

// MarkBar basic implementation progress bar mark
type MarkBar struct {
	Kind         string `json:"_Kind"`
	Width        int
	Numer, Denom *MarkInput

	// Padding
	PaddingFiller string
	PaddingLeft   *Marks `json:",omitempty"`
	PaddingMid    *Marks `json:",omitempty"`
	PaddingRight  *Marks `json:",omitempty"`

	// Bar mid
	Mid *Marks

	// Pending
	PendingFiller string
	PendingLeft   *Marks `json:",omitempty"`
	PendingMid    *Marks `json:",omitempty"`
	PendingRight  *Marks `json:",omitempty"`
}

// MarkFormat returns mark string
func (p *MarkBar) String() string {
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

	if p.Mid != nil {
		mid = p.Mid.String()
	}

	if fill := p.PaddingFiller; fill != "" {
		padding = barFormat(fill, p.PaddingLeft, p.PaddingMid, p.PaddingRight, 0, cur, tol)
	}

	if fill := p.PendingFiller; fill != "" {
		pending = barFormat(fill, p.PendingLeft, p.PendingMid, p.PendingRight, cur, tol-cur, tol)
	}

	return strings.Join([]string{padding, mid, pending}, "")
}

// barFormat returns bar string
func barFormat(filler string, left, mid, right *Marks, offset, length, total int) string {
	mask := []rune(filler)
	end := offset + length
	if m := total / len(mask); m > 0 {
		mask = []rune(strings.Repeat(string(mask), m+1))
	}

	if left != nil {
		ss := []rune(left.String())
		copy(mask[:len(ss)], ss)
	}

	if right != nil {
		ss := []rune(right.String())
		copy(mask[len(mask)-len(ss)-1:], ss)
	}

	if mid != nil {
		ss := []rune(mid.String())
		beg := len(mask)/2 - len(ss)/2
		copy(mask[beg:beg+len(ss)], ss)
	}

	return string(mask[offset:end])
}
