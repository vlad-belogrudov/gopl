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

type LissajousParams struct {
	Cycles int
	Frames int
	Xsize  int
	Ysize  int
	Res    float64
	Delay  int
}

func DefaultLissajousParams() LissajousParams {
	return LissajousParams{
		Cycles: 5,
		Res:    0.001,
		Xsize:  100,
		Ysize:  100,
		Frames: 64,
		Delay:  8,
	}
}

func Lissajous(out io.Writer, params LissajousParams) {
	rand.Seed(time.Now().UTC().UnixNano())
	colorIndex := uint8(rand.Intn(len(palette)-1) + 1)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: params.Frames}
	phase := 0.0
	for i := 0; i < params.Frames; i++ {
		rect := image.Rect(0, 0, 2*params.Xsize+1, 2*params.Ysize+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(params.Cycles*2)*math.Pi; t += params.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				params.Xsize+int(x*float64(params.Xsize)+0.5),
				params.Ysize+int(y*float64(params.Ysize)+0.5),
				colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, params.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
