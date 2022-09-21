package view_content

import (
	"appconsole/module"
	"fmt"
	"image/color"
	"sort"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

func UserListView(win fyne.Window) *fyne.Container {
	userList := NewUserList(win, 50, 0).RefreshTableDatas([]module.User{})
	userList.View = container.NewAdaptiveGrid(
		1,
		container.NewBorder(
			userList.SetTopView(), userList.SetAddButton(), nil, nil,
			userList.SetTableView(),
		),
	)

	//
	userList.MyTableViewContainer = userList.View.Objects[0].(*fyne.Container).Objects[0].(*fyne.Container)
	// userList.MyTableViewContainer.Hide()
	userList.TopViewContainer = userList.View.Objects[0].(*fyne.Container).Objects[1].(*fyne.Container)
	// userList.TopViewContainer.Hide()
	userList.MyTableView = userList.MyTableViewContainer.Objects[0].(*widget.Table)
	userList.NodataMaskContainer = userList.MyTableViewContainer.Objects[1].(*fyne.Container)

	//
	userList.RefreshTableDatas([]module.User{})
	return userList.View
}

// -
type UserList struct {
	Window fyne.Window
	View   *fyne.Container
	//
	Page         int
	NumOfPage    int
	AllItemCount binding.String
	Datas        []module.User
	Tabledatas   []binding.Struct
	//
	TopViewContainer     *fyne.Container
	MyTableViewContainer *fyne.Container
	NodataMaskContainer  *fyne.Container
	MyTableView          *widget.Table
	//
	MyTableDelItems []UserTableDelItem
}

type UserTableDelItem struct {
	CellID   widget.TableCellID
	DataID   int
	Checkbox *widget.Check
}

func NewUserList(win fyne.Window, numOfPage, page int) *UserList {
	userList := &UserList{
		Window:    win,
		Page:      page,
		NumOfPage: numOfPage,
	}
	return userList
}

func (view *UserList) SetTopView() *fyne.Container {
	//
	delButton := widget.NewButton("delete", func() {
		for _, item := range view.MyTableDelItems {
			if err := module.NewUser().Del(item.DataID); err != nil {
				logrus.Error(err)
			} else {
				item.Checkbox.Checked = false
				item.Checkbox.Refresh()
			}
		}
		view.RefreshTableDatas([]module.User{})
		view.MyTableDelItems = []UserTableDelItem{}
	})

	//
	topView := container.NewVBox(
		container.NewBorder(
			nil,
			nil,
			delButton,
			widget.NewLabelWithData(view.AllItemCount),
			view.SetUserSearchView([]string{"name", "member_id"}, func(result []module.User) {
				view.Datas = result
				view.RefreshTableDatas(result)
			}),
		),
		canvas.NewLine(color.Transparent),
		widget.NewSeparator(), // 分隔的線段
		canvas.NewLine(color.Transparent),
	)
	return topView
}

func (view *UserList) SetTableView() *fyne.Container {
	table := widget.NewTable(
		view.tableSize,
		view.tableCreateCell,
		view.tableUpdateCell,
	)
	table.SetColumnWidth(0, 34)  //
	table.SetColumnWidth(1, 34)  //
	table.SetColumnWidth(2, 320) //
	table.SetColumnWidth(3, 100) //
	table.SetColumnWidth(4, 110) //
	table.SetColumnWidth(5, 80)  //
	table.SetColumnWidth(6, 60)  //
	//
	myTableView := container.NewMax(
		table,
		NodataMaskView(),
	)
	return myTableView
}

func (view *UserList) RefreshTableDatas(user_datas []module.User) *UserList {
	//
	view.Tabledatas = []binding.Struct{}
	if len(user_datas) == 0 {
		view.Datas = module.NewUser().List(view.NumOfPage, view.Page)
	}
	for index := range view.Datas {
		view.Tabledatas = append(view.Tabledatas, binding.BindStruct(&view.Datas[index]))
	}

	//
	if view.AllItemCount == nil {
		view.AllItemCount = binding.NewString()
	}
	view.AllItemCount.Set(fmt.Sprintf("all count : %v", len(view.Tabledatas)))

	if view.NodataMaskContainer != nil {
		if len(view.Datas) == 0 {
			view.NodataMaskContainer.Show()
		} else {
			view.NodataMaskContainer.Hide()
		}
	}

	//
	if view.MyTableView != nil {
		view.MyTableView.Refresh()
	}
	return view
}

func (view *UserList) SetAddButton() *fyne.Container {
	addButton := widget.NewButton("", func() {
		userEdit := NewUserEdit(view.Window, module.User{
			MemberID: uuid.New().String(),
		}, func(new_user module.User) {
			new_user.ID = module.NewUser().Count() + 1
			module.NewUser().Create(&new_user)
			view.RefreshTableDatas([]module.User{})
			SendNotification("Add User", "Success!!")
		})
		userEdit.ShowModalView()
	})
	addButton.SetIcon(theme.ContentAddIcon())
	return container.NewMax(addButton)
}

/******************
****** Table ******
*******************/

/**/
func (view *UserList) tableSize() (rows int, columns int) {
	return len(view.Tabledatas), 7
}

/**/
func (view *UserList) tableCreateCell() fyne.CanvasObject {
	label := widget.NewLabel("")
	label.Wrapping = fyne.TextTruncate
	label.Alignment = fyne.TextAlignCenter
	//
	editBtn := widget.NewButton("edit", func() {})
	editBtn.Hide()
	//
	c := container.NewMax(
		widget.NewCheck("", func(ok bool) {}),
		label,
		editBtn,
	)
	return c
}

/**/
func (view *UserList) tableUpdateCell(id widget.TableCellID, cell fyne.CanvasObject) {
	checkbox := cell.(*fyne.Container).Objects[0].(*widget.Check)
	label := cell.(*fyne.Container).Objects[1].(*widget.Label)
	edit_btn := cell.(*fyne.Container).Objects[2].(*widget.Button)

	//
	checkbox.Hide()
	label.Hide()
	edit_btn.Hide()

	//
	switch id.Col {
	case 0:
		checkbox.Show()
		checkbox.OnChanged = func(ok bool) {
			data_id := view.tableCellGetValue(id.Row, "ID").(int)
			if ok {
				view.MyTableDelItems = append(view.MyTableDelItems, UserTableDelItem{
					DataID:   data_id,
					Checkbox: checkbox,
					CellID:   id,
				})
				sort.Slice(view.MyTableDelItems, func(i int, j int) bool { return i < j })
			} else {
				for index, val := range view.MyTableDelItems {
					if val.DataID == data_id {
						view.MyTableDelItems = append(view.MyTableDelItems[:index], view.MyTableDelItems[index+1:]...)
					}
				}
			}
		}
	case 1:
		label.Show()
		label.SetText(fmt.Sprintf("%d", view.tableCellGetValue(id.Row, "ID").(int)))
	case 2:
		label.Show()
		label.SetText(view.tableCellGetValue(id.Row, "MemberID").(string))
	case 3:
		label.Show()
		label.SetText(view.tableCellGetValue(id.Row, "Name").(string))
	case 4:
		label.Show()
		label.SetText(view.tableCellGetValue(id.Row, "Phone").(string))
	case 5:
		label.Show()
		label.SetText(view.tableCellGetValue(id.Row, "Gender").(string))
	case 6:
		edit_btn.Show()
		edit_btn.OnTapped = func() {
			userEdit := NewUserEdit(view.Window, view.Datas[id.Row], func(edited_user module.User) {
				view.Datas[id.Row] = edited_user
				view.RefreshTableDatas([]module.User{})
			})
			userEdit.ShowModalView()
		}
	default:
		label.SetText("undefined cell")
	}
}

func (view *UserList) tableCellGetValue(index int, key string) interface{} {
	val, err := view.Tabledatas[index].GetValue(key)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return val
}

/******************
****** Search ******
*******************/

func (view *UserList) SetUserSearchView(search_keys []string, callback func(search_result []module.User)) fyne.CanvasObject {

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
		keys, _, err := view.GetUserResult(SelectedSearchKey, s)
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
		_, results, err := view.GetUserResult(SelectedSearchKey, s)
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

func (view *UserList) GetUserResult(search_key, search_value string) (keys []string, results []module.User, err error) {
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
