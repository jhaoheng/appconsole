package view_content

import (
	"appconsole/module"
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
	"github.com/sirupsen/logrus"
)

func UserLogView(w fyne.Window) *fyne.Container {
	userLog := NewUserLog(w).RefreshTableDatas([]module.UserLog{})

	userLog.View = container.NewAdaptiveGrid(
		1,
		container.NewBorder(
			userLog.SetTopView(), nil, nil, nil,
			userLog.SetTableView(),
		),
	)
	//
	userLog.MyTableViewContainer = userLog.View.Objects[0].(*fyne.Container).Objects[0].(*fyne.Container)
	// userLog.MyTableViewContainer.Hide()
	userLog.TopViewContainer = userLog.View.Objects[0].(*fyne.Container).Objects[1].(*fyne.Container)
	// userLog.TopViewContainer.Hide()
	userLog.MyTableView = userLog.MyTableViewContainer.Objects[0].(*widget.Table)
	userLog.NodataMaskContainer = userLog.MyTableViewContainer.Objects[1].(*fyne.Container)

	//
	userLog.RefreshTableDatas([]module.UserLog{})
	return userLog.View
}

type UserLog struct {
	Window fyne.Window
	View   *fyne.Container
	//
	AllItemCount binding.String
	Datas        []module.UserLog
	Tabledatas   []binding.Struct
	//
	TopViewContainer     *fyne.Container
	MyTableViewContainer *fyne.Container
	NodataMaskContainer  *fyne.Container
	MyTableView          *widget.Table
}

func NewUserLog(win fyne.Window) *UserLog {
	return &UserLog{
		Window: win,
	}
}

func (view *UserLog) SetTopView() *fyne.Container {
	// 可以透過 canvas 來製造畫面的 padding
	space := canvas.NewLine(color.Transparent)
	space.StrokeWidth = 5

	//
	topView := container.NewVBox(
		container.NewBorder(
			nil, nil, nil, widget.NewLabelWithData(view.AllItemCount),
			view.SetUserLogSearchView([]string{"name"}, func(result []module.UserLog) {
				view.Datas = result
				view.RefreshTableDatas(result)
			}),
		),
		space,
		widget.NewSeparator(), // 分隔的線段
		space,
	)
	return topView
}

func (view *UserLog) SetTableView() *fyne.Container {
	table := widget.NewTable(
		view.tableSize,
		view.tableCreateCell,
		view.tableUpdateCell,
	)
	table.SetColumnWidth(0, 34)  //
	table.SetColumnWidth(1, 100) //
	table.SetColumnWidth(2, 170) //
	table.SetColumnWidth(3, 150) //
	table.SetColumnWidth(4, 170) //
	//
	myTableView := container.NewMax(
		table,
		NodataMaskView(),
	)
	return myTableView
}

func (view *UserLog) RefreshTableDatas(datas []module.UserLog) *UserLog {
	//
	view.Tabledatas = []binding.Struct{}
	if len(datas) == 0 {
		view.Datas = module.NewUserLog().GetAll()
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

/******************
****** Table ******
*******************/

/**/
func (view *UserLog) tableSize() (rows int, columns int) {
	return len(view.Tabledatas), 5
}

/**/
func (view *UserLog) tableCreateCell() fyne.CanvasObject {
	label := widget.NewLabel("")
	label.Wrapping = fyne.TextTruncate
	label.Alignment = fyne.TextAlignCenter
	c := container.NewMax(
		label,
	)
	return c
}

/**/
func (view *UserLog) tableUpdateCell(id widget.TableCellID, cell fyne.CanvasObject) {
	label := cell.(*fyne.Container).Objects[0].(*widget.Label)

	//
	switch id.Col {
	case 0:
		label.SetText(fmt.Sprintf("%d", view.tableCellGetValue(id.Row, "ID").(int)))
	case 1:
		label.SetText(view.tableCellGetValue(id.Row, "Name").(string))
	case 2:
		label.SetText(view.tableCellGetValue(id.Row, "RecordTime").(time.Time).Format("2006-01-02 15:04:05"))
	case 3:
		label.SetText(view.tableCellGetValue(id.Row, "Label").(string))
	case 4:
		label.SetText(view.tableCellGetValue(id.Row, "Created").(time.Time).Format("2006-01-02 15:04:05"))
	default:
		label.SetText("undefined cell")
	}
}

func (view *UserLog) tableCellGetValue(index int, key string) interface{} {
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

func (view *UserLog) SetUserLogSearchView(search_keys []string, callback func(search_result []module.UserLog)) fyne.CanvasObject {

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

func (view *UserLog) GetUserResult(search_key, search_value string) (keys []string, results []module.UserLog, err error) {
	keys = []string{}
	fmt.Println(search_key, search_value)
	if search_key == "name" {
		results, _ = module.NewUserLog().SearchNameLike(search_value)
		for _, result := range results {
			keys = append(keys, result.Name)
		}
	} else {
		err = fmt.Errorf("not set useful key")
		return
	}
	return
}
