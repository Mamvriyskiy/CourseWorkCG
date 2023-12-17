package graphics

import (
	"image/color"
	//"math"

	"../inter"
)

func NewMyGraphicsEngine(cnv ImageCanvas) *inter.MyGraphicsEngine {
	var engine *inter.MyGraphicsEngine
	engine = new(inter.MyGraphicsEngine)

	engine.Cnv = cnv

	engine.Cnv.Fill(color.RGBA{3, 215, 252, 140})
	engine.ZBuf = CreateZBuf(engine.Cnv.Width(), engine.Cnv.Height())

	engine.SBuf =  CreateZBuf(engine.Cnv.Width(), engine.Cnv.Height())

	return engine
}

func CreateZBuf(width, height int) [][]float64 {
	ZBuf := make([][]float64, width)
	for i := range ZBuf {
		ZBuf[i] = make([]float64, height)
		for j := range ZBuf[i] {
			ZBuf[i][j] = 100000
		}
	}

	return ZBuf
}
