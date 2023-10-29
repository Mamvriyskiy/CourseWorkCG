package graphics

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/canvas"
)

func CreateMenu(a fyne.App) *fyne.Container {

	entryA := widget.NewEntry()
	entryA.SetPlaceHolder("от 1 до 100")
	entryB := widget.NewEntry()
	entryB.SetPlaceHolder("от 1 до 100")

	createSceneButton := widget.NewButton("Создать сцену", func() {
	})

	clearButton := widget.NewButton("Очистить", func() {
	})


	labelTextCamera := widget.NewLabel("Управление камерой")
	labelTextCamera.TextStyle = fyne.TextStyle{Bold: true, Italic: false}

	spacerText := widget.NewLabel("")
	spacer := container.NewCenter(spacerText)

	buttonUpCamera := widget.NewButtonWithIcon("", theme.MoveUpIcon(), func() {
        // Действие, выполняемое при нажатии на кнопку
    })

	buttonDownCamera := widget.NewButtonWithIcon("", theme.MoveDownIcon(), func() {
        // Действие, выполняемое при нажатии на кнопку
    })

	buttonLeftCamera := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
        // Действие, выполняемое при нажатии на кнопку
    })

	buttonRightCamera := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
        // Действие, выполняемое при нажатии на кнопку
    })

	text := container.NewCenter(labelTextCamera)

	btnDark := widget.NewButton("Темная тема", func() {
		a.Settings().SetTheme(theme.DarkTheme())
	})	

	btnLight := widget.NewButton("Светлая тема", func() {
		a.Settings().SetTheme(theme.LightTheme())
	})	

	menu := container.New(layout.NewVBoxLayout(), entryA, entryB, createSeparatorLine(1), createSceneButton, 
		clearButton, createSeparatorLine(3), text, buttonUpCamera,
		buttonLeftCamera, buttonRightCamera, buttonDownCamera,  createSeparatorLine(3), spacer, btnDark, btnLight)

	return menu
}

func createSeparatorLine(width float32) *canvas.Line {
    line := canvas.NewLine(theme.SeparatorColor()) // Используйте цвет темы или укажите свой цвет
    line.StrokeWidth = width // Ширина линии
    line.Resize(fyne.NewSize(100, 1)) // Размер линии (100 пикселей в ширину, 1 пиксель в высоту)
    return line
}
