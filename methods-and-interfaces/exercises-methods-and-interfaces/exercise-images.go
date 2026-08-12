package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// Image defines a custom image implementation of image.Image
type Image struct {
	Width  int
	Height int
}

// ColorModel returns color.RGBAModel
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns an image.Rectangle such as image.Rect(0, 0, w, h)
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

// At computes the pixel color at (x, y) matching color.RGBA{v, v, 255, 255}
func (img Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{R: v, G: v, B: 255, A: 255}
}

func main() {
	m := Image{Width: 256, Height: 256}

	f, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, m)
}
