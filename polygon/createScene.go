package polygon

import (
	"fmt"
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

func CreateSceneEx(a, b int) []inter.Square {
	// countSquare := a * b * 2 + a * 2 + b * 2
	// fmt.Println(countSquare)

	sizeA := (float64(a) / 40.0) * 10.0
	sizeB := (float64(b) / 40.0) * 10.0

	size := max(sizeA, sizeB)

	step := (size * 2.0) / max(float64(a), float64(b))

	//TO DO: change size scene. Lenght && Width

	scene := createScene(size, step)	
	
	
	return scene
}

func createScene(size, step float64) []inter.Square {
	scene := make([]inter.Square, 80)
	color_cron := color.NRGBA{0, 154, 23, 255}
	k := 0

	for x := -size; x < size; x += step {
		for z := -size; z < size; z += step {
			//низ и вверх
			scene[k] = createSquareUpDown(x, 0, z, x + step, 0, z + step, color_cron)	
			k++
			scene[k] = createSquareUpDown(x, 0.2, z, x + step, 0.2, z + step, color_cron)	
			k++
		}
	}

	//левая и правая
	for z := -size; z < size; z += step {
		scene[k] = createSquareLeftRight(-size, 0, z, -size, 0.2, z + step, color_cron)
		k++
		scene[k] = createSquareLeftRight(size, 0, z, size, 0.2, z + step, color_cron)
		k++
	}

	//задняя и передняя
	for x := -size; x < size; x += step {
		scene[k] = createSquareBackFront(x, 0, size, x + step, 0.2, size, color_cron)
		k++
		scene[k] = createSquareBackFront(x, 0, -size, x + step, 0.2, -size, color_cron)
		k++
	}
	
	return scene
}
