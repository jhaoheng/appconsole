package mainmenu

import (
	"fyne.io/fyne/v2"
)

/*
- 主要 menu
*/
func MakeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	//
	file := build_file(a, w)

	// a quit item will be appended to our first (File) menu
	device := fyne.CurrentDevice()
	if !device.IsMobile() && !device.IsBrowser() {
		file.Items = append(
			file.Items,
			fyne.NewMenuItemSeparator(),
			build_settings_item(a, w),
			build_preferences_item(a, w),
		)
	}
	main := fyne.NewMainMenu(
		file,
		build_edit(w),
		build_help(a),
	)

	// 動態變更
	file_checked_item := file.Items[1]
	file_checked_item.Action = func() {
		file_checked_item.Checked = !file_checked_item.Checked
		main.Refresh()
	}
	return main
}
