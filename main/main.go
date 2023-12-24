package main

import (
	//"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"./camera"
	"./graphics"
	"./mathfunc"
	"./menu"
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
	engine := graphics.NewMyGraphicsEngine(cnv)
	engine.ProjMatrix = mathfunc.MakeFovProjectionM(90.0, float64(engine.Cnv.Height())/float64(engine.Cnv.Width()), 1.0, 100.0)
	engine.Camera = camera.InitCamera()
	engine.LightCamera = camera.InitLightCamera()
	// engine.LightCamera.Matrix, _ = camera.CreateCamera(engine.LightCamera, engine.ProjMatrix, 0)
	engine.LightCamera.Matrix = [4][4]float64{
		{-0.023098907820484, 0, -1.0096703061025305, -0.9995736030415052},
		{0, 1, 0, 0},
		{0.7907341181203336, 0, -0.02949446697099902, -0.02919952230128903},
		{0.138593446922904, -3, 5.047920826514173, 5.997441618249031},
	}



	rast := canvas.NewRasterFromImage(cnv.Image())

	img := container.New(layout.NewGridWrapLayout(fyne.NewSize(width, height)), rast)

	// боковое меню
	menu := menu.MenuEx(w, a, img, engine, cnv)

	menuColumn := container.New(layout.NewGridWrapLayout(fyne.NewSize(310, height)), menu)
	form := container.New(layout.NewFormLayout(), menuColumn, img)
	w.SetContent(form)
	//w.SetContent(container.New(layout.NewCenterLayout(), canvasM))

	w.Resize(fyne.NewSize(width, height))
	w.ShowAndRun()
}
