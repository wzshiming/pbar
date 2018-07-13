package pbar

import (
	"fmt"
	"strings"
	"time"
)

type Marks struct {
	Split string
	Slice []Mark
}

func (m *Marks) MarkFormat(info *BaseInfo) string {
	ss := []string{}
	for _, v := range m.Slice {
		ss = append(ss, v.MarkFormat(info))
	}
	return strings.Join(ss, m.Split)
}

type MarkRatio struct{}

func (b *MarkRatio) MarkFormat(info *BaseInfo) string {
	return fmt.Sprintf("%d/%d", info.Current, info.Total)
}

type MarkPercent struct{}

func (b *MarkPercent) MarkFormat(info *BaseInfo) string {
	per := 0.0
	if info.Total != 0 {
		per = 100 * float64(info.Current) / float64(info.Total)
	}
	return fmt.Sprintf("%6.2f%%", per)
}

type MarkRoll struct {
	Over string
	Roll []string
}

func (b *MarkRoll) MarkFormat(info *BaseInfo) string {
	if info.IsComplete() {
		return b.Over
	}
	i := info.Refresh % len(b.Roll)
	return b.Roll[i]
}

type MarkText struct {
	Text   string
	Width  int
	Roll   int
	Filler string
}

func (b *MarkText) MarkFormat(info *BaseInfo) string {
	text := b.Text
	if b.Width <= 0 {
		return text
	}

	sub := (len(text) - b.Width)
	if sub > 0 {
		if b.Roll > 0 && !info.IsComplete() {
			index := (info.Refresh / b.Roll) % (2 * sub)
			if sub > index {
				text = text[index : index+b.Width]
			} else {
				index -= sub
				text = text[len(text)-index-b.Width : len(text)-index]
			}
		} else {
			text = text[:b.Width-1] + "."
		}
	} else if sub < 0 && b.Filler != "" {
		sub = -sub
		mask := []rune(b.Filler)
		if m := sub / len(mask); m > 0 {
			mask = []rune(strings.Repeat(string(mask), m+1))
		}
		text += string(mask[:sub])
	}
	return text
}

type MarkAfter struct {
	endAt time.Time
}

func (b *MarkAfter) MarkFormat(info *BaseInfo) string {
	if info.StartTime.IsZero() {
		info.StartTime = time.Now()
	}
	if info.IsComplete() {
		if b.endAt.IsZero() {
			b.endAt = time.Now()
		}
		return b.endAt.Sub(info.StartTime).Truncate(time.Second).String()
	}
	return time.Now().Sub(info.StartTime).Truncate(time.Second).String()
}
