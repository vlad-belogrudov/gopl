// Mandelbrot creates PNG of Mandelbrot
package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

func Mandelbrot(w io.Writer, xmin, xmax, ymin, ymax float64, width, height int) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, iterate(z))
		}
	}
	png.Encode(w, img)
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

func iterate(z complex128) color.Color {
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
