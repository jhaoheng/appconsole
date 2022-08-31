package view_content

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SendNotificationScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewMax(
		container.NewCenter(
			widget.NewButton("send", sendNotification),
		),
	)
}

func sendNotification() {
	myApp := fyne.CurrentApp()
	myApp.SendNotification(fyne.NewNotification("title", "content"))
}
