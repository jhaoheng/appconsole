package view_content

import (
	"appconsole/config"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MetadataView(win fyne.Window) *fyne.Container {
	metadata := NewMetaData()
	metadata.makeui()
	return metadata.UI
}

type Metadata struct {
	UI *fyne.Container
	//
	AppID       string
	Version     string
	Build       int
	UniqueID    string
	CommitCode  string
	Licence     string
	StoragePath string
}

func NewMetaData() *Metadata {
	return &Metadata{
		AppID:       fyne.CurrentApp().Metadata().ID,
		Version:     fyne.CurrentApp().Metadata().Version,
		Build:       fyne.CurrentApp().Metadata().Build,
		UniqueID:    fyne.CurrentApp().UniqueID(),
		CommitCode:  config.Setting.CommitCode,
		Licence:     "",
		StoragePath: config.Setting.DefaultStoratePath,
	}
}

func (obj *Metadata) makeui() {

	block := func(size fyne.Size) *canvas.Raster {
		tmp := canvas.NewRasterWithPixels(func(x int, y int, w int, h int) color.Color {
			return theme.BackgroundColor()
		})
		tmp.SetMinSize(size)
		return tmp
	}

	data := container.NewBorder(
		block(fyne.NewSize(0, 50)), nil, block(fyne.NewSize(50, 0)), nil,
		obj.SetFormView(),
	)
	data.Resize(fyne.NewSize(800, 50))

	obj.UI = container.NewWithoutLayout(data)
}

func (obj *Metadata) SetFormView() *widget.Form {

	items := []*widget.FormItem{
		0: {
			Text:   "AppID",
			Widget: obj.SetFormItem(obj.AppID),
		},
		1: {
			Text:   "Version",
			Widget: obj.SetFormItem(obj.Version),
		},
		2: {
			Text:   "Build",
			Widget: obj.SetFormItem(fmt.Sprintf("%v", obj.Build)),
		},
		3: {
			Text:   "UniqueID",
			Widget: obj.SetFormItem(obj.UniqueID),
		},
		4: {
			Text:   "CommitCode",
			Widget: obj.SetFormItem(obj.CommitCode),
		},
		5: {
			Text:   "Licence",
			Widget: obj.SetFormItem(obj.Licence),
		},
		6: {
			Text:   "StoragePath",
			Widget: obj.SetFormItem(obj.StoragePath),
		},
	}

	f := widget.NewForm(items...)

	return f
}

func (obj *Metadata) SetFormItem(text string) *widget.Entry {
	entry := widget.NewEntry()
	entry.Text = text
	entry.Disable()
	return entry
}
