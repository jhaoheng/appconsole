package view_content

import (
	"appconsole/module"
	"fmt"
	"image/color"
	"sort"

	"github.com/sirupsen/logrus"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var userList *UserList

func UserListScreen(win fyne.Window) fyne.CanvasObject {
	NewUserList(50, 0, win).RefreshTableDatas().SetTopView().SetTableView()
	//
	c := container.NewAdaptiveGrid(
		1,
		container.NewBorder(
			userList.TopView, nil, nil, nil,
			userList.MyTableView,
		),
	)
	return c
}

// -
type UserList struct {
	Window fyne.Window
	//
	Page           int
	NumOfPage      int
	AllItemCount   binding.String
	UserDatas      []module.User
	UserTabledatas []binding.Struct
	//
	TopView     fyne.CanvasObject
	MyTableView *widget.Table
	//
	MyTableDelItems []MyTableDelItem
}

// data_id
type MyTableDelItem struct {
	CellID   widget.TableCellID
	DataID   int
	Checkbox *widget.Check
}

func NewUserList(numOfPage, page int, win fyne.Window) *UserList {
	userList = &UserList{
		Window:    win,
		Page:      page,
		NumOfPage: numOfPage,
	}
	return userList
}

func (ul *UserList) SetTopView() *UserList {
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
		userList.RefreshTableDatas()
		ul.MyTableDelItems = []MyTableDelItem{}
	})

	//
	ul.TopView = container.NewVBox(
		container.NewHBox(
			space,
			delButton,
			layout.NewSpacer(),
			widget.NewLabelWithData(ul.AllItemCount),
		),
		space,
		widget.NewSeparator(), // 分隔的線段
		space,
	)
	return ul
}

func (ul *UserList) SetTableView() *UserList {
	ul.MyTableView = widget.NewTable(
		ul.tableSize,
		ul.tableCreateCell,
		ul.tableUpdateCell,
	)
	ul.MyTableView.SetColumnWidth(0, 34)  //
	ul.MyTableView.SetColumnWidth(1, 34)  //
	ul.MyTableView.SetColumnWidth(2, 320) //
	ul.MyTableView.SetColumnWidth(3, 100) //
	ul.MyTableView.SetColumnWidth(4, 110) //
	ul.MyTableView.SetColumnWidth(5, 80)  //
	ul.MyTableView.SetColumnWidth(6, 60)  //
	return ul
}

func (ul *UserList) RefreshTableDatas() *UserList {
	//
	ul.UserTabledatas = []binding.Struct{}
	ul.UserDatas = module.NewUser().List(ul.NumOfPage, ul.Page)
	for index := range ul.UserDatas {
		ul.UserTabledatas = append(ul.UserTabledatas, binding.BindStruct(&ul.UserDatas[index]))
	}

	//
	if ul.AllItemCount == nil {
		ul.AllItemCount = binding.NewString()
	}
	ul.AllItemCount.Set(fmt.Sprintf("all count : %v", len(ul.UserTabledatas)))

	//
	if userList.MyTableView != nil {
		userList.MyTableView.Refresh()
	}
	return ul
}

/******************
****** Table ******
*******************/

/**/
func (ul *UserList) tableSize() (rows int, columns int) {
	return len(userList.UserTabledatas), 7
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
				userList.MyTableDelItems = append(userList.MyTableDelItems, MyTableDelItem{
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
			userEdit := NewUserEdit(ul.Window, ul.UserDatas[id.Row], func(edited module.User) {
				// fmt.Printf("%+v\n", edited)
				ul.UserDatas[id.Row] = edited
				ul.RefreshTableDatas()
			})
			userEdit.ShowModalScreen()
		}
	default:
		label.SetText("undefined cell")
	}
}

func (ul *UserList) tableCellGetValue(index int, key string) interface{} {
	val, err := userList.UserTabledatas[index].GetValue(key)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return val
}
