package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// On Ubuntu you need install some packages:
// apt-get install libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev

func initMainDialog(app fyne.App) *MainDialog {
	newDialog := &MainDialog{app: app, window: app.NewWindow("Main Window")}

	newDialog.Init()

	return newDialog
}

func main() {
	myApp := app.New()

	mainDialog := initMainDialog(myApp)

	mainDialog.Show()

	myApp.Run()

	displayExited()
}

func displayExited() {
	fmt.Println("Exited")
}
