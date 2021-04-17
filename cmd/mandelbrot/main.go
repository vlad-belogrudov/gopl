// Mandelbrot creates PNG of Mandelbrot
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
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1920, 1080
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func getColor(n, nmax uint8) color.Color {
	x := uint8(uint32(255) * uint32(n) / uint32(nmax))
	return color.RGBA{
		R: x,
		G: 128 + x,
		B: 255 - x,
		A: 255,
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 20
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return getColor(n, iterations)
		}
	}
	return color.Black
}
