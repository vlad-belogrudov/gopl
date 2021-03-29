// Lissajous generates animated GIF from random Lissajous figures
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

var palette = []color.Color{color.Black,
	color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	},
	color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	},
	color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	},
	color.RGBA{
		R: 255,
		G: 255,
		B: 0,
		A: 255,
	},
}

func Lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	rand.Seed(time.Now().UTC().UnixNano())
	colorIndex := uint8(rand.Intn(len(palette)-1) + 1)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
