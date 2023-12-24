package drawobj

import (
	"../inter"
	"time"
	"fmt"
)

func DrawSceneEx(engine *inter.MyGraphicsEngine) {
	copySlice := make([]inter.Square, len(engine.Object))
	LightSlice := make([]inter.Square, len(engine.Object))
	copy(copySlice, engine.Object)
	copy(LightSlice, engine.Object)
	createSceneObj(copySlice, engine.Camera.Matrix)
	createSceneObj(LightSlice, engine.LightCamera.Matrix)

	//testFunc(engine.Object[0], engine.Camera.Matrix, engine.PreobMatrix, engine.LightCamera.Matrix)

	startTimeA := time.Now()
	DrawSquare(engine, LightSlice, false)
	endTimeA := time.Now()
	elapsedTimeA := endTimeA.Sub(startTimeA)

	startTimeB := time.Now()
	DrawSquare(engine, copySlice, true)
	endTimeB := time.Now()
	elapsedTimeB := endTimeB.Sub(startTimeB)
	
	fmt.Printf("Время выполнения в секундах: %.2f\n", elapsedTimeA.Seconds())
	fmt.Printf("Время выполнения в секундах: %.2f\n", elapsedTimeB.Seconds())
}

func createSceneObj(slice []inter.Square, matrix [4][4]float64) {
	for i := 0; i < len(slice); i++ {
		slice[i].Triagle1 = modificationPoligon(slice[i].Triagle1, matrix)
		slice[i].Triagle2 = modificationPoligon(slice[i].Triagle2, matrix)
	}
}

func modificationPoligon(polygon inter.Polygon, matrix [4][4]float64) inter.Polygon {
	polygon.P1 = modifVec4(polygon.P1, matrix)
	polygon.P2 = modifVec4(polygon.P2, matrix)
	polygon.P3 = modifVec4(polygon.P3, matrix)

	return polygon
}

func modifVec4(dot inter.Vec4, matrix [4][4]float64) inter.Vec4 {
	vec := [4]float64{dot.X, dot.Y, dot.Z, dot.W}
	vec = MulMatr(matrix, vec)
 
	dot.CreateVec4(vec[0], vec[1], vec[2], vec[3])

	return dot
}

func MulMatr(matrix [4][4]float64, vector [4]float64) [4]float64 {
	var result [4]float64

	for i := 0; i < 4; i++ {
		result[i] = vector[0] * matrix[0][i] + vector[1] * matrix[1][i] + vector[2] * matrix[2][i] + matrix[3][i]
	}

	return result
}

//func testFunc(polygon inter.Square, a, b, c [4][4]float64) {
	// 	p := polygon.Triagle1
	// 	k := polygon.Triagle1
	
	// 	fmt.Println("1)X:", p.P1.X, "Y:", p.P1.Y, "Началальное значение")
	
	// 	p = modificationPoligon(polygon.Triagle1, a)
	// 	fmt.Println("2)X:", p.P1.X, "Y:", p.P1.Y, "Камера пользвоателя")
	
	// 	p.P1.X = (p.P1.X / p.P1.W + 1) * 0.5 * float64(1024)
	// 	p.P1.Y = (1 - p.P1.Y / p.P1.W) * 0.5 * float64(840)
	// 	//fmt.Println("3)X:", p.P1.X, "Y:", p.P1.Y)
	
	// 	p.P1.X  = (p.P1.X  / float64(1024) / 0.5 - 1) * p.P1.W
	// 	p.P1.Y  = (p.P1.Y / float64(840) / 0.5 - 1) * (-p.P1.W)
	// 	fmt.Println("3)X:", p.P1.X, "Y:", p.P1.Y, "Вернуть к камере пользователя")
	
	// 	matr := [4]float64{p.P1.X, p.P1.Y, p.P1.Z, p.P1.W}
	// 	matr = Mulmatrices(matr, b)
	// 	fmt.Println("4)X:", matr[0], "Y:", matr[1], "Возврат к начальным")
	
	// 	result := Mulmatrices(matr, c) 
	// 	fmt.Println("5)X:", result[0], "Y:", result[1], result[2], result[3], "Переход к свету")
	
	
	// 	resultK := modificationPoligon(k, c)
	// 	fmt.Println("6)X:", resultK.P1.X, "Y:", resultK.P1.Y, resultK.P1.Z, resultK.P1.W, "Получение света")
	
	// 	// x = (matr[0] / matr[3] + 1) * 0.5 * float64(cnv.Cnv.Width())
	
	// 	// y = (1 - matr[1] / matr[3]) * 0.5 * float64(cnv.Cnv.Height())
	
	// 	// px := int(math.Round(x))
	// 	// py := int(math.Round(y))
	// }

