package view_content

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SendNotificationScreen(_ fyne.Window) *fyne.Container {
	return container.NewMax(
		container.NewCenter(
			widget.NewButton("send", func() {
				SendNotification("title", "content")
			}),
		),
	)
}

func SendNotification(title, content string) {
	myApp := fyne.CurrentApp()
	myApp.SendNotification(fyne.NewNotification(title, content))
}
