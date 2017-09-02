// Madelbrot emits a PNG image of the Mandelbrot fractal
package main

import (
//	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

type Point struct {
	x, y float64
}

func abs(i int16) int16 {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -0.01, -0.01, +0.01, +0.01
		width, height          = 1024, 1024
		dx, dy                 = float64(1) / width, float64(1) / height
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func newton(z complex128) color.Color {
	const tolerance = 1.0e-12
	const contrast = 15
	const iterations = 768

	z0 := z
	n := 0
	if f(z) != complex(0, 0) && fp(z) != complex(0, 0) {
		for {
			z1 := z0
			z0 = z0 - f(z0)/fp(z0)
			n++
			if cmplx.Abs(z0 - z1) < tolerance || n >= iterations {
				r := 255 - contrast*(abs(int16(n)-512) % 256)
				g := 255 - contrast*(abs(int16(n)-256) % 256)
				b := 255 - contrast*(n % 256)
				return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
			}
		}
	}

	return color.RGBA{0, 0, 0, 255}
}

func f(z complex128) complex128 {
	return cmplx.Pow(z, 4) - 1
}

func fp(z complex128) complex128 {
	return 4*cmplx.Pow(z, 3)
}
