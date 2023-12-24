package drawobj

import (
	//"fmt"
	"math"

	"../inter"
)

func checkCoords(x, y, z, w float64, cnv *inter.MyGraphicsEngine) bool{
	x = (x / float64(cnv.Cnv.Width()) / 0.5 - 1) * w
	y = (y / float64(cnv.Cnv.Height()) / 0.5 - 1) * (-w)

	matr := [4]float64{x, y, z, w}

	matr = Mulmatrices(matr, cnv.PreobMatrix)
	matr = Mulmatrices(matr, cnv.LightCamera.Matrix)

	x = (matr[0] / matr[3] + 1) * 0.5 * float64(cnv.Cnv.Width())

	y = (1 - matr[1] / matr[3]) * 0.5 * float64(cnv.Cnv.Height())

	px := int(math.Round(x))
	py := int(math.Round(y))

	if px >= 0 && py >= 0 && px < cnv.Cnv.Width() && py < cnv.Cnv.Height() {
		if matr[2] - cnv.SBuf[px][py] < 0.05 {
			return true	
		}
	}

	return false	
}	

func Mulmatrices(m1 [4]float64, m2 [4][4]float64) [4]float64 {
	var m [4]float64

	for c := 0; c < 4; c++ {
		m[c] = m1[0] * m2[0][c] + m1[1] * m2[1][c] + m1[2] * m2[2][c] + m1[3] * m2[3][c]
	}

	return m
}
