package view

import (
	"appconsole/module"
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

//
func LoginContent(myWindow fyne.Window) fyne.CanvasObject {
	//
	name := widget.NewEntry()
	name.SetPlaceHolder("John Smith")
	name.Validator = validation.NewAllStrings(func(s string) error {
		if len(s) == 0 {
			return errors.New("not allow empty")
		}
		return nil
	})

	// email := widget.NewEntry()
	// email.SetPlaceHolder("test@example.com")
	// email.Validator = validation.NewRegexp(`\w{1,}@\w{1,}\.\w{1,4}`, "not a valid email")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	password.Validator = validation.NewAllStrings(func(s string) error {
		if len(s) == 0 {
			return errors.New("not allow empty")
		}
		return nil
	})

	// disabled := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(string) {})
	// disabled.Horizontal = true
	// disabled.Disable()
	// largeText := widget.NewMultiLineEntry()
	check_box := widget.NewCheck("", func(is_checked bool) {
		fmt.Println("is checked:", is_checked)
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: name, HintText: "Your full name"},
			// {Text: "Email", Widget: email, HintText: "A valid email address"},
		},
		// OnCancel: func() {
		// 	fmt.Println("Cancelled")
		// 	myWindow.Close()
		// },
		OnSubmit: func() {
			fmt.Println("Form submitted")
			if module.NewLogin(name.Text, password.Text).Check() {
				LoginSuccess <- struct{}{}
			}
		},
		SubmitText: "OK",
	}
	form.Append("Password", password)
	form.Append("Remember", check_box)
	return form
}
