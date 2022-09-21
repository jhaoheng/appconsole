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
			View:  viewcontent.WelcomeView,
		},
		"metadata": {
			Title: "Metadata",
			Intro: "The information of application.",
			View:  viewcontent.MetadataView,
		},
		"device": {
			Title: "Device",
			Intro: "",
			View:  viewcontent.DeviceListView,
		},
		"user": {
			Title: "User",
			Intro: "",
			View:  viewcontent.UserListView,
		},
		"log": {
			Title: "Access Log",
			Intro: "",
			View:  viewcontent.UserLogView,
		},
	}

	// ViewIndex  defines how our view should be laid out in the index tree
	ViewIndex = map[string][]string{
		"":        {"welcome", "device", "user", "log"},
		"welcome": {"metadata"},
	}
)
