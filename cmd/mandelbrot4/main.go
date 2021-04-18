// Mandelbrot creates PNG of Mandelbrot
// This is a supersampling version to allow smoother color changes
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
		xmin, ymin, xmax, ymax  = -2, -2, 2, 2
		width, height           = 1920, 1080
		superwidth, superheight = width * 2, height * 2
	)
	var colors [superwidth][superheight]color.RGBA
	for py := 0; py < superheight; py++ {
		y := float64(py)/superheight*(ymax-ymin) + ymin
		for px := 0; px < superwidth; px++ {
			x := float64(px)/superwidth*(xmax-xmin) + xmin
			z := complex(x, y)
			colors[px][py] = mandelbrot(z)
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			img.Set(px, py,
				getAverageColor(colors[2*px][2*py],
					colors[2*px+1][2*py],
					colors[2*px][2*py+1],
					colors[2*px+1][2*py+1]))
		}
	}
	png.Encode(os.Stdout, img)
}

func getAverageColor(colors ...color.RGBA) color.RGBA {
	var red, green, blue int
	for _, c := range colors {
		red += int(c.R)
		green += int(c.G)
		blue += int(c.B)
	}
	return color.RGBA{
		R: uint8(red / len(colors)),
		G: uint8(green / len(colors)),
		B: uint8(blue / len(colors)),
		A: 255,
	}
}

func getColor(n, nmax uint8) color.RGBA {
	x := uint8(uint32(255) * uint32(n) / uint32(nmax))
	return color.RGBA{
		R: x,
		G: 128 + x,
		B: 255 - x,
		A: 255,
	}
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 20
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return getColor(n, iterations)
		}
	}
	return color.RGBA{}
}
