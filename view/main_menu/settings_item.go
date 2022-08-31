package mainmenu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/driver/desktop"
)

func build_settings_item(a fyne.App, w fyne.Window) *fyne.MenuItem {
	openSettings := func() {
		w := a.NewWindow("Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	}
	// 預設設定 Settings, 則會自動歸類在主欄位下
	settingsItem := fyne.NewMenuItem("Settings", openSettings)

	// 加入快捷鍵
	settingsShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	settingsItem.Shortcut = settingsShortcut
	w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings()
	})
	return settingsItem
}
