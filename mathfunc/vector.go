package mathfunc

import (
	"../inter"
	"math"
)

func MakeVec3(x, y, z float64) inter.Vec3 {
	var vec inter.Vec3

	vec.X, vec.Y, vec.Z = x, y, z

	return vec
}

func MakeVec4(x, y, z float64, w... float64) inter.Vec4 {
	var vec4 inter.Vec4
	vec4.Vec3 = MakeVec3(x, y, z)

	if len(w) != 0 {
		vec4.W = w[0]
	} else {
		vec4.W = 1.0
	}

	return vec4
}

func Vec3Sum(a, b inter.Vec3) inter.Vec3 {
	Add(&a, b)

	return a
}

func Add(a *inter.Vec3, b inter.Vec3) {
	(*a).X += b.X
	(*a).Y += b.Y
	(*a).Z += b.Z
}

func Vec3Mul(a inter.Vec3, k float64) inter.Vec3 {
	Mul(&a, k)
	return a
}

func Mul(a *inter.Vec3, k float64){
	(*a).X *= k
	(*a).Y *= k
	(*a).Z *= k
}

func Sub(a *inter.Vec3, b inter.Vec3) {
	(*a).X -= b.X
	(*a).Y -= b.Y
	(*a).Z -= b.Z
}

func Normalize(vec inter.Vec3) inter.Vec3 {
	Div(&vec, Length(vec))

	return vec
}

func Div(a *inter.Vec3, k float64) {
	if k != 0 {
		(*a).X /= k
		(*a).Y /= k
		(*a).Z /= k
	}
}

func Length(vec inter.Vec3) float64 {
	length := math.Sqrt(vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z)

	return length
}

func DotProduct(a, b inter.Vec3) float64 {
	res := a.X * b.X + a.Y * b.Y + a.Z * b.Z

	return res
}

func CrossProduct(a, b inter.Vec3) inter.Vec3 {
	var vec inter.Vec3

	vec.X = (a.Y)*(b.Z) - (a.Z)*(b.Y)
	vec.Y = (a.Z)*(b.X) - (a.X)*(b.Z)
	vec.Z = (a.X)*(b.Y) - (a.Y)*(b.X)

	return vec
}

func Vec4Diff(a, b inter.Vec4) inter.Vec4 {
	a.Sub(b)

	return a
}

func Vec3Diff(a, b inter.Vec3) inter.Vec3 {
	a.Sub(b)

	return a
}
