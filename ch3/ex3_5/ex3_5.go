// Madelbrot emits a PNG image of the Mandelbrot fractal
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, madelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)	
}

func abs(i int16) int16 {
	if i < 0 {
		return -i
	}
	return i
}

func madelbrot(z complex128) color.Color {
	const iterations = 768
	const contrast = 15
	
	var v complex128
	for n := uint16(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := 255 - contrast*(abs(int16(n)-512) % 256)
			b := 255 - contrast*(abs(int16(n)-256) % 256)
			g := 255 - contrast*(n % 256)
			return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		}
	}
	return color.Black
}