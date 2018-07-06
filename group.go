package pbar

import (
	"strings"

	"github.com/wzshiming/cursor"
)

type ProgressBarGroup []ProgressBar

func (p ProgressBarGroup) IsComplete() bool {
	for _, v := range p {
		if !v.IsComplete() {
			return false
		}
	}
	return true
}

func (p ProgressBarGroup) Format() string {
	ss := []string{}
	for _, v := range p {
		ss = append(ss, v.Format())
	}
	if len(ss) > 1 {
		ss = append(ss, cursor.RawMoveUp(uint64(len(ss))))
	}
	return strings.Join(ss, "\n")
}

func (p ProgressBarGroup) End() string {
	return cursor.RawMoveDown(uint64(len(p)))
}
