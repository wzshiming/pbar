package pbar

import (
	"github.com/wzshiming/ctc"
)

func NewBaseProgressBar(content string) ProgressBar {
	return NewNormalStyle(content)
}

func NewNormalStyle(content string) ProgressBar {
	p := &BaseProgressBar{
		Marks: Marks{
			Split: `|`,
			Slice: []Mark{
				&MarkText{
					Max:  30,
					Text: content,
				},
				&BarMark{
					Width: 80,
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

func NewPercentStyle(content string) ProgressBar {
	p := &BaseProgressBar{
		Marks: Marks{
			Slice: []Mark{
				&MarkText{
					Max:  30,
					Text: content,
				},
				&MarkText{
					Text: ctc.Negative.String(),
				},
				&BarMark{
					Width: 80,
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
			Max:  30,
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
					Width:          100,
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

func NewBreakStyle(content string) ProgressBar {
	p := &BaseProgressBar{
		Marks: Marks{
			Split: `|`,
			Slice: []Mark{
				&MarkText{
					Max:  30,
					Text: content,
				},
				&BarMark{
					Width: 80,
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
