package pbar

import (
	"fmt"
	"strings"
	"time"
)

// Marker Marking progress can be any text that represents progress
type Marker interface {
	fmt.Stringer
}

type MarkPerCent struct {
	Kind         string `json:"_Kind"`
	Numer, Denom *MarkInput
}

// MarkPerCent returns mark string
func (b *MarkPerCent) String() string {
	n, _ := b.Numer.Int64()
	d, _ := b.Denom.Int64()

	return fmt.Sprintf("%3d%%", 100*n/d)
}

// MarkPer Show progress in a percent
type MarkPer struct {
	Kind         string `json:"_Kind"`
	Numer, Denom *MarkInput
}

// MarkFormat returns mark string
func (b *MarkPer) String() string {
	n := b.Numer.String()
	d := b.Denom.String()

	r := fmt.Sprintf("%s/%s", n, d)
	s := len(d) - len(n)
	if s > 0 {
		r = strings.Repeat(" ", s) + r
	}
	return r
}

// MarkRoll Indicates that no response has been lost
type MarkRoll struct {
	Kind         string `json:"_Kind"`
	Numer, Denom *MarkInput
	Over         string
	Roll         []string
	Refresh      int `json:",omitempty"`
}

// MarkFormat returns mark string
func (b *MarkRoll) String() string {
	if isComplete(b.Numer, b.Denom) {
		return b.Over
	}
	i := b.Refresh % len(b.Roll)
	b.Refresh++
	return b.Roll[i]
}

// MarkScroll Show text
type MarkScroll struct {
	Kind         string `json:"_Kind"`
	Numer, Denom *MarkInput
	Content      *MarkInput
	Width        int
	Roll         int
	Filler       string
	Refresh      int `json:",omitempty"`
}

// MarkScroll returns mark string
func (b *MarkScroll) String() string {
	text := b.Content.String()
	if b.Width <= 0 {
		return text
	}

	sub := (len(text) - b.Width)
	if sub > 0 {
		if b.Roll > 0 && !isComplete(b.Numer, b.Denom) {
			index := (b.Refresh / b.Roll) % (2 * sub)
			if sub > index {
				text = text[index : index+b.Width]
			} else {
				index -= sub
				text = text[len(text)-index-b.Width : len(text)-index]
			}
		} else {
			text = text[:b.Width-1] + "."
		}
	} else if fill := b.Filler; sub < 0 && fill != "" {
		sub = -sub
		mask := []rune(fill)
		if m := sub / len(mask); m > 0 {
			mask = []rune(strings.Repeat(string(mask), m+1))
		}
		text += string(mask[:sub])
	}
	b.Refresh++
	return text
}

// MarkAfter Show the after time
type MarkAfter struct {
	Kind           string `json:"_Kind"`
	Numer, Denom   *MarkInput
	startAt, endAt time.Time
}

// MarkFormat returns mark string
func (b *MarkAfter) String() string {
	if b.startAt.IsZero() {
		b.startAt = time.Now()
	}

	ret := ""
	if isComplete(b.Numer, b.Denom) {
		if b.endAt.IsZero() {
			b.endAt = time.Now()
		}
		ret = b.endAt.Sub(b.startAt).Truncate(time.Second).String()
	} else {
		ret = time.Now().Sub(b.startAt).Truncate(time.Second).String()
	}

	return fmt.Sprintf("%5s", ret)
}

func isComplete(numer, denom *MarkInput) bool {
	n, _ := numer.Float64()
	d, _ := denom.Float64()
	return d != 0 && n >= d
}
