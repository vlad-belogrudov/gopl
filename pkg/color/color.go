// Package color represents simple 3 byte color type
package color

import "fmt"

type ColorByte byte

type Color struct {
	Red   ColorByte
	Green ColorByte
	Blue  ColorByte
}

func (b ColorByte) Scale(x float64) ColorByte {
	y := float64(b) * x
	if y > 255 {
		return ColorByte(255)
	} else if y < 0 {
		return ColorByte(0)
	}
	return ColorByte(y)
}

func MixColorBytes(bytes ...ColorByte) ColorByte {
	var b ColorByte
	for _, x := range bytes {
		if x > b {
			b = x
		}
	}
	return b
}

func MixColors(colors ...Color) Color {
	var reds, greens, blues []ColorByte
	for _, x := range colors {
		reds = append(reds, x.Red)
		greens = append(greens, x.Green)
		blues = append(blues, x.Blue)
	}
	return Color{
		Red:   MixColorBytes(reds...),
		Green: MixColorBytes(greens...),
		Blue:  MixColorBytes(blues...),
	}
}

func (c Color) Scale(x float64) Color {
	return Color{
		Red:   c.Red.Scale(x),
		Green: c.Green.Scale(x),
		Blue:  c.Blue.Scale(x),
	}
}

func (c Color) String() string {
	return fmt.Sprintf("#%02X%02X%02X", c.Red, c.Green, c.Blue)
}
