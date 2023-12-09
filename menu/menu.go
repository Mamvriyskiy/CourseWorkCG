package menu

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"../polygon"
	"../inter"
	"../camera"
	"../drawobj"
	"../graphics"
	"../mathfunc"
	"image/color"
)

func MenuEx(w fyne.Window, a fyne.App, img *fyne.Container, engine *inter.MyGraphicsEngine, cnv graphics.ImageCanvas) *fyne.Container {
	//var scene []inter.Square
	// Ввод размерности сцены
	var entryA, entryB *widget.Entry
	{
		entryA = widget.NewEntry()
		entryA.SetPlaceHolder("от 1 до 100")
		entryB = widget.NewEntry()
		entryB.SetPlaceHolder("от 1 до 100")
	}


	// Кнопки создания и отчистки сцены
	var createSceneButton, clearButton *widget.Button
	{
		createSceneButton = widget.NewButton("Создать сцену", func() {
		sizeSceneA, sizeSceneB, errA, errB := CheckEntrySize(entryA.Text, entryB.Text)
		if errA.message == "" && errB.message == "" {
			engine.Object, engine.Step = polygon.CreateSceneEx(sizeSceneA, sizeSceneB)
			engine.Camera.Matrix = camera.CreateCamera(engine.Camera, engine.ProjMatrix)
			drawobj.DrawSceneEx(engine)
			img.Refresh()
		} else {
			entryA.SetText("")
			entryA.SetPlaceHolder(errA.Error())

			entryB.SetText("")
			entryB.SetPlaceHolder(errB.Error())
		}
		})

		clearButton = widget.NewButton("Очистить", func() {
			engine.ZBuf = nil
			engine.Cnv.Fill(color.RGBA{3, 215, 252, 140})
			engine.ZBuf = graphics.CreateZBuf(engine.Cnv.Height(), engine.Cnv.Width())
	
			img.Refresh()
		})
	}

	// Надписи в меню
	var labelTextCamera, spacerText, labelTextObj, labelTextRotate, labelTextControl *widget.Label
	{
		labelTextCamera = widget.NewLabel("Управление камерой")
		labelTextCamera.TextStyle = fyne.TextStyle{Bold: true, Italic: false}

		spacerText = widget.NewLabel("")

		labelTextObj = widget.NewLabel("Выберите объект")
		labelTextObj.TextStyle = fyne.TextStyle{Bold: true, Italic: false}

		labelTextRotate = widget.NewLabel("Действие")
		labelTextRotate.TextStyle = fyne.TextStyle{Bold: true, Italic: false}

		labelTextControl = widget.NewLabel("Управление сценой")
		labelTextControl.TextStyle = fyne.TextStyle{Bold: true, Italic: false}
	}
	spacer := container.NewCenter(spacerText)
	text := container.NewCenter(labelTextCamera)

	labelTextControlScene := container.NewCenter(labelTextControl)

	// Управление камерой через меню
	var buttonUpCamera, buttonDownCamera, buttonLeftCamera, buttonRightCamera, 
	buttonZoomInCamera, buttonZoomOutCamera, buttonRotateLeftCamera, buttonRotateRightCamera *widget.Button
	{
		buttonUpCamera = widget.NewButtonWithIcon("", theme.MoveUpIcon(), func() {
			engine.Camera.VCamera.Add(mathfunc.MakeVec3(0, 1, 0))
			renderScene(engine)
			img.Refresh()
		})

		buttonDownCamera = widget.NewButtonWithIcon("", theme.MoveDownIcon(), func() {
			engine.Camera.VCamera.Add(mathfunc.MakeVec3(0, -1, 0))
			renderScene(engine)
			img.Refresh()
		})

		buttonLeftCamera = widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
			engine.Camera.VCamera.Add(mathfunc.MakeVec3(-1, 0, 0))
			renderScene(engine)
			img.Refresh()
		})

		buttonRightCamera = widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
			engine.Camera.VCamera.Add(mathfunc.MakeVec3(1, 0, 0))
			renderScene(engine)
			img.Refresh()
		})

		buttonZoomInCamera = widget.NewButton("W", func() {
			engine.Camera.VCamera.Add(engine.Camera.VForward)
			renderScene(engine)
			img.Refresh()
		})

		buttonZoomOutCamera = widget.NewButton("S", func() {
			engine.Camera.VCamera.Sub(engine.Camera.VForward)
			renderScene(engine)
			img.Refresh()
		})

		buttonRotateLeftCamera = widget.NewButton("A", func() {
			engine.Camera.FYaw += 0.1
			renderScene(engine)
			img.Refresh()
		})

		buttonRotateRightCamera = widget.NewButton("D", func() {
			engine.Camera.FYaw -= 0.1
			renderScene(engine)
			img.Refresh()
		})
	}

	//Управление камерой через клавиатуру
	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == "Right" {
			engine.Camera.VCamera.Add(mathfunc.MakeVec3(1, 0, 0))
			renderScene(engine)
			img.Refresh()
		}
		if k.Name == "Left" {
			engine.Camera.VCamera.Add(mathfunc.MakeVec3(-1, 0, 0))
			renderScene(engine)
			img.Refresh()
		}
		if k.Name == "Up" {
			engine.Camera.VCamera.Add(mathfunc.MakeVec3(0, 1, 0))
			renderScene(engine)
			img.Refresh()
		}
		if k.Name == "Down" {
			engine.Camera.VCamera.Add(mathfunc.MakeVec3(0, -1, 0))
			renderScene(engine)
			img.Refresh()
		}
		if k.Name == "W" {
			engine.Camera.VCamera.Add(engine.Camera.VForward)
			renderScene(engine)
			img.Refresh()
		}
		if k.Name == "S" {
			engine.Camera.VCamera.Sub(engine.Camera.VForward)
			renderScene(engine)
			img.Refresh()
		}
		if k.Name == "A" {
			engine.Camera.FYaw += 0.1
			renderScene(engine)
			img.Refresh()
		}
		if k.Name == "D" {
			engine.Camera.FYaw -= 0.1
			renderScene(engine)
			img.Refresh()
		}
	})

	cntCamera := container.NewCenter(container.New(layout.NewHBoxLayout(), buttonUpCamera, 
		buttonDownCamera, buttonLeftCamera, buttonRightCamera, buttonZoomInCamera, buttonZoomOutCamera, 
		buttonRotateLeftCamera, buttonRotateRightCamera))

	//Выбор номер яячейки фигуры
	var entryX, entryY *widget.Entry
	{
		entryX = widget.NewEntry()
		entryX.SetPlaceHolder("номер по X")
		entryY = widget.NewEntry()
		entryY.SetPlaceHolder("номер по Y")
	}

	var entryX2, entryY2 *widget.Entry
	{
		entryX2 = widget.NewEntry()
		entryX2.SetPlaceHolder("номер по X")
		entryY2 = widget.NewEntry()
		entryY2.SetPlaceHolder("номер по Y")
	}

	// Выбор фигуры
	var radioGroupObj, radioRotateObj *widget.RadioGroup
	var radioButton *widget.Button
	{
		radioGroupObj = widget.NewRadioGroup([]string{"станция", "головной вагон", 
														"вагон", "закругленные рельсы", 
														"прямые рельсы", "дерево"}, func(s string) {
		})

		radioRotateObj = widget.NewRadioGroup([]string{"создать", "удалить", 
			"повернуть"}, func(s string) {
		})

		radioButton = widget.NewButton("Создать", func() {

			numA, numB, errA, errB := CheckEntrySize(entryX.Text, entryY.Text)
			var errA2, errB2 MyError
			var numA2, numB2 int
			if radioGroupObj.Selected == "головной вагон" || radioGroupObj.Selected == "вагон" {
				numA2, numB2, errA2, errB2 = CheckEntrySize(entryX2.Text, entryY2.Text)
			}
			if errA.message == "" && errB.message == "" && errA2.message == "" && errB2.message == ""{
				engine.Object = polygon.CreateObjectForScene(engine.Object, numA, numB, numA2, numB2, radioGroupObj.Selected, radioRotateObj.Selected, engine.Step)
				renderScene(engine)
				img.Refresh()
			} else {
				entryX.SetText("")
				entryX.SetPlaceHolder(errA.Error())
	
				entryY.SetText("")
				entryY.SetPlaceHolder(errB.Error())
			}
			fmt.Println()
		})
	}


	choiceObj := container.New(layout.NewVBoxLayout(), labelTextObj, radioGroupObj)
	rotateObj := container.New(layout.NewVBoxLayout(), labelTextRotate, radioRotateObj, entryX, entryY, entryX2, entryY2)

	settingsObj := container.New(layout.NewHBoxLayout(), choiceObj, rotateObj)

	// Управление сценой
	var buttonUpScene, buttonDownScene, buttonLeftScene, buttonRightScene, 
	buttonRotateLeftScene, buttonRotateRightScene *widget.Button
	{
		buttonUpScene = widget.NewButton("T", func() {
			// Действие, выполняемое при нажатии на кнопку
		})

		buttonDownScene = widget.NewButton("F", func() {
			// Действие, выполняемое при нажатии на кнопку
		})

		buttonLeftScene = widget.NewButton("G", func() {
			// Действие, выполняемое при нажатии на кнопку
		})

		buttonRightScene = widget.NewButton("H", func() {
			// Действие, выполняемое при нажатии на кнопку
		})

		buttonRotateLeftScene = widget.NewButton("L", func() {
			// Действие, выполняемое при нажатии на кнопку
		})

		buttonRotateRightScene = widget.NewButton("R", func() {
			// Действие, выполняемое при нажатии на кнопку
		})
	}

	cntScene := container.NewCenter(container.New(layout.NewHBoxLayout(), buttonUpScene, 
		buttonDownScene, buttonLeftScene, buttonRightScene,
		buttonRotateLeftScene, buttonRotateRightScene))


	// Смена тем(светлый, темный)
	var btnDark, btnLight *widget.Button
	{	
		btnDark = widget.NewButton("Темная тема", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		})	

		btnLight = widget.NewButton("Светлая тема", func() {
			a.Settings().SetTheme(theme.LightTheme())
		})	
	}

	menu := container.New(layout.NewVBoxLayout(), 
		entryA, entryB, 
		createSeparatorLine(1), createSceneButton, clearButton, 
		createSeparatorLine(3), text,cntCamera,  
		createSeparatorLine(3), settingsObj, radioButton,
		createSeparatorLine(3), labelTextControlScene, cntScene,
		createSeparatorLine(3), 
		btnDark, btnLight)

	menu = container.New(layout.NewHBoxLayout(), menu, createSeparatorLine(3))

	fmt.Println(spacer)

	return menu
}

func createSeparatorLine(width float32) *canvas.Line {
    line := canvas.NewLine(theme.SeparatorColor())
    line.StrokeWidth = width
    line.Resize(fyne.NewSize(100, 1))
    return line
}

func renderScene(engine *inter.MyGraphicsEngine) {
	engine.ZBuf = nil
	engine.ZBuf = graphics.CreateZBuf(engine.Cnv.Width(), engine.Cnv.Height())

	engine.Cnv.Fill(color.RGBA{3, 215, 252, 140})

	engine.Camera.Matrix = camera.CreateCamera(engine.Camera, engine.ProjMatrix)
	drawobj.DrawSceneEx(engine)
}
