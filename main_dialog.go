package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
)

type MainDialog struct {
	app    fyne.App
	window fyne.Window
}

func (dialog *MainDialog) Init() {
	infoText := widget.NewLabel("What is two plus two?")

	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), infoText, layout.NewSpacer())

	button1 := widget.NewButton("It's four", func() {
		infoText.SetText("Excelent!")
	})

	button2 := widget.NewButton("It's five", func() {
		infoText.SetText("You're wrong!")
	})

	dialog.window.SetContent(container.NewVBox(centered, button1, button2))
}

func (dialog *MainDialog) Show() {
	dialog.window.Show()
}
