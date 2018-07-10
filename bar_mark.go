package pbar

import (
	"strings"
)

type BarMark struct {
	Width          int
	PaddingBarForm Bar
	MidMarkForm    Mark
	PendingBarForm Bar
}

func (p *BarMark) MarkFormat(info *BaseInfo) string {

	cur := int(info.Current)
	tol := int(info.Total)
	if cur > tol {
		cur = tol
	}

	if p.Width > 0 && cur != 0 {
		width := p.Width
		cur *= width
		cur /= tol
		tol = width
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
