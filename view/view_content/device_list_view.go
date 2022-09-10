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
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

var deviceList *DeviceList

func DeviceListView(win fyne.Window) *fyne.Container {
	deviceList = NewDeviceList(win, 50, 0).RefreshTableDatas()
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
	deviceList = &DeviceList{
		Window:    win,
		Page:      page,
		NumOfPage: numOfPage,
	}
	return deviceList
}

func (dl *DeviceList) SetTopView() *fyne.Container {
	// 可以透過 canvas 來製造畫面的 padding
	space := canvas.NewLine(color.Transparent)
	space.StrokeWidth = 5
	//
	delButton := widget.NewButton("delete", func() {
		for _, item := range dl.MyTableDelItems {
			if err := module.NewDevice().Del(item.DataID); err != nil {
				logrus.Error(err)
			} else {
				item.Checkbox.Checked = false
				item.Checkbox.Refresh()
			}
		}
		deviceList.RefreshTableDatas()
		dl.MyTableDelItems = []DeviceTableDelItem{}
	})

	//
	topView := container.NewVBox(
		container.NewHBox(
			space,
			delButton,
			layout.NewSpacer(),
			widget.NewLabelWithData(dl.AllItemCount),
		),
		space,
		widget.NewSeparator(), // 分隔的線段
		space,
	)
	return topView
}

func (dl *DeviceList) SetTableView() *fyne.Container {

	table := widget.NewTable(
		dl.tableSize,
		dl.tableCreateCell,
		dl.tableUpdateCell,
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

func (dl *DeviceList) RefreshTableDatas() *DeviceList {
	//
	dl.Tabledatas = []binding.Struct{}
	dl.Datas = module.NewDevice().List(dl.NumOfPage, dl.Page)
	for index := range dl.Datas {
		dl.Tabledatas = append(dl.Tabledatas, binding.BindStruct(&dl.Datas[index]))
	}

	//
	if dl.AllItemCount == nil {
		dl.AllItemCount = binding.NewString()
	}
	dl.AllItemCount.Set(fmt.Sprintf("all count : %v", len(dl.Tabledatas)))

	if deviceList.NodataMaskContainer != nil {
		if len(deviceList.Datas) == 0 {
			deviceList.NodataMaskContainer.Show()
		} else {
			deviceList.NodataMaskContainer.Hide()
		}
	}

	//
	if deviceList.MyTableView != nil {
		deviceList.MyTableView.Refresh()
	}
	return dl
}

func (dl *DeviceList) SetAddButton() *fyne.Container {
	addButton := widget.NewButton("", func() {
		module.FakeDataDevices = append(module.FakeDataDevices, module.Device{
			ID:           len(module.FakeDataDevices) + 1,
			Name:         "device_01",
			IP:           "192.168.0.1",
			MacAddress:   "xx:xx:xx:xx:xx:xx",
			DeviceSerial: "J91322386",
			Status:       true,
		})
		dl.RefreshTableDatas()
		SendNotification("Add New Device", "Success")
	})
	addButton.SetIcon(theme.ContentAddIcon())
	return container.NewMax(addButton)
}

/******************
****** Table ******
*******************/

/**/
func (dl *DeviceList) tableSize() (rows int, columns int) {
	return len(deviceList.Tabledatas), 7
}

/**/
func (dl *DeviceList) tableCreateCell() fyne.CanvasObject {
	c := container.NewMax(
		widget.NewCheck("", func(ok bool) {}),
		widget.NewLabel(""),
	)
	return c
}

/**/
func (dl *DeviceList) tableUpdateCell(id widget.TableCellID, cell fyne.CanvasObject) {
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
	switch id.Col {
	case 0:
		checkbox.OnChanged = func(ok bool) {
			data_id := dl.tableCellGetValue(id.Row, "ID").(int)
			if ok {
				deviceList.MyTableDelItems = append(deviceList.MyTableDelItems, DeviceTableDelItem{
					DataID:   data_id,
					Checkbox: checkbox,
					CellID:   id,
				})
				sort.Slice(deviceList.MyTableDelItems, func(i int, j int) bool { return i < j })
			} else {
				for index, val := range deviceList.MyTableDelItems {
					if val.DataID == data_id {
						deviceList.MyTableDelItems = append(deviceList.MyTableDelItems[:index], deviceList.MyTableDelItems[index+1:]...)
					}
				}
			}
		}
	case 1:
		label.SetText(fmt.Sprintf("%d", dl.tableCellGetValue(id.Row, "ID").(int)))
	case 2:
		label.SetText(dl.tableCellGetValue(id.Row, "Name").(string))
	case 3:
		label.SetText(dl.tableCellGetValue(id.Row, "IP").(string))
	case 4:
		label.SetText(dl.tableCellGetValue(id.Row, "MacAddress").(string))
	case 5:
		label.SetText(dl.tableCellGetValue(id.Row, "DeviceSerial").(string))
	case 6:
		label.SetText(fmt.Sprintf("%v", dl.tableCellGetValue(id.Row, "Status").(bool)))
	default:
		label.SetText("undefined cell")
	}
}

func (dl *DeviceList) tableCellGetValue(index int, key string) interface{} {
	val, err := deviceList.Tabledatas[index].GetValue(key)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return val
}
