package drawobj

import (
	"math"
	//"image/color"
	"../inter"
	"../mathfunc"
)

// func DrawLine3D(engine *inter.MyGraphicsEngine, p1, p2 inter.Vec4, c color.Color) {
// 	dx := p2.X - p1.X
// 	dy := p2.Y - p1.Y
// 	dz := p2.Z - p1.Z

// 	steps := int(math.Max(math.Abs(dx), math.Abs(dy)))
// 	xInc := dx / float64(steps)
// 	yInc := dy / float64(steps)
// 	zInc := dz / float64(steps)

// 	for i := 0; i < steps; i++ {
// 		x := int(p1.X + float64(i) * xInc)
// 		y := int(p1.Y + float64(i) * yInc)
// 		z := p1.Z + float64(i) * zInc

// 		if x >= 0 && x < engine.Cnv.Width() && y >= 0 && y < engine.Cnv.Height() && z < engine.ZBuf[x][y] {
// 			engine.ZBuf[x][y] = z
// 			engine.Cnv.SetPixel(x, y, color.Black)
// 		}
// 	}
// }

func DrawSquare(engine *inter.MyGraphicsEngine, slice []inter.Square) {
	for _, square := range slice {
		square.Triagle1.UpdatePolygon(engine.Cnv.Width(), engine.Cnv.Height())
		square.Triagle2.UpdatePolygon(engine.Cnv.Width(), engine.Cnv.Height())
		// fmt.Println(square.Triagle1.P3)
		// fmt.Println(square.Triagle2.P3)
		drawPolygon(engine, square.Triagle1)
		drawPolygon(engine, square.Triagle2)
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

		// DrawLine3D(engine, polygon.P2, polygon.P3, color.Black)
		// DrawLine3D(engine, polygon.P1, polygon.P3, color.Black)
		// DrawLine3D(engine, polygon.P1, polygon.P2, color.Black)
	}
}
