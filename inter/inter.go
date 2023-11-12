package inter

import (
	"image"
	"image/color"
)

func (dot *Vec4) CreateDot (x, y, z, w float64) {
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

func (p *Polygon) UpdatePolygon(width, height int) {
	p.P1.X = (p.P1.X / p.P1.W + 1) * 0.5 * float64(width)
	p.P2.X = (p.P2.X / p.P2.W + 1) * 0.5 * float64(width)
	p.P3.X = (p.P3.X / p.P3.W + 1) * 0.5 * float64(width)

	p.P1.Y = (1 - p.P1.Y / p.P1.W) * 0.5 * float64(height)
	p.P2.Y = (1 - p.P2.Y / p.P2.W) * 0.5 * float64(height)
	p.P3.Y = (1 - p.P3.Y / p.P3.W) * 0.5 * float64(height)
}

type Polygon struct {
	P1 Vec4
	P2 Vec4
	P3 Vec4

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

func (a *Vec3) Sub(b Vec3) {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z
}

func (a *Vec3) Add(b Vec3) {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
}

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (a *Vec4) Sub(b Vec4) {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z
	a.W -= b.W
}

func (a *Vec4) Mul(b float64) {
	a.X *= b
	a.Y *= b
	a.Z *= b
	a.W *= b
}

func (a *Vec4) Add(b Vec4) {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
	a.W += b.W
}

func (a *Vec4) CreateVec4(x, y, z, w float64) {
	a.X = x
	a.Y = y
	a.Z = z
	a.W = w
}

type Vec4 struct {
	Vec3
	W float64
}