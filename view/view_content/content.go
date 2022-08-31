package view_content

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var PresentContent *Content

type Content struct {
	Content      *fyne.Container
	ContentTitle *widget.Label
	ContentIntro *widget.Label
}

func NewContent() *Content {
	PresentContent = &Content{}
	PresentContent.Content = container.NewMax()
	PresentContent.ContentTitle = widget.NewLabel("Component name")
	PresentContent.ContentIntro = widget.NewLabel("Intro...")
	PresentContent.ContentIntro.Wrapping = fyne.TextWrapWord
	return PresentContent
}

// 取得畫布
func (c *Content) GetCanvas() *fyne.Container {
	return container.NewBorder(
		container.NewVBox(c.ContentTitle, widget.NewSeparator(), c.ContentIntro), nil, nil, nil, c.Content,
	)
}
