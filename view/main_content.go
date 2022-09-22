package view

import (
	"appconsole/view/index"
	"appconsole/view/view_content"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var loadPreviousKey = "loadPreviousKey"

var ShowLoginView = false

// 主畫面
func MainContainer(myWindow fyne.Window) fyne.CanvasObject {
	// 右邊框架
	presentContent := view_content.NewContent()

	// 左邊框架
	myLeftCanvas := buildLeftCanvas(
		myWindow,
		func(v index.ViewInfo) {
			// 更新 content
			presentContent.ContentTitle.SetText(v.Title)
			presentContent.ContentIntro.SetText(v.Intro)
			if len(v.Intro) == 0 {
				presentContent.ContentIntro.Hide()
			} else {
				presentContent.ContentIntro.Show()
			}
			if v.View != nil {
				presentContent.Content.Objects = []fyne.CanvasObject{v.View(myWindow)}
			} else {
				presentContent.Content.Objects = []fyne.CanvasObject{}
			}
			presentContent.Content.Refresh()
		},
		true,
	)
	//
	split := container.NewHSplit(myLeftCanvas, presentContent.GetCanvas())
	split.Offset = 0.2

	return split
}

// 側邊欄
func buildLeftCanvas(myWindow fyne.Window, setContent func(v index.ViewInfo), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()
	//
	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return index.ViewIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := index.ViewIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			viewMetadata, ok := index.ViewMetadata[uid]
			if !ok {
				fyne.LogError("Missing UI: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(viewMetadata.Title)
			obj.(*widget.Label).TextStyle = fyne.TextStyle{}
		},
		OnSelected: func(uid string) {
			if viewMetadata, ok := index.ViewMetadata[uid]; ok {
				a.Preferences().SetString(loadPreviousKey, uid)
				setContent(viewMetadata)
			}
		},
	}

	//
	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(loadPreviousKey, "welcome")
		tree.Select(currentPref)
	}

	//
	logout_btn := widget.NewButton("Logout", func() {
		myApp := fyne.CurrentApp()
		myApp.Preferences().SetBool("remember_me", false)
		SwitchLoginView(false, true)
	})
	info_text := fmt.Sprintf("%v", a.Preferences().String("version"))
	info := container.NewVBox(
		container.NewHBox(logout_btn),
		widget.NewLabel(info_text),
	)

	return container.NewBorder(nil, info, nil, nil, tree)
}

/**/
func SwitchLoginView(openMainWin bool, openLoginWin bool) {
	myApp := fyne.CurrentApp()
	main_window := myApp.Driver().AllWindows()[0]
	login_window := myApp.Driver().AllWindows()[1]

	if openMainWin {
		main_window.Show()
	} else {
		main_window.Hide()
	}

	if openLoginWin {
		login_window.Show()
	} else {
		login_window.Hide()
	}
}
