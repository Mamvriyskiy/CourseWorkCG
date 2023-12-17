package polygon

import (
	//"fmt"
	"fmt"
	"image/color"

	"../inter"
)

func CreateObjectForScene(objects []inter.Square, a, b int, obj, do string, step float64) []inter.Square {
	for i := 0; i < len(objects); i++ {
		if objects[i].NumberX == a && objects[i].NumberY == b {
			if do == "создать" {
				objects[i].Object = len(objects)
				objects = createObj(objects, obj, step, objects[i].Triagle1.P1.X, objects[i].Triagle1.P1.Z, i)
			} else if do == "удалить" {
				fmt.Println("Len:", len(objects))
				objects = deleteObj(objects, i)
				fmt.Println("Len:", len(objects))
			} else {
				rotateObj(objects[i], 90)
				fmt.Println("Повернуть");
			}
			fmt.Println(len(objects))
			break
		}
	}

	return objects
}

func deleteObj(square []inter.Square, i int) []inter.Square {
	startDelet := square[i].Object
	endDelet := square[i].Object + square[i].TypeObj
	newSquare := append(square[: startDelet], square[endDelet : ]...)

	return newSquare
}

func createObj(square []inter.Square, obj string, step float64, x, z float64, k int) []inter.Square {
	switch obj {
		case "станция":
			square = append(square, createStation(x, z, step)...)
			square[k].TypeObj = inter.STATION
		case "вагон":
			square[k].TypeObj = inter.TRAIN
			square = append(square, createTrain(x, z, step)...)
		case "головной вагон":
			square[k].TypeObj = inter.TRAINHEAD
			square = append(square, createTrainHead(x, z, step)...)
		case "закругленные рельсы":
			square = append(square, createCircleRails(x, z, step)...)
		case "прямые рельсы":
			square = append(square, createStraightRails(x, z, step)...)
			square[k].TypeObj = inter.RAILS
		case "дерево":
			square = append(square, createTree(x, z, step)...)
			square[k].TypeObj = inter.TREE
	}

	return square
}

func createTrain(x, z, step float64) []inter.Square {
	slice := make([]inter.Square, 6)
	k := 0

	slice[k] = createSquareUpDown(x + step / 5, 0.205, z, x + step / 5 * 4, 0.205, z + step, color.NRGBA{245, 243, 240, 255})	
	k++
	slice[k] = createSquareUpDown(x + step / 5, 0.530, z, x + step / 5 * 4, 0.530, z + step, color.NRGBA{245, 243, 240, 255})
	k++

	slice[k] = createSquareLeftRight(x + step / 5, 0.205, z, x + step / 5, 0.530, z + step, color.NRGBA{245, 243, 240, 255})
	k++
	slice[k] = createSquareLeftRight(x + step / 5 * 4, 0.205, z, x + step / 5 * 4, 0.530, z + step, color.NRGBA{245, 243, 240, 255})
	k++

	slice[k] = createSquareBackFront(x + step / 5, 0.205, z, x + step / 5 * 4, 0.53, z, color.NRGBA{245, 243, 240, 255})
	k++
	slice[k] = createSquareBackFront(x + step / 5, 0.205, z + step, x + step / 5 * 4, 0.53, z + step, color.NRGBA{245, 243, 240, 255})
	k++

	return slice
}
	
func createTrainHead(x, z, step float64) []inter.Square {
	slice := make([]inter.Square, 10)
	k := 0

	slice[k].Triagle1 = createHead(x + step / 5, 0.255, z + step / 4, x + step / 5, 0.53, z + step, color.NRGBA{245, 243, 240, 255})
	slice[k].Triagle2 = createHead(x + step / 5 * 4, 0.255, z + step / 4, x + step / 5 * 4, 0.53, z + step, color.NRGBA{245, 243, 240, 255})
	k++

	slice[k] = createSquareLeftRight(x + step / 5, 0.205, z + step / 4, x + step / 5, 0.255, z + step, color.NRGBA{245, 243, 240, 255})
	k++
	slice[k] = createSquareLeftRight(x + step / 5 * 4, 0.205, z + step / 4, x + step / 5 * 4, 0.255, z + step, color.NRGBA{245, 243, 240, 255})
	k++

	slice[k] = createMirror(x + step / 5, 0.255, z + step / 4, x + step / 5 * 4, 0.53, z + step, color.NRGBA{245, 243, 240, 255})
	k++

	slice[k] = createMirror(x + step / 5 * 2, 0.336, z + step / 3, x + step / 5 * 3, 0.5, z + step / 5 * 4, color.Black)
	k++

	slice[k] = createSquareBackFront(x + step / 5, 0.205, z + step, x + step / 5 * 4, 0.255, z + step, color.NRGBA{245, 243, 240, 255})
	k++

	slice[k] = createSquareLeftRight(x + step / 5, 0.205, z, x + step / 5, 0.53, z + step / 4, color.NRGBA{245, 243, 240, 255})
	k++
	slice[k] = createSquareLeftRight(x + step / 5 * 4, 0.205, z, x + step / 5 * 4, 0.53, z + step / 4, color.NRGBA{245, 243, 240, 255})
	k++

	slice[k] = createSquareBackFront(x + step / 5, 0.205, z, x + step / 5 * 4, 0.53, z, color.NRGBA{245, 243, 240, 255})
	k++

	slice[k] = createSquareUpDown(x + step / 5, 0.53, z, x + step / 5 * 4, 0.53, z + step / 4, color.NRGBA{245, 243, 240, 255})
	k++

	return slice
}

