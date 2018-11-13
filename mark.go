package pbar

import (
	"fmt"
	"strings"
	"time"
)

// Marker Marking progress can be any text that represents progress
type Marker interface {
	MarkFormat(info *Info) string
}

type MarkPerCent struct {
	Numer, Denom *MarkInput
}

// MarkPerCent returns mark string
func (b *MarkPerCent) MarkFormat(info *Info) string {
	n, _ := b.Numer.Float64()
	d, _ := b.Denom.Float64()
	return fmt.Sprintf("%3.0f%%", n/d*100)
}

// MarkPer Show progress in a percent
type MarkPer struct {
	Numer, Denom *MarkInput
}

// MarkFormat returns mark string
func (b *MarkPer) MarkFormat(info *Info) string {
	n := b.Numer.String()
	d := b.Denom.String()
	return strings.Repeat(" ", len(d)-len(n)) + fmt.Sprintf("%s/%s", n, d)
}

// MarkRoll Indicates that no response has been lost
type MarkRoll struct {
	Numer, Denom *MarkInput
	Over         string
	Roll         []string
}

// MarkFormat returns mark string
func (b *MarkRoll) MarkFormat(info *Info) string {
	if isComplete(b.Numer, b.Denom) {
		return b.Over
	}
	i := info.Refresh % len(b.Roll)
	return b.Roll[i]
}

// MarkText Show text
type MarkText struct {
	Numer, Denom *MarkInput
	Text         *MarkInput
	Width        int
	Roll         int
	Filler       string
}

// MarkFormat returns mark string
func (b *MarkText) MarkFormat(info *Info) string {
	text := b.Text.String()
	if b.Width <= 0 {
		return text
	}

	sub := (len(text) - b.Width)
	if sub > 0 {
		if b.Roll > 0 && !isComplete(b.Numer, b.Denom) {
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

// MarkAfter Show the after time
type MarkAfter struct {
	Numer, Denom *MarkInput
	endAt        time.Time
}

// MarkFormat returns mark string
func (b *MarkAfter) MarkFormat(info *Info) string {
	if info.StartTime.IsZero() {
		info.StartTime = time.Now()
	}

	ret := ""
	if isComplete(b.Numer, b.Denom) {
		if b.endAt.IsZero() {
			b.endAt = time.Now()
		}
		ret = b.endAt.Sub(info.StartTime).Truncate(time.Second).String()
	} else {
		ret = time.Now().Sub(info.StartTime).Truncate(time.Second).String()
	}

	return fmt.Sprintf("%5s", ret)
}

func isComplete(numer, denom *MarkInput) bool {
	n, _ := numer.Float64()
	d, _ := denom.Float64()
	return d != 0 && n >= d
}
