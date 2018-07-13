package pbar

import (
	"github.com/wzshiming/ctc"
)

func NewNormalStyle(content string) ProgressBar {
	p := &BaseProgressBar{
		Marks: Marks{
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
		},
	}
	return p
}

func NewTTYShowStyle(content string) ProgressBar {
	p := &BaseProgressBar{
		Marks: Marks{
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
		},
	}
	return p
}

func NewBreakStyle(content string) ProgressBar {
	p := &BaseProgressBar{
		Marks: Marks{
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
		},
	}
	return p
}

func NewAddedStyle(content string) ProgressBar {
	p := &BaseProgressBar{
		Marks: Marks{
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
		},
	}
	return p
}

func NewFillStyle(content string) ProgressBar {
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
	p := &BaseProgressBar{
		Marks: Marks{
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
		},
	}
	return p
}

func NewFillRedStyle(content string) ProgressBar {
	bb := &BaseBar{
		Filler: ` `,
		Left:   &MarkRatio{},
		Right: &MarkRoll{
			Over: ` `,
			Roll: []string{`*`, `*`, ` `, ` `},
		},
	}
	p := &BaseProgressBar{
		Marks: Marks{
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
		},
	}
	return p
}
