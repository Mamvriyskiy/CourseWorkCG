package polygon

import (
	"math"

	"../inter"
)	

func rotateObj(object []inter.Square, i int, angle float64) []inter.Square {
	startRotate:= object[i].Object
	endRotate := object[i].Object + object[i].TypeObj
	radian := angleToRadians(angle)

	for j := startRotate; j < endRotate; j++ {
		object[j].Triagle1 = rotatePolygonY(object[j].Triagle1, radian)
		object[j].Triagle2 = rotatePolygonY(object[j].Triagle2, radian)
	}

	return object
}

func angleToRadians(angle float64) float64 {
	return angle * math.Pi / 180
}	

func rotatePolygonY(triagle inter.Polygon, angle float64) inter.Polygon {
	triagle.P1 = rotateVec4(triagle.P1, angle)
	triagle.P2 = rotateVec4(triagle.P2, angle)
	triagle.P3 = rotateVec4(triagle.P3, angle)

	return triagle
}

func rotateVec4(vec inter.Vec4, angle float64) inter.Vec4 {
	newX := vec.X * math.Cos(angle) + vec.Z * math.Sin(angle)
	newZ := vec.Z * math.Cos(angle) - vec.X * math.Sin(angle)

	vec.X = newX
	vec.Z = newZ

	return vec
}
