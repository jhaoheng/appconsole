package view_content

import (
	"appconsole/module"
	"io"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var userEdit *UserEdit

type UserEdit struct {
	Window              fyne.Window
	UserEditScreen      *fyne.Container
	UserEditScreenModal *widget.PopUp
	//
	UserBinding  binding.Struct
	ImageBinding binding.Bytes
	UserImg      *canvas.Image
	//
	TopView     fyne.CanvasObject
	ContentView fyne.CanvasObject
	//
	edited_callback func(edited module.User)
}

func NewUserEdit(window fyne.Window, user module.User, edited_callback func(edited module.User)) *UserEdit {
	//
	userEdit = &UserEdit{
		Window:          window,
		UserBinding:     binding.BindStruct(&user),
		edited_callback: edited_callback,
	}
	img := binding.NewBytes()
	img.Set(user.Picture)
	userEdit.ImageBinding = img
	//
	userEdit.UserEditScreen = container.NewAdaptiveGrid(
		1,
		container.NewBorder(
			userEdit.SetTop(), nil, nil, nil,
			userEdit.SetContent(),
		),
	)
	userEdit.UserEditScreenModal = widget.NewModalPopUp(userEdit.UserEditScreen, window.Canvas())
	return userEdit
}

func (ue *UserEdit) GetView() fyne.CanvasObject {
	return ue.UserEditScreen
}

func (ue *UserEdit) ShowModalScreen() {
	ue.UserEditScreenModal.Show()
}

func (ue *UserEdit) HideModalScreen() {
	ue.UserEditScreenModal.Hide()
}

func (ue *UserEdit) GetUserData(key string) (value interface{}) {
	v, err := ue.UserBinding.GetValue(key)
	if err != nil {
		panic(err)
	}
	return v
}

/*
- 存放 image
*/
func (ue *UserEdit) SetTop() fyne.CanvasObject {
	img, err := ue.ImageBinding.Get()
	if err != nil {
		panic(err)
	}
	ue.UserImg = canvas.NewImageFromResource(fyne.NewStaticResource("", img))
	ue.UserImg.SetMinSize(fyne.NewSize(300, 300))
	ue.UserImg.Resize(fyne.NewSize(300, 300))

	//
	imageContainer := container.NewHBox(
		layout.NewSpacer(),
		container.NewMax(ue.UserImg),
		layout.NewSpacer(),
	)

	//
	top := container.NewVBox(
		imageContainer,
		ue.OpenImageBtn(fyne.CurrentApp().Driver().AllWindows()[0]),
	)
	return top
}

/*
- 存放其他欄位
*/
func (ue *UserEdit) SetContent() fyne.CanvasObject {
	NameEntry := func() *widget.Entry {
		tmp := widget.NewEntry()
		tmp.SetText(ue.GetUserData("Name").(string))
		return tmp
	}()
	PhoneEntry := func() *widget.Entry {
		tmp := widget.NewEntry()
		tmp.SetText(ue.GetUserData("Phone").(string))
		return tmp
	}()
	GenderEntry := func() *widget.Entry {
		tmp := widget.NewEntry()
		tmp.SetText(ue.GetUserData("Gender").(string))
		return tmp
	}()
	TextArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "MemberID", Widget: widget.NewLabel(ue.GetUserData("MemberID").(string))},
			{Text: "Name", Widget: NameEntry},
			{Text: "Phone", Widget: PhoneEntry},
			{Text: "Gender", Widget: GenderEntry},
		},
		OnSubmit: func() {
			user := &module.User{
				ID:       5,
				MemberID: ue.GetUserData("MemberID").(string),
				Name:     NameEntry.Text,
				Phone:    PhoneEntry.Text,
				Gender:   GenderEntry.Text,
				Picture: func() []byte {
					tmp, _ := ue.ImageBinding.Get()
					return tmp
				}(),
			}
			ue.edited_callback(*user)
			userEdit.HideModalScreen()
		},
		OnCancel: func() {
			userEdit.HideModalScreen()
		},
	}
	// we can also append items
	form.Append("Text", TextArea)

	return form
}

func (ue *UserEdit) OpenImageBtn(parentWindow fyne.Window) *fyne.Container {
	openFile := widget.NewButton("File Open With Filter (.jpg or .png)", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, parentWindow)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}

			ue.imageOpened(reader)
		}, parentWindow)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
		fd.Show()
	})
	return container.NewVBox(openFile)
}

func (ue *UserEdit) imageOpened(f fyne.URIReadCloser) {
	if f == nil {
		log.Println("Cancelled")
		return
	}
	defer f.Close()

	ue.showImage(f)
}

func (ue *UserEdit) showImage(f fyne.URIReadCloser) {
	data, err := io.ReadAll(f)
	if err != nil {
		fyne.LogError("Failed to load image data", err)
		panic(err)
	}
	//
	ue.ImageBinding.Set(data)
	//
	imgContainer := ue.UserEditScreen.Objects[0].(*fyne.Container).Objects[1].(*fyne.Container).Objects[0].(*fyne.Container).Objects[1].(*fyne.Container)
	imgContainer.Add(canvas.NewImageFromResource(fyne.NewStaticResource("", data)))
}
