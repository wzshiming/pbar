package pbar

import (
	"fmt"
	"strings"
	"time"
)

type MarkRatio struct{}

func (b MarkRatio) MarkFormat(info *Info) string {
	return fmt.Sprintf("%.1f/%.1f", float64(info.Current), float64(info.Total))
}

type MarkPercent struct{}

func (b MarkPercent) MarkFormat(info *Info) string {
	return fmt.Sprintf("%5.2f%%", 100*float64(info.Current)/float64(info.Total))
}

type MarkRoll string

func (b MarkRoll) MarkFormat(info *Info) string {
	text := string(b)
	i := info.Refresh % len(text)
	return text[i : i+1]
}

type MarkText string

func (b MarkText) MarkFormat(info *Info) string {
	return string(b)
}

type MarkAfter struct{}

func (b MarkAfter) MarkFormat(info *Info) string {
	return time.Now().Sub(info.StartTime).Truncate(time.Second).String()
}

type MarkMerge []Mark

func (b MarkMerge) MarkFormat(info *Info) string {
	ss := []string{}
	for _, v := range b {
		ss = append(ss, v.MarkFormat(info))
	}
	return strings.Join(ss, " ")
}
