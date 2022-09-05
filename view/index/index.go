package index

import (
	viewcontent "appconsole/view/view_content"

	"fyne.io/fyne/v2"
)

type ViewInfo struct {
	Title string
	Intro string
	View  func(w fyne.Window) fyne.CanvasObject
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
			View:  nil,
		},
		"add_device": {
			Title: "Add New Device",
			Intro: "",
			View:  nil,
		},
		"list_device": {
			Title: "Device List",
			Intro: "",
			View:  nil,
		},
		"user": {
			Title: "User",
			Intro: "",
			View:  nil,
		},
		"add_user": {
			Title: "Add New User",
			Intro: "",
			View:  nil,
		},
		"list_user": {
			Title: "User List",
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
		"device":   {"add_device", "list_device"},
		"user":     {"add_user", "list_user"},
		"log":      {"member_log", "admin_log"},
		"advanced": {},
	}
)
