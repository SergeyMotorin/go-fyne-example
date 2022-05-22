package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// On Ubuntu you need install some packages:
// apt-get install libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("My Window")

	helloMessage := widget.NewLabel("Hi! You were able to launch me")
	myButton := widget.NewButton("Press me", func() {
		helloMessage.SetText("Now you've pressed the button!")
	})

	myWindow.SetContent(container.NewVBox(
		helloMessage,
		myButton,
	))

	myWindow.ShowAndRun()
}