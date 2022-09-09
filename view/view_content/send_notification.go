package view_content

import (
	"fyne.io/fyne/v2"
)

func SendNotification(title, content string) {
	myApp := fyne.CurrentApp()
	myApp.SendNotification(fyne.NewNotification(title, content))
}
