package pbar

import (
	"strings"

	"github.com/wzshiming/cursor"
)

type ProgressBarGroup []ProgressBar

func (p ProgressBarGroup) Count() int {
	sum := 0
	for _, v := range p {
		sum += v.Count()
	}
	return sum
}

func (p ProgressBarGroup) IsComplete() bool {
	for _, v := range p {
		if !v.IsComplete() {
			return false
		}
	}
	return true
}

func (p ProgressBarGroup) Format() string {
	switch len(p) {
	case 0:
		return ""
	case 1:
		return p[0].Format()
	default:
		ss := []string{}
		ss = append(ss, cursor.RawMoveUp(uint64(len(p))))
		for _, v := range p {
			ss = append(ss, v.Format())
		}
		return strings.Join(ss, "\n")
	}
}
