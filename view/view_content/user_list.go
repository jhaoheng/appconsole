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

var UserTabledatas []binding.Struct
var page = 0
var numOfPage = 50

//
var myTable *widget.Table

//
func get_user_list() {
	UserTabledatas = []binding.Struct{}
	mydata := module.NewUser().List(numOfPage, page)
	for i := range mydata {
		UserTabledatas = append(UserTabledatas, binding.BindStruct(&mydata[i]))
	}
}

//
func UserListScreen(_ fyne.Window) fyne.CanvasObject {
	//
	get_user_list()
	//
	c := container.NewAdaptiveGrid(
		1,
		container.NewBorder(
			setTop(),
			nil,
			nil,
			nil,
			setTable(),
		),
	)
	myTable.Refresh()
	return c
}

func setTop() fyne.CanvasObject {
	// 可以透過 canvas 來製造畫面的 padding
	space := canvas.NewLine(color.Transparent)
	space.StrokeWidth = 5

	// count
	count := widget.NewLabel(fmt.Sprintf("all count : %v", len(UserTabledatas)))

	//
	top := container.NewVBox(
		container.NewHBox(
			space,
			setDelBtn(),
			layout.NewSpacer(),
			count,
		),
		space,
		widget.NewSeparator(), // 分隔的線段
		space,
	)
	return top
}

// data_id
type DelItem struct {
	CellID   widget.TableCellID
	DataID   int
	Checkbox *widget.Check
}

var selected_delete_items []DelItem

func setDelBtn() *widget.Button {
	delButton := widget.NewButton("delete", func() {
		for _, item := range selected_delete_items {
			if err := module.NewUser().Del(item.DataID); err != nil {
				logrus.Error(err)
			} else {
				item.Checkbox.Checked = false
				item.Checkbox.Refresh()
			}
		}
		get_user_list()
		myTable.Refresh()
		selected_delete_items = []DelItem{}
	})
	return delButton
}

//
func setTable() fyne.CanvasObject {
	myTable = widget.NewTable(
		tableSize,
		tableCreateCell,
		tableUpdateCell,
	)
	myTable.SetColumnWidth(0, 34)  //
	myTable.SetColumnWidth(1, 34)  //
	myTable.SetColumnWidth(2, 34)  //
	myTable.SetColumnWidth(3, 320) //
	myTable.SetColumnWidth(4, 100) //
	myTable.SetColumnWidth(5, 100) //
	myTable.SetColumnWidth(6, 100) //
	return myTable
}

/**/
func tableSize() (rows int, columns int) {
	return len(UserTabledatas), 7
}

/**/
func tableCreateCell() fyne.CanvasObject {
	c := container.NewMax(
		widget.NewCheck("", func(ok bool) {}),
		widget.NewLabel(""),
		container.NewMax(canvas.NewImageFromFile("./resources/img/default_user.jpg")),
	)
	return c
}

/**/
func tableUpdateCell(id widget.TableCellID, cell fyne.CanvasObject) {
	checkbox := cell.(*fyne.Container).Objects[0].(*widget.Check)
	label := cell.(*fyne.Container).Objects[1].(*widget.Label)
	img_container := cell.(*fyne.Container).Objects[2].(*fyne.Container)

	checkbox.Hide()
	label.Hide()
	img_container.Hide()
	//
	if id.Col == 0 {
		checkbox.Show()
	} else if id.Col == 2 {
		img_container.Show()
	} else {
		label.Show()
	}

	//
	switch id.Col {
	case 0:
		checkbox.OnChanged = func(ok bool) {
			data_id := tableCellGetValue(id.Row, "ID").(int)
			if ok {
				selected_delete_items = append(selected_delete_items, DelItem{
					DataID:   data_id,
					Checkbox: checkbox,
					CellID:   id,
				})
				sort.Slice(selected_delete_items, func(i int, j int) bool { return i < j })
			} else {
				for index, val := range selected_delete_items {
					if val.DataID == data_id {
						selected_delete_items = append(selected_delete_items[:index], selected_delete_items[index+1:]...)
					}
				}
			}
		}
	case 1:
		label.SetText(fmt.Sprintf("%d", tableCellGetValue(id.Row, "ID").(int)))
	case 2:
		resource := tableCellGetValue(id.Row, "Picture").([]byte)
		if resource != nil {
			img_container.Objects[0] = canvas.NewImageFromResource(fyne.NewStaticResource("img", resource))
			// resource := tableCellGetValue(id.Row, "PictureFilePath").(string)
			// img_container.Objects[0] = canvas.NewImageFromFile(resource)
			img_container.Refresh()
		}
	case 3:
		label.SetText(tableCellGetValue(id.Row, "MemberID").(string))
	case 4:
		label.SetText(tableCellGetValue(id.Row, "Name").(string))
	case 5:
		label.SetText(tableCellGetValue(id.Row, "Phone").(string))
	case 6:
		label.SetText(tableCellGetValue(id.Row, "Gender").(string))
	default:
		label.SetText("unknow")
	}
}

//
func tableCellGetValue(index int, key string) interface{} {
	val, err := UserTabledatas[index].GetValue(key)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return val
}
