package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainDialog struct {
	app    fyne.App
	window fyne.Window
}

func (dialog *MainDialog) Init() {
	helloMessage := widget.NewLabel("Hi! You were able to launch me")

	myButton := widget.NewButton("Press me", func() {
		helloMessage.SetText("Now you've pressed the button!")
	})

	dialog.window.SetContent(container.NewVBox(
		helloMessage,
		myButton,
	))
}

func (dialog *MainDialog) Show() {
	dialog.window.Show()
}
