package view_content

import (
	"appconsole/config"
	"appconsole/module"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func WelcomeView(_ fyne.Window) *fyne.Container {
	logo := canvas.NewImageFromResource(fyne.NewStaticResource("", module.NewResourceOP(config.Setting.Resource).GetImage("resources/logo/logo.png")))
	logo.FillMode = canvas.ImageFillContain
	if fyne.CurrentDevice().IsMobile() {
		logo.SetMinSize(fyne.NewSize(192, 192))
	} else {
		logo.SetMinSize(fyne.NewSize(256, 256))
	}

	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to the demo app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		logo,
		container.NewHBox(
			widget.NewHyperlink("google", parseURL("https://developers.google.com/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("documentation", parseURL("https://developers.google.com/learn")),
			widget.NewLabel("-"),
			widget.NewHyperlink("sponsor", parseURL("https://www.google.com/search?q=sponsor&oq=sponsor&aqs=chrome..69i57.3063j0j4&sourceid=chrome&ie=UTF-8")),
		),
		widget.NewLabel(""), // balance the header on the tutorial screen we leave blank on this content
	))
}

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}
