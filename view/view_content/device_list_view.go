package view_content

import (
	"appconsole/module"
	"fmt"
	"image/color"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

// var deviceList *DeviceList

func DeviceListView(win fyne.Window) *fyne.Container {
	deviceList := NewDeviceList(win, 50, 0).RefreshTableDatas()
	deviceList.DeviceListView = container.NewAdaptiveGrid(
		1,
		container.NewBorder(
			deviceList.SetTopView(), deviceList.SetAddButton(), nil, nil,
			deviceList.SetTableView(),
		),
	)
	deviceList.MyTableViewContainer = deviceList.DeviceListView.Objects[0].(*fyne.Container).Objects[0].(*fyne.Container)
	// deviceList.MyTableViewContainer.Hide()
	deviceList.TopViewContainer = deviceList.DeviceListView.Objects[0].(*fyne.Container).Objects[1].(*fyne.Container)
	// deviceList.TopViewContainer.Hide()
	deviceList.MyTableView = deviceList.MyTableViewContainer.Objects[0].(*widget.Table)
	deviceList.NodataMaskContainer = deviceList.MyTableViewContainer.Objects[1].(*fyne.Container)
	// deviceList.NodataMaskContainer.Hide()

	//
	deviceList.RefreshTableDatas()

	return deviceList.DeviceListView
}

type DeviceList struct {
	Window         fyne.Window
	DeviceListView *fyne.Container
	//
	Page         int
	NumOfPage    int
	AllItemCount binding.String
	Datas        []module.Device
	Tabledatas   []binding.Struct
	//
	TopViewContainer     *fyne.Container
	MyTableViewContainer *fyne.Container
	NodataMaskContainer  *fyne.Container
	MyTableView          *widget.Table
	//
	MyTableDelItems []DeviceTableDelItem
}

type DeviceTableDelItem struct {
	CellID   widget.TableCellID
	DataID   int
	Checkbox *widget.Check
}

func NewDeviceList(win fyne.Window, numOfPage, page int) *DeviceList {
	deviceList := &DeviceList{
		Window:    win,
		Page:      page,
		NumOfPage: numOfPage,
	}
	return deviceList
}

func (view *DeviceList) SetTopView() *fyne.Container {
	// 可以透過 canvas 來製造畫面的 padding
	space := canvas.NewLine(color.Transparent)
	space.StrokeWidth = 5

	//
	delButton := widget.NewButton("delete", func() {
		for _, item := range view.MyTableDelItems {
			if err := module.NewDevice().Del(item.DataID); err != nil {
				logrus.Error(err)
			} else {
				item.Checkbox.Checked = false
				item.Checkbox.Refresh()
			}
		}
		view.RefreshTableDatas()
		view.MyTableDelItems = []DeviceTableDelItem{}
	})

	//
	topView := container.NewVBox(
		container.NewBorder(
			nil,
			nil,
			delButton,
			widget.NewLabelWithData(view.AllItemCount),
			view.SetPingView(),
		),
		space,
		widget.NewSeparator(), // 分隔的線段
		space,
	)
	return topView
}

func (view *DeviceList) SetPingView() fyne.CanvasObject {
	info := widget.NewButton("", func() {
		d := dialog.NewInformation("INFORMATION", "this is message", fyne.CurrentApp().Driver().AllWindows()[0])
		d.Show()
	})
	info.SetIcon(theme.InfoIcon())

	//
	entry := widget.NewEntry()
	entry.ActionItem = info
	entry.OnSubmitted = func(s string) {
		fmt.Println("驗證正確")
		alert := dialog.NewInformation("", "Do something after press btn", view.Window)
		alert.Show()
	}
	entry.PlaceHolder = "ping device, ex: 192.168.1.49"
	return container.NewMax(container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		entry,
	))
}

func (view *DeviceList) SetTableView() *fyne.Container {

	table := widget.NewTable(
		view.tableSize,
		view.tableCreateCell,
		view.tableUpdateCell,
	)

	table.SetColumnWidth(0, 34)  //
	table.SetColumnWidth(1, 34)  //
	table.SetColumnWidth(2, 100) //
	table.SetColumnWidth(3, 100) //
	table.SetColumnWidth(4, 150) //
	table.SetColumnWidth(5, 100) //
	table.SetColumnWidth(6, 100) //

	//
	myTableView := container.NewMax(
		table,
		NodataMaskView(),
	)

	return myTableView
}

