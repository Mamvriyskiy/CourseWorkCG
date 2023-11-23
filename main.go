package main

import (
	//"fmt"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	//"fyne.io/fyne/v2/widget"

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
	engine := graphics.NewMyGraphicsEngine(cnv, false)
	engine.ProjMatrix = mathfunc.MakeFovProjectionM(90.0, float64(engine.Cnv.Height())/float64(engine.Cnv.Width()), 1.0, 100.0)
	fmt.Println(engine.ProjMatrix)
	engine.Camera = camera.InitCamera()
	engine.LightCamera = camera.InitLightCamera()

	rast := canvas.NewRasterFromImage(cnv.Image())

	img := container.New(layout.NewGridWrapLayout(fyne.NewSize(width, height)), rast)

	rast.OnTapped = func(e *fyne.PointEvent) {
		x, y := int(e.Position.X), int(e.Position.Y)
		fmt.Printf("Mouse clicked at coordinates (%d, %d)\n", x, y)
	}

	// боковое меню
	menu := menu.MenuEx(w, a, img, engine)

	menuColumn := container.New(layout.NewGridWrapLayout(fyne.NewSize(310, height)), menu)
	form := container.New(layout.NewFormLayout(), menuColumn, img)
	w.SetContent(form)

	w.Resize(fyne.NewSize(width, height))
	w.ShowAndRun()
}
