package pbar

import (
	"bytes"
	"text/template"
)

type BaseProgressBarData struct {
	Padding string
	Pending string
	Head    string
	Mid     string
	Tail    string
}

type BaseProgressBar struct {
	Data BaseProgressBarData
	Info
	Width          int
	Template       *template.Template
	PaddingBarForm Bar
	PendingBarForm Bar
	HeadMarkForm   Mark
	MidMarkForm    Mark
	TailMarkForm   Mark
	Buf            bytes.Buffer
}

func (p *BaseProgressBar) Count() int {
	return 1
}

func (p *BaseProgressBar) Format() string {
	if p.Info.Total == 0 {
		p.Info.Total = 1
	}
	p.Info.calculate()
	p.calculateMark()
	p.calculateBar()
	p.Buf.Reset()
	p.Template.Execute(&p.Buf, p.Data)
	return p.Buf.String()
}

func (p *BaseProgressBar) MarkFormat(info *Info) string {
	p.Info = *info
	return p.Format()
}

func (p *BaseProgressBar) calculateMark() {

	if p.HeadMarkForm != nil {
		p.Data.Head = p.HeadMarkForm.MarkFormat(&p.Info)
	}

	if p.MidMarkForm != nil {
		if p.Info.Current != p.Info.Total {
			p.Data.Mid = p.MidMarkForm.MarkFormat(&p.Info)
		} else {
			p.Data.Mid = ""
		}
	}

	if p.TailMarkForm != nil {
		p.Data.Tail = p.TailMarkForm.MarkFormat(&p.Info)
	}

}
func (p *BaseProgressBar) calculateBar() {

	cur := int(p.Info.Current)
	tol := int(p.Info.Total)

	if p.Width > 0 {
		width := p.Width
		width -= len(p.Data.Head)
		width -= len(p.Data.Tail)
		width -= len(p.Data.Mid)

		cur *= width
		cur /= tol
		tol = width
	}

	sub := tol - cur

	if p.PaddingBarForm != nil {
		p.Data.Padding = p.PaddingBarForm.BarFormat(&p.Info, 0, cur, tol)
	}

	if p.PendingBarForm != nil {
		p.Data.Pending = p.PendingBarForm.BarFormat(&p.Info, cur, sub, tol)
	}

	p.Info.Refresh++
	return
}
