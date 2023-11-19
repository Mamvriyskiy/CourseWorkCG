package graphics

import (
	//"fmt"
	//"fmt"
	//"fmt"
	"image/color"
	//"math"

	"github.com/Mamvriyskiy/CourseWorkCG/inter"
)

func NewMyGraphicsEngine(cnv ImageCanvas, flag bool) *inter.MyGraphicsEngine {
	var engine *inter.MyGraphicsEngine
	if !flag {
		engine = new(inter.MyGraphicsEngine)

		engine.Cnv = cnv
	}

	engine.ZBuf = CreateZBuf(engine.Cnv.Width(), engine.Cnv.Height())

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
