package mainmenu

import (
	"fyne.io/fyne/v2"
)

func build_preferences_item(a fyne.App, w fyne.Window) *fyne.MenuItem {
	item := fyne.NewMenuItem("Preferences", nil)
	return item
}
