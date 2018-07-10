package pbar

import (
	"time"
)

type Mark interface {
	MarkFormat(info *BaseInfo) string
}

type Bar interface {
	BarFormat(info *BaseInfo, offset, length, total int) string
}

type ProgressBar interface {
	Format() string
	IsComplete() bool
	Count() int
	Info() *BaseInfo
}

type BaseInfo struct {
	Total     uint64
	Current   uint64
	Refresh   int
	StartTime time.Time
}

func (i *BaseInfo) Info() *BaseInfo {
	return i
}

func (i *BaseInfo) IsComplete() bool {
	return i.Total != 0 && i.Current >= i.Total
}

func (i *BaseInfo) SetTotal(total uint64) {
	i.Total = total
}

func (i *BaseInfo) SetCurrent(current uint64) {
	i.Current = current
}

func (i *BaseInfo) AddCurrent(val uint64) {
	i.Current += val
}

func (i *BaseInfo) calculate() {
	if i.Current >= i.Total {
		i.Current = i.Total
	}
	if i.StartTime.IsZero() {
		i.StartTime = time.Now()
	}
}
