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
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var userList *UserList

func UserListView(win fyne.Window) *fyne.Container {
	userList = NewUserList(win, 50, 0).RefreshTableDatas([]module.User{})
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
	userList = &UserList{
		Window:    win,
		Page:      page,
		NumOfPage: numOfPage,
	}
	return userList
}

func (ul *UserList) SetTopView() *fyne.Container {
	// 可以透過 canvas 來製造畫面的 padding
	space := canvas.NewLine(color.Transparent)
	space.StrokeWidth = 5
	//
	delButton := widget.NewButton("delete", func() {
		for _, item := range ul.MyTableDelItems {
			if err := module.NewUser().Del(item.DataID); err != nil {
				logrus.Error(err)
			} else {
				item.Checkbox.Checked = false
				item.Checkbox.Refresh()
			}
		}
		userList.RefreshTableDatas([]module.User{})
		ul.MyTableDelItems = []UserTableDelItem{}
	})

	//
	topView := container.NewVBox(
		container.NewBorder(
			nil,
			nil,
			delButton,
			widget.NewLabelWithData(ul.AllItemCount),
			SetUserSearchView([]string{"name", "member_id"}, func(result []module.User) {
				ul.Datas = result
				ul.RefreshTableDatas(result)
			}),
		),
		space,
		widget.NewSeparator(), // 分隔的線段
		space,
	)
	return topView
}

func (ul *UserList) SetTableView() *fyne.Container {
	table := widget.NewTable(
		ul.tableSize,
		ul.tableCreateCell,
		ul.tableUpdateCell,
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

func (ul *UserList) RefreshTableDatas(user_datas []module.User) *UserList {
	//
	ul.Tabledatas = []binding.Struct{}
	if len(user_datas) == 0 {
		ul.Datas = module.NewUser().List(ul.NumOfPage, ul.Page)
	}
	for index := range ul.Datas {
		ul.Tabledatas = append(ul.Tabledatas, binding.BindStruct(&ul.Datas[index]))
	}

	//
	if ul.AllItemCount == nil {
		ul.AllItemCount = binding.NewString()
	}
	ul.AllItemCount.Set(fmt.Sprintf("all count : %v", len(ul.Tabledatas)))

	if userList.NodataMaskContainer != nil {
		if len(userList.Datas) == 0 {
			userList.NodataMaskContainer.Show()
		} else {
			userList.NodataMaskContainer.Hide()
		}
	}

	//
	if userList.MyTableView != nil {
		userList.MyTableView.Refresh()
	}
	return ul
}

func (ul *UserList) SetAddButton() *fyne.Container {
	addButton := widget.NewButton("", func() {
		userEdit := NewUserEdit(ul.Window, module.User{
			MemberID: uuid.New().String(),
		}, func(new_user module.User) {
			new_user.ID = module.NewUser().Count() + 1
			module.NewUser().Create(&new_user)
			ul.RefreshTableDatas([]module.User{})
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
func (ul *UserList) tableSize() (rows int, columns int) {
	return len(userList.Tabledatas), 7
}

/**/
func (ul *UserList) tableCreateCell() fyne.CanvasObject {
	c := container.NewMax(
		widget.NewCheck("", func(ok bool) {}),
		widget.NewLabel(""),
		container.NewCenter(widget.NewButton("edit", func() {})),
	)
	return c
}

/**/
func (ul *UserList) tableUpdateCell(id widget.TableCellID, cell fyne.CanvasObject) {
	checkbox := cell.(*fyne.Container).Objects[0].(*widget.Check)
	label := cell.(*fyne.Container).Objects[1].(*widget.Label)
	edit_btn := cell.(*fyne.Container).Objects[2].(*fyne.Container).Objects[0].(*widget.Button)

	//
	checkbox.Hide()
	label.Hide()
	edit_btn.Hide()
	//
	if id.Col == 0 {
		checkbox.Show()
	} else if id.Col == 6 {
		edit_btn.Show()
	} else {
		label.Show()
	}

	//
	switch id.Col {
	case 0:
		checkbox.OnChanged = func(ok bool) {
			data_id := ul.tableCellGetValue(id.Row, "ID").(int)
			if ok {
				userList.MyTableDelItems = append(userList.MyTableDelItems, UserTableDelItem{
					DataID:   data_id,
					Checkbox: checkbox,
					CellID:   id,
				})
				sort.Slice(userList.MyTableDelItems, func(i int, j int) bool { return i < j })
			} else {
				for index, val := range userList.MyTableDelItems {
					if val.DataID == data_id {
						userList.MyTableDelItems = append(userList.MyTableDelItems[:index], userList.MyTableDelItems[index+1:]...)
					}
				}
			}
		}
	case 1:
		label.SetText(fmt.Sprintf("%d", ul.tableCellGetValue(id.Row, "ID").(int)))
	case 2:
		label.SetText(ul.tableCellGetValue(id.Row, "MemberID").(string))
	case 3:
		label.SetText(ul.tableCellGetValue(id.Row, "Name").(string))
	case 4:
		label.SetText(ul.tableCellGetValue(id.Row, "Phone").(string))
	case 5:
		label.SetText(ul.tableCellGetValue(id.Row, "Gender").(string))
	case 6:
		edit_btn.OnTapped = func() {
			userEdit := NewUserEdit(ul.Window, ul.Datas[id.Row], func(edited_user module.User) {
				ul.Datas[id.Row] = edited_user
				ul.RefreshTableDatas([]module.User{})
			})
			userEdit.ShowModalView()
		}
	default:
		label.SetText("undefined cell")
	}
}

func (ul *UserList) tableCellGetValue(index int, key string) interface{} {
	val, err := userList.Tabledatas[index].GetValue(key)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return val
}
