package mathfunc

import (
	"math"
	"../inter"
)

func MakeFovProjectionM(fov, ar, n, f float64) [4][4]float64 {
	var proj [4][4]float64

	fovRad := fov * math.Pi / 180.0
	w := 1.0 / math.Tan(fovRad / 2.0)
	h := w * ar

	proj[0][0] = h
	proj[1][1] = w
	proj[2][2] = f / (f - n)
	proj[3][2] = -n * f / (f - n)
	proj[2][3] = 1.0

	return proj
}

func MulVecMat(vec inter.Vec4, m [4][4]float64) inter.Vec4 {
	var res inter.Vec4

	res.X = vec.X*m[0][0] + vec.Y*m[1][0] + vec.Z*m[2][0] + m[3][0]
	res.Y = vec.X*m[0][1] + vec.Y*m[1][1] + vec.Z*m[2][1] + m[3][1]
	res.Z = vec.X*m[0][2] + vec.Y*m[1][2] + vec.Z*m[2][2] + m[3][2]
	res.W = vec.X*m[0][3] + vec.Y*m[1][3] + vec.Z*m[2][3] + m[3][3]

	return res
}

func MakeRotationYM(fAngleRad float64) [4][4]float64 {
	var m[4][4]float64
	m[0][0] = math.Cos(fAngleRad)
	m[2][2] = math.Cos(fAngleRad)
	m[0][2] = math.Sin(fAngleRad)
	m[2][0] = -math.Sin(fAngleRad)
	m[1][1] = 1.0
	m[3][3] = 1.0

	return m
}

func MakePointAtM(pos, target, up inter.Vec3) [4][4]float64 {
	newForward := Vec3Diff(target, pos)
	newForward = Normalize(newForward)

	a := Vec3Mul(newForward, DotProduct(up, newForward))
	newUp := Vec3Diff(up, a)
	newUp = Normalize(newUp)

	newRight := CrossProduct(newUp, newForward)

	var m [4][4]float64
	m[0][0] = newRight.X
	m[0][1] = newRight.Y
	m[0][2] = newRight.Z
	m[0][3] = 0.0

	m[1][0] = newUp.X
	m[1][1] = newUp.Y
	m[1][2] = newUp.Z
	m[1][3] = 0.0

	m[2][0] = newForward.X
	m[2][1] = newForward.Y
	m[2][2] = newForward.Z
	m[2][3] = 0.0

	m[3][0] = pos.X
	m[3][1] = pos.Y
	m[3][2] = pos.Z
	m[3][3] = 1.0

	return m
}

func InverseTranslationM(m [4][4]float64) [4][4]float64 {
	var res [4][4]float64

	res[0][0] = m[0][0]
	res[0][1] = m[1][0]
	res[0][2] = m[2][0]
	res[0][3] = 0.0

	res[1][0] = m[0][1]
	res[1][1] = m[1][1]
	res[1][2] = m[2][1]
	res[1][3] = 0.0

	res[2][0] = m[0][2]
	res[2][1] = m[1][2]
	res[2][2] = m[2][2]
	res[2][3] = 0.0

	res[3][0] = -(m[3][0]*res[0][0] + m[3][1]*res[1][0] + m[3][2]*res[2][0])
	res[3][1] = -(m[3][0]*res[0][1] + m[3][1]*res[1][1] + m[3][2]*res[2][1])
	res[3][2] = -(m[3][0]*res[0][2] + m[3][1]*res[1][2] + m[3][2]*res[2][2])
	res[3][3] = 1.0

	return res
}

func MulMatrices(m1, m2 [4][4]float64) [4][4]float64 {
	var m [4][4]float64

	for c := 0; c < 4; c++ {
		for r := 0; r < 4; r++ {
			m[r][c] = m1[r][0] * m2[0][c] + m1[r][1] * m2[1][c] + m1[r][2] * m2[2][c] + m1[r][3] * m2[3][c]
		}
	}

	return m
}
