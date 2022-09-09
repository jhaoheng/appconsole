package view_content

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NodataMaskScreen() *fyne.Container {
	background := canvas.NewRasterWithPixels(func(x int, y int, w int, h int) color.Color {
		return theme.BackgroundColor()
	})
	background.SetMinSize(fyne.NewSize(280, 280))
	//

	return container.NewMax(
		background,
		container.NewCenter(widget.NewLabel("nodata")),
	)
}