func createMirror(x1, y1, z1, x2, y2, z2 float64, clr color.Color) inter.Square {
	var square inter.Square

	var triangle1 inter.Polygon

	triangle1.P1.CreateDot(x1, y1, z2, 1)
	triangle1.P2.CreateDot(x2, y2, z1, 1)
	triangle1.P3.CreateDot(x2, y1, z2, 1)
	triangle1.Color = clr

	var triangle2 inter.Polygon
	triangle2.P1.CreateDot(x2, y2, z1, 1)
	triangle2.P2.CreateDot(x1, y1, z2, 1)
	triangle2.P3.CreateDot(x1, y2, z1, 1)
	triangle2.Color = clr

	square.Triagle1 = triangle1
	square.Triagle2 = triangle2
	square.Object = 0
	return square
}

func createHead(x1, y1, z1, x2, y2, z2 float64, clr color.Color) inter.Polygon {
	var triangle1 inter.Polygon

	triangle1.P1.CreateDot(x2, y1, z2, 1)
	triangle1.P2.CreateDot(x1, y2, z1, 1)
	triangle1.P3.CreateDot(x1, y1, z1, 1)
	triangle1.Color = clr

	return triangle1
}

func createRoofStation(x1, y1, z1, x2, y2, z2 float64, clr color.Color) inter.Square {
	var square inter.Square

	var triangle1 inter.Polygon

	triangle1.P1.CreateDot(x1, y1, z1, 1)
	triangle1.P2.CreateDot(x2, y2, z2, 1)
	triangle1.P3.CreateDot(x1, y1, z2, 1)
	triangle1.Color = clr

	var triangle2 inter.Polygon
	triangle2.P1.CreateDot(x1, y1, z1, 1)
	triangle2.P2.CreateDot(x2, y2, z2, 1)
	triangle2.P3.CreateDot(x2, y2, z1, 1)
	triangle2.Color = clr

	square.Triagle1 = triangle1
	square.Triagle2 = triangle2
	square.Object = 0

	return square
}

func createStation(x, z, step float64) []inter.Square {
	slice := make([]inter.Square, 9)
	k := 0

	slice[k] = createSquareUpDown(x, 0.201, z, x + step, 0.201, z + step, color.NRGBA{74, 53, 56, 255})	
	k++
	slice[k] = createSquareUpDown(x, 0.28, z, x + step, 0.28, z + step, color.NRGBA{74, 53, 56, 255})
	k++

	slice[k] = createSquareLeftRight(x, 0.201, z, x, 0.28, z + step, color.NRGBA{74, 53, 56, 255})
	k++
	slice[k] = createSquareLeftRight(x + step, 0.201, z, x + step, 0.28, z + step, color.NRGBA{74, 53, 56, 255})
	k++
	
	slice[k] = createSquareBackFront(x, 0.201, z, x + step, 0.28, z, color.NRGBA{74, 53, 56, 255})
	k++
	slice[k] = createSquareBackFront(x, 0.201, z + step, x + step, 0.28, z + step, color.NRGBA{74, 53, 56, 255})
	k++

	slice[k] = createSquareLeftRight(x + step / 2, 0.28, z + step / 3, x + step / 2, 0.65, z + step / 3 * 2, color.NRGBA{219, 200, 202, 255})
	k++

	slice[k] = createRoofStation(x, 0.72, z, x + step / 2, 0.65, z + step,  color.NRGBA{10, 10, 140, 255})
	k++

	slice[k] = createRoofStation(x + step / 2, 0.65, z, x + step, 0.72, z + step,  color.NRGBA{10, 10, 140, 255})
	k++


	return slice
}

