package pbar

import (
	"strings"
)

// BarMark basic implementation progress bar mark
type BarMark struct {
	Width          int
	PaddingBarForm Bar
	MidMarkForm    Mark
	PendingBarForm Bar
}

func (p *BarMark) MarkFormat(info *Info) string {

	cur := int(info.Current)
	tol := int(info.Total)
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
