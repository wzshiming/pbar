package pbar

import (
	"github.com/wzshiming/ctc"
)

// NewNormalStyle normal style
func NewNormalStyle(content string) Mark {
	p := &Marks{
		Split: `|`,
		Slice: []Mark{
			&MarkText{
				Filler: ` `,
				Roll:   5,
				Width:  20,
				Text:   content,
			},
			&BarMark{
				Width: 60,
				PaddingBarForm: &BaseBar{
					Filler: `=`,
				},
				MidMarkForm: &MarkRoll{
					Over: `=`,
					Roll: []string{`-`, `\`, `|`, `/`},
				},
				PendingBarForm: &BaseBar{
					Filler: `-`,
				},
			},
			&MarkPercent{},
			&MarkRatio{},
			&MarkAfter{},
		},
	}
	return p
}

// NewTTYShowStyle TTY show style
func NewTTYShowStyle(content string) Mark {
	p := &Marks{
		Split: `|`,
		Slice: []Mark{
			&BarMark{
				Width: len(content),
				PaddingBarForm: &BaseBar{
					Filler: content,
				},
				MidMarkForm: &MarkRoll{
					Over: ` `,
					Roll: []string{`_`, `_`, ` `, ` `},
				},
				PendingBarForm: &BaseBar{
					Filler: ` `,
				},
			},
			&MarkPercent{},
			&MarkRatio{},
			&MarkAfter{},
		},
	}
	return p
}

// NewBreakStyle break style
func NewBreakStyle(content string) Mark {
	p := &Marks{
		Split: `|`,
		Slice: []Mark{
			&MarkText{
				Filler: ` `,
				Roll:   5,
				Width:  20,
				Text:   content,
			},
			&BarMark{
				Width: 60,
				PaddingBarForm: &BaseBar{
					Filler: `>`,
				},
				MidMarkForm: &MarkRoll{
					Over: `>`,
					Roll: []string{`\`, `|`, `/`, `|`},
				},
				PendingBarForm: &BaseBar{
					Filler: `|`,
				},
			},
			&MarkPercent{},
			&MarkRatio{},
			&MarkAfter{},
		},
	}
	return p
}

// NewAddedStyle added style
func NewAddedStyle(content string) Mark {
	p := &Marks{
		Slice: []Mark{
			&MarkText{
				Text: content,
			},
			&MarkText{
				Text: ctc.Negative.String(),
			},
			&BarMark{
				Width: 60,
				PaddingBarForm: &BaseBar{
					Filler: ` `,
				},
				MidMarkForm: &Marks{
					Slice: []Mark{
						&MarkRoll{
							Over: ` `,
							Roll: []string{`-`, `\`, `|`, `/`},
						},
						&MarkText{
							Text: ctc.Reset.String(),
						},
						&Marks{
							Split: "|",
							Slice: []Mark{
								&MarkPercent{},
								&MarkRatio{},
								&MarkAfter{},
							},
						},
					},
				},
			},
		},
	}
	return p
}

// NewFillStyle fill style
func NewFillStyle(content string) Mark {
	bb := &BaseBar{
		Filler: ` `,
		Left: &MarkText{
			Text: content,
		},
		Mid: &Marks{
			Split: ` | `,
			Slice: []Mark{
				&MarkPercent{},
				&MarkRatio{},
			},
		},
		Right: &Marks{
			Slice: []Mark{
				&MarkAfter{},
				&MarkRoll{
					Over: ` `,
					Roll: []string{`-`, `\`, `|`, `/`},
				},
			},
		},
	}
	p := &Marks{
		Slice: []Mark{
			&MarkText{
				Text: ctc.Negative.String(),
			},
			&BarMark{
				Width:          80,
				PaddingBarForm: bb,
				MidMarkForm: &MarkText{
					Text: ctc.Reset.String(),
				},
				PendingBarForm: bb,
			},
		},
	}
	return p
}

// NewFillRedStyle fill red style
func NewFillRedStyle(content string) Mark {
	bb := &BaseBar{
		Filler: ` `,
		Left:   &MarkRatio{},
		Right: &MarkRoll{
			Over: ` `,
			Roll: []string{`*`, `*`, ` `, ` `},
		},
	}
	p := &Marks{
		Split: ``,
		Slice: []Mark{
			&MarkText{
				Text: content,
			},
			&MarkText{
				Text: ctc.BackgroundRed.String(),
			},
			&BarMark{
				Width:          60,
				PaddingBarForm: bb,
				MidMarkForm: &MarkText{
					Text: ctc.Reset.String(),
				},
				PendingBarForm: bb,
			},
			&Marks{
				Split: ` `,
				Slice: []Mark{
					&MarkPercent{},
					&MarkAfter{},
				},
			},
		},
	}
	return p
}
