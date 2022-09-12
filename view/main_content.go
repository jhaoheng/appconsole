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

var LoginSuccess = make(chan struct{})

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
		myWindow.Close()
	})
	info_text := fmt.Sprintf("%v - %v", a.Preferences().String("version"), a.Preferences().String("buildDate"))
	info := container.NewVBox(
		container.NewHBox(logout_btn),
		widget.NewLabel(info_text),
	)

	return container.NewBorder(nil, info, nil, nil, tree)
}
