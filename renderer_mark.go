package pbar

import (
	"strings"

	"github.com/wzshiming/cursor"
)

type infoMark struct {
	Info
	Mark Mark
}

func (i *infoMark) String() string {
	return i.Mark.MarkFormat(&i.Info)
}

type rendererMark []*infoMark

func (p rendererMark) Count() int {
	return len(p)
}

func (p rendererMark) IsComplete() bool {
	for _, v := range p {
		if !v.Info.IsComplete() {
			return false
		}
	}
	return true
}

func (p rendererMark) String() string {
	switch len(p) {
	case 0:
		return ""
	default:
		ss := []string{}
		ss = append(ss, cursor.RawMoveUp(uint64(len(p))))
		for _, v := range p {
			ss = append(ss, v.String())
		}
		return strings.Join(ss, "\n")
	}
}
