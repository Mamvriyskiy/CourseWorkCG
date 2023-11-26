package polygon

import (
	//"fmt"
	"image/color"

	"../inter"
)

func createSquareUpDown(x1, y1, z1, x2, y2, z2 float64, clr color.Color) inter.Square {
	var square inter.Square

	square.Triagle1 = createTriangleBTA(x1, y1, z1, x2, y2, z2, clr)
	square.Triagle2 = createTriangleBTB(x1, y1, z1, x2 , y2, z2, clr)

	square.Object = 0

	return square
}

func createSquareLeftRight(x1, y1, z1, x2, y2, z2 float64, clr color.Color) inter.Square {
	var square inter.Square

	square.Triagle1 = createTriangleLRA(x1, y1, z1, x2, y2, z2, clr)
	square.Triagle2 = createTriangleLRB(x1, y1, z1, x2 , y2, z2, clr)

	square.Object = 0

	return square
}

func createSquareBackFront(x1, y1, z1, x2, y2, z2 float64, clr color.Color) inter.Square {
	var square inter.Square

	square.Triagle1 = createTriangleBBA(x1, y1, z1, x2, y2, z2, clr)
	square.Triagle2 = createTriangleBBB(x1, y1, z1, x2, y2, z2, clr)

	square.Object = 0

	return square
}

func CreateSceneEx(a, b int) ([]inter.Square, float64) {
	countSquare := a * b * 2 + a * 2 + b * 2
	// fmt.Println(countSquare)

	sizeA := (float64(a) / 40.0) * 10.0
	sizeB := (float64(b) / 40.0) * 10.0

	stepA := (sizeA * 2.0) / float64(a)
	stepB := (sizeB * 2.0) / float64(b)

	scene := createScene(sizeA, sizeB, stepA, stepB, countSquare)	
	
	return scene, stepA
}

func createScene(sizeA, sizeB, stepA, stepB float64, countSquare int) []inter.Square {
	scene := make([]inter.Square, countSquare)
	color_cron := color.NRGBA{0, 154, 23, 255}
	k := 0

	numX := 1
	numY := 1
	for x := -sizeA; x < sizeA; x += stepA {
		for z := -sizeB; z < sizeB; z += stepB {
			//низ и вверх
			scene[k] = createSquareUpDown(x, 0, z, x + stepA, 0, z + stepB, color_cron)	
			k++
			scene[k] = createSquareUpDown(x, 0.2, z, x + stepA, 0.2, z + stepB, color_cron)	
			scene[k].NumberX = numX
			scene[k].NumberY = numY
			numY++
			k++
		}
		numY = 1
		numX++
	}

	//левая и правая
	for z := -sizeB; z < sizeB; z += stepB {
		scene[k] = createSquareLeftRight(-sizeA, 0, z, -sizeA, 0.2, z + stepB, color_cron)
		k++
		scene[k] = createSquareLeftRight(sizeA, 0, z, sizeA, 0.2, z + stepB, color_cron)
		k++
	}

	//задняя и передняя
	for x := -sizeA; x < sizeA; x += stepA {
		scene[k] = createSquareBackFront(x, 0, sizeB, x + stepA, 0.2, sizeB, color_cron)
		k++
		scene[k] = createSquareBackFront(x, 0, -sizeB, x + stepA, 0.2, -sizeB, color_cron)
		k++
	}

	return scene
}
