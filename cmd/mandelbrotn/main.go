// Mandelbrot creates PNG of Mandelbrot N**4-1
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
	"runtime"
	"sync"
)

type job struct {
	xstart, xend, ystart, yend int
}

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 1920, 1080
	chunks                 = 16
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	jobs := make(chan job, 256)
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	for i := 0; i < numWorkers; i++ {
		go worker(&wg, jobs, img)
		wg.Add(1)
	}

	increment := height / chunks
	for py := 0; py < height; py += increment {
		jobs <- job{
			xstart: 0,
			xend:   width,
			ystart: py,
			yend:   py + increment,
		}
	}
	// last chunk if needed
	if height%chunks != 0 {
		jobs <- job{
			xstart: 0,
			xend:   width,
			ystart: height - height%chunks,
			yend:   height,
		}
	}

	close(jobs)
	wg.Wait()
	if err := png.Encode(os.Stdout, img); err != nil {
		fmt.Fprintf(os.Stderr, "Error while encoding image: %v\n", err)
	}
}

func worker(wg *sync.WaitGroup, jobs chan job, img *image.RGBA) {
	defer wg.Done()
	for j := range jobs {
		for py := j.ystart; py < j.yend; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := j.xstart; px < j.xend; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z, n, ok := neuton(complex(x, y))
				if !ok {
					img.Set(px, py, color.Black)
					continue
				}
				c, ok := colors(z)
				if !ok {
					img.Set(px, py, color.White)
					continue
				}
				img.Set(px, py, reduceColor(c, n))
			}
		}
	}
}

func reduceColor(c color.RGBA, n uint8) color.RGBA {
	return color.RGBA{
		R: c.R / n * 2,
		G: c.G / n * 2,
		B: c.B / n * 2,
		A: 255,
	}
}

func neuton(z complex128) (complex128, uint8, bool) {
	const iterations = 255
	for n := uint8(1); n < iterations; n++ {
		d := delta(z)
		if d == 0 {
			return 0, 0, false
		}
		z -= f(z) / d
		if cmplx.Abs(f(z)) < 0.0001 {
			return z, n, true
		}
	}
	return 0, 0, false
}

// function specifics

func f(z complex128) complex128 {
	return cmplx.Pow(z, 4) - 1
}

func delta(z complex128) complex128 {
	return 4 * cmplx.Pow(z, 3)
}

var colormap = map[complex128]color.RGBA{
	1:   {R: 255, G: 0, B: 0, A: 255},
	-1:  {R: 0, G: 255, B: 0, A: 255},
	-1i: {R: 0, G: 0, B: 255, A: 255},
	1i:  {R: 255, G: 255, B: 0, A: 255},
}

func colors(z complex128) (c color.RGBA, ok bool) {
	c, ok = colormap[complex(math.RoundToEven(real(z)), math.RoundToEven(imag(z)))]
	return
}
