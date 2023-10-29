package graphics

import (
	"image"
	// "image/color"
	// "image/draw"

	// "github.com/StephaneBunel/bresenham"
)

type ImageCanvas struct {
	Img *image.NRGBA
}

func MakeImageCanvas(w, h int) ImageCanvas {
	var img ImageCanvas
	img.Img = image.NewNRGBA(image.Rect(0, 0, w, h))
	return img
}

func (cnv ImageCanvas) Image() image.Image {
	return cnv.Img
}
