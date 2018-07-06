package pbar

import (
	"time"
)

type Info struct {
	Total     uint64
	Current   uint64
	Refresh   int
	StartTime time.Time
}

func (i *Info) IsComplete() bool {
	return i.Total != 0 && i.Current >= i.Total
}

func (i *Info) SetTotal(total uint64) {
	i.Total = total
}

func (i *Info) SetCurrent(current uint64) {
	i.Current = current
}

func (i *Info) AddCurrent(val uint64) {
	i.Current += val
}

func (i *Info) calculate() {
	if i.Current >= i.Total {
		i.Current = i.Total
	}
	if i.StartTime.IsZero() {
		i.StartTime = time.Now()
	}
}

type Mark interface {
	MarkFormat(info *Info) string
}

type Bar interface {
	BarFormat(info *Info, offset, length, total int) string
}

type ProgressBar interface {
	Format() string
	IsComplete() bool
}
