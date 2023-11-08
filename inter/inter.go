package inter

import (
	"image/color"
	"image"
)

func (dot Dot3D) CreateDot (x, y, z, w float64) {
	dot.X = x
	dot.Y = y
	dot.Z = z
	dot.W = w
}

type Dot3D struct {
	X float64
	Y float64
	Z float64
	W float64
}

type Polygon struct {
	P1 Dot3D
	P2 Dot3D
	P3 Dot3D

	color.Color
}

type Square struct  {
	Triagle1 Polygon
	Triagle2 Polygon

	Object int
}

type Canvas interface {
	SetPixel(x, y int, c color.Color)

	DrawLine(x1, y1, x2, y2 int, c color.Color)

	Fill(c color.Color)

	Height() int
	Width() int

	Image() image.Image
}

type MyGraphicsEngine struct {
	Cnv Canvas

	ZBuf [][] float64
	SBuf [][] float64

	Camera Camera
	LightCamera Camera
	ZBufToSBuf [4][4]float64
	ProjMatrix [4][4]float64

	Object []Square
}

type Camera struct {
	Pos Vec3

	VCamera  Vec3
	VLookDir Vec4
	FYaw     float64
	FTheta   float64
	VTarget  Vec3

	VForward Vec3

	Matrix [4][4]float64
}

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

type Vec4 struct {
	Vec3
	W float64
}