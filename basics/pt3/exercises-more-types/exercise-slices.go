package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// Pic returns a 2D slice of uint8 values
func Pic(dx, dy int) [][]uint8 {
	// Allocate the outer slice
	pic := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		// Allocate each inner slice
		pic[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// You can change this formula to (x + y) / 2 or (x * y)
			pic[y][x] = uint8(x ^ y)
		}
	}

	return pic
}

func main() {
	dx, dy := 256, 256
	imgData := Pic(dx, dy)

	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := imgData[y][x]
			m.SetNRGBA(x, y, color.NRGBA{R: v, G: v, B: 255, A: 255})
		}
	}

	f, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, m)
}
