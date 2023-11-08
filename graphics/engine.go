package graphics

import (
	//"fmt"
	//"fmt"
	//"fmt"
	"image/color"
	//"math"

	"../inter"
)

func NewMyGraphicsEngine(cnv ImageCanvas, flag bool) *inter.MyGraphicsEngine {
	var engine *inter.MyGraphicsEngine
	if !flag {
		engine = new(inter.MyGraphicsEngine)

		engine.Cnv = cnv
	}

	engine.ZBuf = make([][]float64, engine.Cnv.Width())
	for i := range engine.ZBuf {
		engine.ZBuf[i] = make([]float64, engine.Cnv.Height())
		for j := range engine.ZBuf[i] {
			engine.ZBuf[i][j] = 100000
		}
	}

	if !flag {
		engine.SBuf = make([][]float64, engine.Cnv.Width())
		for i := range engine.SBuf {
			engine.SBuf[i] = make([]float64, engine.Cnv.Height())
			for j := range engine.SBuf[i] {
				engine.SBuf[i][j] = 100000
			}
		}
	}

	engine.Cnv.Fill(color.RGBA{3, 215, 252, 140})

	return engine
}
