package polygon

import (
	"math"

	"../inter"
)	

func calculateCenter(triangle inter.Polygon) (float64, float64) {
	return (triangle.P1.X + triangle.P2.X + triangle.P3.X) / 3, (triangle.P1.Z + triangle.P2.Z + triangle.P3.Z) / 3
}

func rotateObj(object []inter.Square, i int, angle float64) []inter.Square {
	startRotate:= object[i].Object
	endRotate := object[i].Object + object[i].TypeObj
	radian := angleToRadians(angle)

	//centerX, centerZ := calculateCenter(object[i].Triagle1)
	//fmt.Println(object[i].Triagle1.P3.X, object[i].Triagle2.P3.X, math.Abs(object[i].Triagle1.P3.X - object[i].Triagle2.P3.X) / 2)
	centerX := max(object[i].Triagle1.P3.X, object[i].Triagle2.P3.X) - math.Abs(object[i].Triagle1.P3.X - object[i].Triagle2.P3.X) / 2
	centerZ := max(object[i].Triagle1.P3.Z, object[i].Triagle2.P3.Z) - math.Abs(object[i].Triagle1.P3.Z - object[i].Triagle2.P3.Z) / 2
	//fmt.Println(centerZ, centerZ)

	for j := startRotate; j < endRotate; j++ {
		object[j].Triagle1 = translateCenter(object[j].Triagle1, -centerX, -centerZ)
		object[j].Triagle2 = translateCenter(object[j].Triagle2, -centerX, -centerZ)

		object[j].Triagle1 = rotatePolygonY(object[j].Triagle1, radian)
		object[j].Triagle2 = rotatePolygonY(object[j].Triagle2, radian)

		object[j].Triagle1 = translateCenter(object[j].Triagle1, centerX, centerZ)
		object[j].Triagle2 = translateCenter(object[j].Triagle2, centerX, centerZ)
	}

	if object[i].Object2 != 0 {
		startRotate = object[i].Object2
		endRotate = object[i].Object2 + object[i].TypeObj2

		for j := startRotate; j < endRotate; j++ {
			object[j].Triagle1 = translateCenter(object[j].Triagle1, -centerX, -centerZ)
			object[j].Triagle2 = translateCenter(object[j].Triagle2, -centerX, -centerZ)

			object[j].Triagle1 = rotatePolygonY(object[j].Triagle1, radian)
			object[j].Triagle2 = rotatePolygonY(object[j].Triagle2, radian)

			object[j].Triagle1 = translateCenter(object[j].Triagle1, centerX, centerZ)
			object[j].Triagle2 = translateCenter(object[j].Triagle2, centerX, centerZ)	
		}
	}

	return object
}

func translateCenter(object inter.Polygon, offsetX, offsetZ float64) inter.Polygon {
	object.P1.X += offsetX
	object.P2.X += offsetX
	object.P3.X += offsetX

	object.P1.Z += offsetZ
	object.P2.Z += offsetZ
	object.P3.Z += offsetZ

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
