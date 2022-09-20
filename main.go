package main

import (
	"embed"
	"log"
	"os"
	"time"

	"appconsole/module"
	"appconsole/view"
	mainmenu "appconsole/view/main_menu"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

//go:embed resources
var resource embed.FS

var (
	Version    string = "develop ver"
	CommitCode string = ""
	BuildDate  string = time.Now().Format("2006-01-02")
	Title             = "app console"
)

var myApp fyne.App

func init() {
	module.LoadFont()
}

func main() {
	myApp = app.New()
	myApp = app.NewWithID("app.console.demo")
	//
	makeTray(myApp)
	logLifecycle(myApp)
	//
	module.Resource = &resource
	//
	myApp.Preferences().SetString("version", Version)
	myApp.Preferences().SetString("buildDate", BuildDate)

	//
	myWindow := myApp.NewWindow(Title)
	myWindow.SetMainMenu(mainmenu.MakeMenu(myApp, myWindow))
	myWindow.Resize(fyne.NewSize(1200, 750))
	myWindow.CenterOnScreen()
	myWindow.SetContent(view.MainContainer(myWindow))
	myWindow.SetMaster()

	//
	if myApp.Preferences().Bool("remember_me") {
		myWindow.Show()
	} else {
		myWindow.Hide()
		//
		loginWindow := myApp.NewWindow("login")
		loginWindow.Resize(fyne.NewSize(400, 150))
		loginWindow.CenterOnScreen()
		loginWindow.SetContent(view.LoginContent(loginWindow))
		loginWindow.SetFixedSize(true)
		loginWindow.SetMaster()
		loginWindow.Show()
		//
		go func() {
			<-view.LoginSuccess
			loginWindow.Hide()
			myWindow.Show()
		}()
	}

	//
	myApp.Run()
	//
	os.Unsetenv("FYNE_FONT")
}

func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

/*
- osx 右上角的小圖示
- window 右下角的小圖示
- 點選後, 會有選單
*/
func makeTray(a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		h := fyne.NewMenuItem("Hello", func() {})
		menu := fyne.NewMenu("Hello World", h)
		h.Action = func() {
			log.Println("System tray menu tapped")
			h.Label = "Welcome"
			myApp.SendNotification(fyne.NewNotification("notification", "hello world"))
			menu.Refresh()
		}
		desk.SetSystemTrayMenu(menu)
	}
}
