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

// Info returns into
func (i *Info) Info() *Info {
	return i
}

// IsComplete returns completed
func (i *Info) IsComplete() bool {
	return i.Total != 0 && i.Current >= i.Total
}

// SetTotal sets total
func (i *Info) SetTotal(total uint64) {
	i.Total = total
}

// SetCurrent sets current
func (i *Info) SetCurrent(current uint64) {
	i.Current = current
}

// AddCurrent adds current
func (i *Info) AddCurrent(val uint64) {
	i.Current += val
}
