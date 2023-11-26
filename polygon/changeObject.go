package polygon

import (
	//"fmt"
	"fmt"
	"image/color"

	"../inter"
)

func CreateObjectForScene(objects []inter.Square, a, b int, obj, naprav string, step float64) []inter.Square {
	for i := 0; i < len(objects); i++ {
		if objects[i].NumberX == a && objects[i].NumberY == b {
			fmt.Println(objects[i])
			objects[i].Object = len(objects)
			objects = createObj(objects, obj, naprav, step, objects[i].Triagle1.P1.X, objects[i].Triagle1.P1.Z)
			break
		}
	}

	return objects
}

func createObj(square []inter.Square, obj, npr string, step float64, x, z float64) []inter.Square {
	switch obj {
		case "станция":
			fmt.Println("+")
		case "головной вагон":
			fmt.Println("+")
		case "вагон":
			fmt.Println("+")
		case "закругленные рельсы":
			fmt.Println("+")
		case "прямые рельсы":
			square = append(square, createStraightRails(x, z, step)...)
		case "дерево":
			square = append(square, createTree(x, z, step)...)
	}

	return square
}

func createStraightRails(x, z, step float64) []inter.Square {
	slice := make([]inter.Square, 3)
	railsColor := color.NRGBA{107, 106, 104, 255}
	sleepersColor := color.NRGBA{74, 56, 1, 255}

	slice[0] = createSquareUpDown(x + step / 3, 0.203, z, x + step / 3 * 2, 0.203, z + step, railsColor)
	slice[0].Object = 100

	slice[1] = createSquareUpDown(x + step / 4, 0.201, z, x + step / 4 * 3, 0.201, z + step / 4, sleepersColor)
	slice[1].Object = 101

	slice[2] = createSquareUpDown(x + step / 4, 0.201, z + step / 4 * 3, x + step / 4 * 3, 0.201, z + step, sleepersColor)
	slice[2].Object = 101

	return slice
}

func createTree(x, z, step float64) []inter.Square {
	color_stolb := color.NRGBA{125, 85, 4, 255}
	color_cron := color.NRGBA{0, 154, 23, 255}
	k := 0
	slice := make([]inter.Square, 15)
	// Дерево
	//передняя часть
	slice[k] = createSquareBackFront(x + step / 3, 1, z + step / 3, x + step / 3 * 2, 0.2, z + step / 3, color_stolb)
	k++

	//задняя часть
	slice[k] = createSquareBackFront(x + step / 3, 1, z + step / 3 * 2, x + step / 3 * 2, 0.2, z + step / 3 * 2, color_stolb)
	k++

	//левая часть
	slice[k] = createSquareLeftRight(x + step / 3, 0.2, z + step / 3 * 2, x + step / 3, 1, z + step / 3, color_stolb)
	k++

	//правая часть
	slice[k] = createSquareLeftRight(x + step / 3 * 2, 0.2, z + step / 3 * 2, x + step / 3 * 2, 1, z + step / 3, color_stolb)
	k++

	//нижняя часть кроны
	//низ
	slice[k] = createSquareUpDown(x + step / 6, 1, z + step / 6, x + step - step / 6, 1, z + step - step / 6, color_cron)
	k++

	// вверх
	slice[k] = createSquareUpDown(x + step / 6, 1.2, z + step / 6, x + step - step / 6, 1.2, z + step - step / 6, color_cron)
	k++

	// передняя
	slice[k] = createSquareBackFront(x + step / 6, 1, z + step / 6, x + step - step / 6, 1.2, z + step / 6, color_cron)
	k++

	// задняя
	slice[k] = createSquareBackFront(x + step / 6, 1, z + step - step / 6, x + step - step / 6, 1.2, z + step - step / 6, color_cron)
	k++
	
	//левая
	slice[k] = createSquareLeftRight(x + step / 6, 1, z + step / 6, x + step / 6, 1.2, z + step - step / 6, color_cron)
	k++

	//правая
	slice[k] = createSquareLeftRight(x + step - step / 6, 1, z + step / 6, x + step - step / 6, 1.2, z + step - step / 6, color_cron)
	k++

	//верхняя часть кроны
	//передняя часть
	slice[k] = createSquareBackFront(x + step / 3, 1.2, z + step / 3, x + step / 3 * 2, 1.4, z + step / 3, color_cron)
	k++

	//задняя часть
	slice[k] = createSquareBackFront(x + step / 3, 1.2, z + step / 3 * 2, x + step / 3 * 2, 1.4, z + step / 3 * 2, color_cron)
	k++

	//левая часть
	slice[k] = createSquareLeftRight(x + step / 3, 1.2, z + step / 3 * 2, x + step / 3, 1.4, z + step / 3, color_cron)
	k++

	//правая часть
	slice[k] = createSquareLeftRight(x + step / 3 * 2, 1.2, z + step / 3 * 2, x + step / 3 * 2, 1.4, z + step / 3, color_cron)
	k++
	//вернхяя часть
	slice[k] = createSquareUpDown(x + step / 3, 1.4, z + step / 3, x + step / 3 * 2, 1.4, z + step / 3 * 2, color_cron)
	k++

	return slice
}