func createCircleRails(x, z, step float64) []inter.Square {
	slice := make([]inter.Square, 1)
	railsColor := color.NRGBA{107, 106, 104, 255}
	//sleepersColor := color.NRGBA{74, 56, 1, 255}

	triagle1 := createTriangleBTB(x + step / 3, 0.201, z , x + step / 3 * 2, 0.201, z + step, railsColor)
	triagle2 := createTriangleBTB(x, 0.201, z + step / 3, x + step, 0.201, z + step / 3 * 2, railsColor)

	slice[0].Triagle1 = triagle1
	slice[0].Triagle2 = triagle2
	slice[0].Object = 102

	return slice
}

func createStraightRails(x, z, step float64) []inter.Square {
	slice := make([]inter.Square, 3)
	railsColor := color.NRGBA{107, 106, 104, 255}
	sleepersColor := color.NRGBA{74, 56, 1, 255}

	slice[0] = createSquareUpDown(x + step / 3, 0.203, z, x + step / 3 * 2, 0.203, z + step, railsColor)
	slice[0].Object = 100

	slice[1] = createSquareUpDown(x + step / 4, 0.201, z, x + step / 4 * 3, 0.201, z + step / 3, sleepersColor)
	slice[1].Object = 101

	slice[2] = createSquareUpDown(x + step / 4, 0.201, z + step / 3 * 2, x + step / 4 * 3, 0.201, z + step, sleepersColor)
	slice[2].Object = 101

	return slice
}

func createTree(x, z, step float64) []inter.Square {
	color_stolb := color.NRGBA{125, 85, 4, 255}
	color_cron := color.NRGBA{6, 82, 22, 255}
	k := 0
	slice := make([]inter.Square, 15)
	// Дерево
	//передняя часть
	slice[k] = createSquareBackFront(x + step / 3, 0.6, z + step / 3, x + step / 3 * 2, 0.2, z + step / 3, color_stolb)
	k++

	//задняя часть
	slice[k] = createSquareBackFront(x + step / 3, 0.6, z + step / 3 * 2, x + step / 3 * 2, 0.2, z + step / 3 * 2, color_stolb)
	k++

	//левая часть
	slice[k] = createSquareLeftRight(x + step / 3, 0.2, z + step / 3 * 2, x + step / 3, 0.6, z + step / 3, color_stolb)
	k++

	//правая часть
	slice[k] = createSquareLeftRight(x + step / 3 * 2, 0.2, z + step / 3 * 2, x + step / 3 * 2, 0.6, z + step / 3, color_stolb)
	k++

	//нижняя часть кроны
	//низ
	slice[k] = createSquareUpDown(x + step / 6, 0.6, z + step / 6, x + step - step / 6, 0.6, z + step - step / 6, color_cron)
	k++

	// вверх
	slice[k] = createSquareUpDown(x + step / 6, 0.8, z + step / 6, x + step - step / 6, 0.8, z + step - step / 6, color_cron)
	k++

	// передняя
	slice[k] = createSquareBackFront(x + step / 6, 0.6, z + step / 6, x + step - step / 6, 0.8, z + step / 6, color_cron)
	k++

	// задняя
	slice[k] = createSquareBackFront(x + step / 6, 0.6, z + step - step / 6, x + step - step / 6, 0.8, z + step - step / 6, color_cron)
	k++
	
	//левая
	slice[k] = createSquareLeftRight(x + step / 6, 0.6, z + step / 6, x + step / 6, 0.8, z + step - step / 6, color_cron)
	k++

	//правая
	slice[k] = createSquareLeftRight(x + step - step / 6, 0.6, z + step / 6, x + step - step / 6, 0.8, z + step - step / 6, color_cron)
	k++

	//верхняя часть кроны
	//передняя часть
	slice[k] = createSquareBackFront(x + step / 3, 0.8, z + step / 3, x + step / 3 * 2, 1, z + step / 3, color_cron)
	k++

	//задняя часть
	slice[k] = createSquareBackFront(x + step / 3, 0.8, z + step / 3 * 2, x + step / 3 * 2, 1, z + step / 3 * 2, color_cron)
	k++

	//левая часть
	slice[k] = createSquareLeftRight(x + step / 3, 0.8, z + step / 3 * 2, x + step / 3, 1, z + step / 3, color_cron)
	k++

	//правая часть
	slice[k] = createSquareLeftRight(x + step / 3 * 2, 0.8, z + step / 3 * 2, x + step / 3 * 2, 1, z + step / 3, color_cron)
	k++
	//вернхяя часть
	slice[k] = createSquareUpDown(x + step / 3, 1, z + step / 3, x + step / 3 * 2, 1, z + step / 3 * 2, color_cron)
	k++

	return slice
}

