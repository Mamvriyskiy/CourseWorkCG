package drawobj

import (
	//"fmt"
	//"fmt"
	"image/color"
	"math"

	"../inter"
	"../mathfunc"
)

func DrawSquare(engine *inter.MyGraphicsEngine, slice []inter.Square) {
	for _, square := range slice {
		square.Triagle1.UpdatePolygon(engine.Cnv.Width(), engine.Cnv.Height())
		square.Triagle2.UpdatePolygon(engine.Cnv.Width(), engine.Cnv.Height())
		if square.Object == 100 {
			drawBoldEdge(engine, square.Triagle2.P2, square.Triagle2.P3, square.Triagle1.Color, 1.5)
			drawBoldEdge(engine, square.Triagle1.P1, square.Triagle1.P3, square.Triagle1.Color, 1.5)
		} else if square.Object == 101 {
			drawBoldEdge(engine, square.Triagle2.P1, square.Triagle2.P3, square.Triagle1.Color, 1.5)
			drawBoldEdge(engine, square.Triagle1.P2, square.Triagle1.P3, square.Triagle1.Color, 1.5)
		} else {
			drawPolygon(engine, square.Triagle1)
			drawPolygon(engine, square.Triagle2)
		}
	}
}

func drawPolygon(engine *inter.MyGraphicsEngine, polygon inter.Polygon) {
	p0, p1, p2 := polygon.P1, polygon.P2, polygon.P3

	if p0.Y > p1.Y {
		p0, p1 = p1, p0
	}

	if p0.Y > p2.Y {
		p0, p2 = p2, p0
	}

	if p1.Y > p2.Y {
		p1, p2 = p2, p1
	}

	dyTotal := p2.Y - p0.Y

	for y := p0.Y; y <= p1.Y; y++ {
		dySegment := p1.Y - p0.Y
		alpha := float64((y - p0.Y) / dyTotal)
		beta := float64((y - p0.Y) / dySegment)

		var a, b inter.Vec4

		a = mathfunc.Vec4Diff(p2, p0)
		a.Mul(alpha)
		a.Add(p0)

		b = mathfunc.Vec4Diff(p1, p0)
		b.Mul(beta)
		b.Add(p0)

		if a.X > b.X {
			a, b = b, a
		}

		for x := a.X; x <= b.X + 4e-1; x++ {
			var (
				phi float64
				p inter.Vec4
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}

			p.Z = a.Z + (b.Z - a.Z) * phi
			p.W = a.W + (b.W - a.W) * phi

			p.X = x
			p.Y = y

			px := int(math.Round(p.X))
			py := int(math.Round(p.Y))
			if px >= 0 && py >= 0 && px < engine.Cnv.Width() && py < engine.Cnv.Height() {
				if p.Z < engine.ZBuf[px][py] {
					engine.ZBuf[px][py] = p.Z
					engine.Cnv.SetPixel(px, py, polygon.Color)
				}
			}
		}
	}

	for y := p1.Y; y <= p2.Y; y++ {
		dySegment := p2.Y - p1.Y
		alpha := float64((y - p0.Y) / dyTotal)
		beta := float64((y - p1.Y) / dySegment)

		var a, b inter.Vec4

		a = mathfunc.Vec4Diff(p2, p0)
		a.Mul(alpha)
		a.Add(p0)

		b = mathfunc.Vec4Diff(p2, p1)
		b.Mul(beta)
		b.Add(p1)

		if a.X > b.X {
			a, b = b, a
		}

		for x := a.X; x <= b.X + 4e-1; x++ {
			var (
				phi float64
				p inter.Vec4
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}

			p.Z = a.Z + (b.Z - a.Z) * phi
			p.W = a.W + (b.W - a.W) * phi
			p.X = x
			p.Y = y

			px := int(math.Round(p.X))
			py := int(math.Round(p.Y))
			if px >= 0 && py >= 0 && px < engine.Cnv.Width() && py < engine.Cnv.Height() {
				if p.Z < engine.ZBuf[px][py] {
					engine.ZBuf[px][py] = p.Z
					engine.Cnv.SetPixel(px, py, polygon.Color)
				}
			}
		}

		drawBoldEdge(engine, polygon.P2, polygon.P3, color.Black, 1)
		drawBoldEdge(engine, polygon.P1, polygon.P3, color.Black, 1)
	}
}

func drawBoldEdge(engine *inter.MyGraphicsEngine, p0, p1 inter.Vec4, borderColor color.Color, thickness float64) {
	var (
		deltaX  = p1.X - p0.X
		deltaY  = p1.Y - p0.Y
		deltaZ  = p1.Z - p0.Z
		deltaW  = p1.W - p0.W
		lengthX = math.Abs(deltaX)
		lengthY = math.Abs(deltaY)
		steps   int
	)

	if lengthX >= lengthY {
		steps = int(lengthX)
	} else {
		steps = int(lengthY)
	}

	var (
		xIncrement = deltaX / float64(steps)
		yIncrement = deltaY / float64(steps)
		zIncrement = deltaZ / float64(steps)
		wIncrement = deltaW / float64(steps)
		p          = p0
	)

	for i := 0; i < steps; i++ {
		px := int(math.Round(p.X))
		py := int(math.Round(p.Y))

		for dx := -int(thickness); dx <= int(thickness); dx++ {
			for dy := -int(thickness); dy <= int(thickness); dy++ {
				if px+dx >= 0 && py+dy >= 0 && px+dx < engine.Cnv.Width() && py+dy < engine.Cnv.Height() {
					if p.Z <= engine.ZBuf[px+dx][py+dy] {
						engine.ZBuf[px+dx][py+dy] = p.Z
						engine.Cnv.SetPixel(px+dx, py+dy, borderColor)
					}
				}
			}
		}

		p.X += xIncrement
		p.Y += yIncrement
		p.Z += zIncrement
		p.W += wIncrement
	}

	// Draw the last pixel with the fill color
	px := int(math.Round(p1.X))
	py := int(math.Round(p1.Y))
	for dx := -int(thickness); dx <= int(thickness); dx++ {
		for dy := -int(thickness); dy <= int(thickness); dy++ {
			if px+dx >= 0 && py+dy >= 0 && px+dx < engine.Cnv.Width() && py+dy < engine.Cnv.Height() {
				if p.Z <= engine.ZBuf[px+dx][py+dy] {
					engine.ZBuf[px+dx][py+dy] = p.Z
					engine.Cnv.SetPixel(px+dx, py+dy, borderColor)
				}
			}
		}
	}
}