func (view *DeviceList) RefreshTableDatas() *DeviceList {
	//
	view.Tabledatas = []binding.Struct{}
	view.Datas = module.NewDevice().List(view.NumOfPage, view.Page)
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

func (view *DeviceList) SetAddButton() *fyne.Container {
	addButton := widget.NewButton("", func() {
		module.FakeDevices = append(module.FakeDevices, module.Device{
			ID:           len(module.FakeDevices) + 1,
			Name:         "device_01",
			IP:           "192.168.0.1",
			MacAddress:   "xx:xx:xx:xx:xx:xx",
			DeviceSerial: "J91322386",
			Status:       true,
		})
		view.RefreshTableDatas()
		SendNotification("Add New Device", "Success")
	})
	addButton.SetIcon(theme.ContentAddIcon())
	return container.NewMax(addButton)
}

/******************
****** Table ******
*******************/

/**/
func (view *DeviceList) tableSize() (rows int, columns int) {
	return len(view.Tabledatas), 7
}

/**/
func (view *DeviceList) tableCreateCell() fyne.CanvasObject {
	c := container.NewMax(
		widget.NewCheck("", func(ok bool) {}),
		widget.NewLabel(""),
	)
	return c
}

/**/
func (view *DeviceList) tableUpdateCell(id widget.TableCellID, cell fyne.CanvasObject) {
	checkbox := cell.(*fyne.Container).Objects[0].(*widget.Check)
	label := cell.(*fyne.Container).Objects[1].(*widget.Label)

	//
	checkbox.Hide()
	label.Hide()

	//
	if id.Col == 0 {
		checkbox.Show()
	} else {
		label.Show()
	}

	//
	data_index := id.Row
	switch id.Col {
	case 0:
		checkbox.OnChanged = func(ok bool) {
			data_id := view.tableCellGetValue(data_index, "ID").(int)
			if ok {
				view.MyTableDelItems = append(view.MyTableDelItems, DeviceTableDelItem{
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
		label.SetText(fmt.Sprintf("%d", view.tableCellGetValue(data_index, "ID").(int)))
	case 2:
		label.SetText(view.tableCellGetValue(data_index, "Name").(string))
	case 3:
		label.SetText(view.tableCellGetValue(data_index, "IP").(string))
	case 4:
		label.SetText(view.tableCellGetValue(data_index, "MacAddress").(string))
	case 5:
		label.SetText(view.tableCellGetValue(data_index, "DeviceSerial").(string))
	case 6:
		label.SetText(fmt.Sprintf("%v", view.tableCellGetValue(data_index, "Status").(bool)))
	default:
		label.SetText("undefined cell")
	}
}

func (view *DeviceList) tableCellGetValue(index int, key string) interface{} {
	val, err := view.Tabledatas[index].GetValue(key)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return val
}

// func (view *DeviceList) tableSetHead(id widget.TableCellID, cell fyne.CanvasObject) {
// 	if id.Row != 0 {
// 		return
// 	}

// 	head_labels := []string{
// 		"",
// 		"ID",
// 		"Name",
// 		"IP",
// 		"MacAddress",
// 		"DeviceSerial",
// 		"Status",
// 	}

// 	label := cell.(*fyne.Container).Objects[1].(*widget.Label)
// 	label.Show()

// 	switch id.Col {
// 	case 0:
// 		label.SetText(head_labels[id.Col])
// 	case 1:
// 		label.SetText(head_labels[id.Col])
// 	case 2:
// 		label.SetText(head_labels[id.Col])
// 	case 3:
// 		label.SetText(head_labels[id.Col])
// 	case 4:
// 		label.SetText(head_labels[id.Col])
// 	case 5:
// 		label.SetText(head_labels[id.Col])
// 	case 6:
// 		label.SetText(head_labels[id.Col])
// 	default:
// 	}
// }
