package drawobj

import (

	"../inter"
)

func DrawSceneEx(engine *inter.MyGraphicsEngine) {
	copySlice := make([]inter.Square, len(engine.Object))
	copy(copySlice, engine.Object)
	createSceneObj(copySlice, engine.Camera.Matrix)

	DrawSquare(engine, copySlice)
}

func createSceneObj(slice []inter.Square, matrix [4][4]float64) {
	for i := 0; i < len(slice); i++ {
		// fmt.Println(tmp.Triagle1)
		slice[i].Triagle1 = modificationPoligon(slice[i].Triagle1, matrix)
		// fmt.Println(tmp.Triagle1)
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

