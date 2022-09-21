package view

import (
	"appconsole/module"
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

func LoginContent(myWindow fyne.Window) fyne.CanvasObject {
	myApp := fyne.CurrentApp()

	//
	name := widget.NewEntry()
	name.SetPlaceHolder("John Smith")
	name.Validator = validation.NewAllStrings(func(s string) error {
		if len(s) == 0 {
			return errors.New("not allow empty")
		}
		return nil
	})
	if len(myApp.Preferences().String("login_name")) != 0 {
		name.Text = myApp.Preferences().String("login_name")
	}

	//
	email := widget.NewEntry()
	email.SetPlaceHolder("test@example.com")
	email.Validator = validation.NewRegexp(`\w{1,}@\w{1,}\.\w{1,4}`, "not a valid email")
	if len(myApp.Preferences().String("login_email")) != 0 {
		email.Text = myApp.Preferences().String("login_email")
	}

	//
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	//
	remember_me_box := widget.NewCheck("", func(is_checked bool) {
		fmt.Println("is remember me:", is_checked)
		myApp := fyne.CurrentApp()
		if is_checked {
			myApp.Preferences().SetBool("remember_me", true)
		} else {
			myApp.Preferences().SetBool("remember_me", false)
		}
	})

	//
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: name, HintText: "Your full name"},
			{Text: "Email", Widget: email, HintText: "A valid email address"},
		},
		OnCancel: func() {
			fmt.Println("Cancelled")
			myWindow.Close()
		},
		OnSubmit: func() {
			fmt.Println("Form submitted")
			if module.NewLogin(name.Text, password.Text).Check() {
				myApp.Preferences().SetString("login_name", name.Text)
				myApp.Preferences().SetString("login_email", email.Text)
				//
				SwitchLoginView()
			}
		},
		SubmitText: "Login",
	}
	form.Append("Password", password)
	form.Append("Remember", remember_me_box)
	return form
}
