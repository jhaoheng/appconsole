package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	extwidget "fyne.io/x/fyne/widget"
)

/*
- 當需要時, 將 loading 加入到內容最上層
*/
func LoadingScreen(size fyne.Size) *fyne.Container {
	gif, err := extwidget.NewAnimatedGif(storage.NewFileURI("../resources/gif/loading.gif"))
	if err != nil {
		panic(err)
	}
	gif.Start()
	//
	container := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// bg color
		canvas.NewRectangle(
			color.NRGBA{R: 255, G: 255, B: 255, A: 255},
		),
		// the content
		gif,
	)

	// 一定要設定畫布大小
	container.Resize(size)
	return container
}
