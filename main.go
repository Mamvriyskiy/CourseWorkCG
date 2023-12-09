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

	//canvasM := canvasWithMouseEvents()

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

// func canvasWithMouseEvents() *fyne.Container {
// 	c := canvas.NewRasterWithPixels(func(x, y, _, _ int) color.Color {
// 		return color.Black
// 	})

// 	c.OnTapped = func(ev *fyne.PointEvent) {
// 		fmt.Printf("Clicked at (%d, %d)\n", int(ev.Position.X), int(ev.Position.Y))
// 	}

// 	return container.New(layout.NewCenterLayout(), c)
// }
