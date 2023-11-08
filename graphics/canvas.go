package graphics

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/StephaneBunel/bresenham"
)

type ImageCanvas struct {
	Img *image.NRGBA
}

func MakeImageCanvas(w, h int) ImageCanvas {
	var img ImageCanvas
	img.Img = image.NewNRGBA(image.Rect(0, 0, w, h))
	return img
}

func (cnv ImageCanvas) SetPixel(x, y int, c color.Color) {
	cnv.Img.Set(x, y, c)
}

func (cnv ImageCanvas) Fill(c color.Color) {
	draw.Draw(cnv.Img, cnv.Img.Bounds(), &image.Uniform{c}, image.Point{}, draw.Src)
}

func (cnv ImageCanvas) DrawLine(x1, y1, x2, y2 int, c color.Color) {
	bresenham.DrawLine(cnv.Img, x1, y1, x2, y2, c)
}

func (cnv ImageCanvas) Height() int {
	return cnv.Img.Rect.Dy()
}

func (cnv ImageCanvas) Width() int {
	return cnv.Img.Rect.Dx()
}

func (cnv ImageCanvas) Image() image.Image {
	return cnv.Img
}
