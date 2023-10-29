package main

import (
	//"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	
	//"fyne.io/fyne/v2/widget"

	"./graphics"
)


const (
	height = 886
	width  = 1120
)



func main() {
	a := app.New()
	w := a.NewWindow("Train")

	// Рабочая зона
	cnv := graphics.MakeImageCanvas(width, height)
	rast := canvas.NewRasterFromImage(cnv.Image())
	img := container.New(layout.NewGridWrapLayout(fyne.NewSize(width, height)), rast)


	// боковое меню
	menu := graphics.CreateMenu(a) 

	menuColumn := container.New(layout.NewGridWrapLayout(fyne.NewSize(300, height)), menu)
	form := container.New(layout.NewFormLayout(), menuColumn, img)
	w.SetContent(form)


	w.Resize(fyne.NewSize(width, height))
	w.ShowAndRun()
}
