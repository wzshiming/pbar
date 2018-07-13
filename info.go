package pbar

import (
	"time"
)

// Info Current information about the progress bar
type Info struct {
	Total     uint64
	Current   uint64
	Refresh   int
	StartTime time.Time
}

func (i *Info) Info() *Info {
	return i
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
