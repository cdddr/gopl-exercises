// Madelbrot emits a PNG image of the Mandelbrot fractal
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

type Point struct {
	x, y float64
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		dx, dy = float64(1)/width, float64(1)/height
	)
	
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, supersample(x, y, dx, dy))
		}
	}
	png.Encode(os.Stdout, img)	
}

func supersample(x, y, dx, dy float64) color.Color {
	points := []Point{Point{ x - dx/2, y - dy/2 }, 
		Point{ x + dx/2, y - dy/2}, 
		Point{ x - dx/2, y + dy/2}, 
		Point{ x + dx/2, y + dy/2}}
	avgColor := uint8(0)
	for i := 0; i < 4; i++ {
		point := points[i]
		avgColor += mandelbrot(complex(point.x, point.y))
	}
	return color.Gray{avgColor/4}
}

func mandelbrot(z complex128) uint8 {
	const iterations = 200
	const contrast = 15
	
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return 255 - contrast*n
		}
	}
	return 0
}