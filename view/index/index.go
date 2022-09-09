package index

import (
	viewcontent "appconsole/view/view_content"

	"fyne.io/fyne/v2"
)

type ViewInfo struct {
	Title string
	Intro string
	View  func(w fyne.Window) *fyne.Container
}

var (
	// Defines the metadata for each view
	ViewMetadata = map[string]ViewInfo{
		"welcome": {
			Title: "Welcome",
			Intro: "",
			View:  viewcontent.WelcomeScreen,
		},
		"send_notification": {
			Title: "Send Notification",
			Intro: "Press btn to see the action.",
			View:  viewcontent.SendNotificationScreen,
		},
		"device": {
			Title: "Device",
			Intro: "",
			View:  viewcontent.DeviceListScreen,
		},
		"user": {
			Title: "User",
			Intro: "",
			View:  viewcontent.UserListScreen,
		},
		"log": {
			Title: "Log",
			Intro: "",
			View:  nil,
		},
		"member_log": {
			Title: "Member Access Log",
			Intro: "",
			View:  nil,
		},
		"admin_log": {
			Title: "Admin Access Log",
			Intro: "",
			View:  nil,
		},
		"advanced": {
			Title: "Advanced",
			Intro: "",
			View:  nil,
		},
	}

	// ViewIndex  defines how our view should be laid out in the index tree
	ViewIndex = map[string][]string{
		"":         {"welcome", "send_notification", "device", "user", "log", "advanced"},
		"log":      {"member_log", "admin_log"},
		"advanced": {},
	}
)
