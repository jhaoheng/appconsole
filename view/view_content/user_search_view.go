package view_content

import (
	"appconsole/module"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
	"github.com/sirupsen/logrus"
)

func SetUserSearchView(search_keys []string, callback func(search_result []module.User)) fyne.CanvasObject {

	var SelectedSearchKey string = ""

	options := []string{} // 初始可選的範圍
	entry := xwidget.NewCompletionEntry(options)
	entry.ActionItem = func() fyne.CanvasObject {
		btn := widget.NewButton("", func() {
			dialog.NewInformation("INFORMATION", "this is message", fyne.CurrentApp().Driver().AllWindows()[0]).Show()
		})
		btn.SetIcon(theme.InfoIcon())
		return btn
	}()

	// When the use typed text, complete the list.
	entry.OnChanged = func(s string) {
		// completion start for text length >= 2
		if len(s) < 2 {
			entry.HideCompletion()
			return
		}

		//
		keys, _, err := NewSearch().GetUserResult(SelectedSearchKey, s)
		if err != nil {
			logrus.Error(err)
			entry.HideCompletion()
			return
		}

		// no results
		if len(keys) == 0 {
			entry.HideCompletion()
			return
		}

		// then show them
		entry.SetOptions(keys) // 設定可選擇的內容
		entry.ShowCompletion()
	}

	//
	entry.OnCursorChanged = func() {
	}

	// 只有按下 button 時才有效, 不太友善
	entry.OnSubmitted = func(s string) {
		_, results, err := NewSearch().GetUserResult(SelectedSearchKey, s)
		if err != nil {
			logrus.Error(err)
			entry.HideCompletion()
			return
		}
		callback(results)
	}

	//
	selectkeyview := widget.NewSelect(search_keys, func(s string) {
		SelectedSearchKey = s
	})
	selectkeyview.SetSelectedIndex(0)

	return container.NewMax(container.NewBorder(
		nil,
		nil,
		selectkeyview,
		nil,
		entry,
	))
}

type Search struct {
}

func NewSearch() *Search {
	return &Search{}
}

func (s *Search) GetUserResult(search_key, search_value string) (keys []string, results []module.User, err error) {
	keys = []string{}
	if search_key == "name" {
		results, err = module.NewUser().SearchNameLike(search_value)
		for _, result := range results {
			keys = append(keys, result.Name)
		}
	} else if search_key == "member_id" {
		results, err = module.NewUser().SearchMemberIDLike(search_value)
		for _, result := range results {
			keys = append(keys, result.MemberID)
		}
	} else {
		err = fmt.Errorf("not set useful key")
		return
	}
	return
}
