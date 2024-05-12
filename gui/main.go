package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	application := app.New()
	window := application.NewWindow()
	window.Resize(fyne.NewSize(600, 500))
	window.ShowAndRun()
}
