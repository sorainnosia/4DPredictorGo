package main

import (
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type Pagination struct {
	PageIndex   int
	PageSize    int
	PageCount   int
	Data        []string
	txtText     *widget.Entry
	btnPrevious *widget.Button
	btnNext     *widget.Button
	lblLabel    *widget.Label
	Container   *fyne.Container
}

func NewPagination() *Pagination {
	p := &Pagination{}
	p.PageSize = 10
	p.btnNext = widget.NewButtonWithIcon("Next", theme.NavigateNextIcon(), func() {
		p.Next()
	})
	p.btnPrevious = widget.NewButtonWithIcon("Previous", theme.NavigateBackIcon(), func() {
		p.Previous()
	})
	p.lblLabel = widget.NewLabel("")
	p.Container = fyne.NewContainerWithLayout(layout.NewBorderLayout(layout.NewSpacer(), layout.NewSpacer(), p.btnPrevious, p.btnNext), p.btnPrevious, p.btnNext, p.lblLabel)
	p.Refresh()
	return p
}

func (p *Pagination) Refresh() {
	p.PageCount = int(len(p.Data) / p.PageSize)
	if len(p.Data)%p.PageSize > 0 {
		p.PageCount++
	}

	if p.txtText != nil {
		p.txtText.SetText(strings.Join(p.Current(), "\n"))
	}

	c := &Convert{}
	p.lblLabel.SetText(c.IntToString(p.PageIndex+1) + "/" + c.IntToString(p.PageCount))
	p.btnNext.Enable()
	p.btnPrevious.Enable()
	if p.PageIndex == 0 {
		p.btnPrevious.Disable()
	}
	if p.PageIndex >= p.PageCount-1 {
		p.btnNext.Disable()
	}
}

func (p *Pagination) Current() []string {
	if len(p.Data) == 0 {
		return nil
	}
	if p.PageIndex > p.PageCount-1 {
		p.PageIndex = 0
	}
	start := p.PageIndex * p.PageSize
	end := start + p.PageSize

	if end >= len(p.Data) {
		end = len(p.Data)
	}
	return p.Data[start:end]
}

func (p *Pagination) Previous() {
	if p.PageIndex-1 >= 0 {
		p.PageIndex--
	}
	p.Refresh()
}

func (p *Pagination) Next() {
	if p.PageIndex+1 <= p.PageCount-1 {
		p.PageIndex++
	}
	p.Refresh()
}
