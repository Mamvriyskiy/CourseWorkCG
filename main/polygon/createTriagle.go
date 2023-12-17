package polygon

import (
	"image/color"

	"../inter"
)

//вверх низ
func createTriangleBTA(x1, y1, z1, x2, y2, z2 float64, clr color.Color) inter.Polygon {
	var triangle inter.Polygon

	triangle.P1.CreateDot(x1, y1, z1, 1)
	triangle.P2.CreateDot(x2, y2, z2, 1)
	triangle.P3.CreateDot(x1, y1, z2, 1)
	
	triangle.Color = clr

	return triangle
}

func createTriangleBTB(x1, y1, z1, x2, y2, z2 float64, clr color.Color) inter.Polygon {
	var triangle inter.Polygon

	triangle.P1.CreateDot(x1, y1, z1, 1)
	triangle.P2.CreateDot(x2, y2, z2, 1)
	triangle.P3.CreateDot(x2, y1, z1, 1)

	triangle.Color = clr

	return triangle
}

//лево право
func createTriangleLRA(x1, y1, z1, x2, y2, z2 float64, clr color.Color) inter.Polygon {
	var triangle inter.Polygon

	triangle.P1.CreateDot(x1, y1, z1, 1)
	triangle.P2.CreateDot(x1, y2, z2, 1)
	triangle.P3.CreateDot(x1, y2, z1, 1)

	triangle.Color = clr

	return triangle
}

func createTriangleLRB(x1, y1, z1, x2, y2, z2 float64, clr color.Color) inter.Polygon {
	var triangle inter.Polygon

	triangle.P1.CreateDot(x1, y1, z1, 1)
	triangle.P2.CreateDot(x2, y2, z2, 1)
	triangle.P3.CreateDot(x1, y1, z2, 1)

	triangle.Color = clr

	return triangle
}

func createTriangleBBA(x1, y1, z1, x2, y2, z2 float64, clr color.Color) inter.Polygon {
	var triangle inter.Polygon

	triangle.P1.CreateDot(x1, y1, z1, 1)
	triangle.P2.CreateDot(x2, y2, z1, 1)
	triangle.P3.CreateDot(x1, y2, z1, 1)

	triangle.Color = clr

	return triangle
}

func createTriangleBBB(x1, y1, z1, x2, y2, z2 float64, clr color.Color) inter.Polygon {
	var triangle inter.Polygon

	triangle.P1.CreateDot(x1, y1, z2, 1)
	triangle.P2.CreateDot(x2, y2, z2, 1)
	triangle.P3.CreateDot(x2, y1, z2, 1)

	triangle.Color = clr

	return triangle
}
