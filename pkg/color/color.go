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

func MixColorByte(b1, b2 ColorByte) ColorByte {
	b := b1
	if b2 > b1 {
		b = b2
	}
	return b
}

func (c Color) Scale(x float64) Color {
	return Color{
		Red:   c.Red.Scale(x),
		Green: c.Green.Scale(x),
		Blue:  c.Blue.Scale(x),
	}
}

func MixColors(c1, c2 Color) Color {
	return Color{
		Red:   MixColorByte(c1.Red, c2.Red),
		Green: MixColorByte(c1.Green, c2.Green),
		Blue:  MixColorByte(c1.Blue, c2.Blue),
	}
}

func (c Color) String() string {
	return fmt.Sprintf("#%02X%02X%02X", c.Red, c.Green, c.Blue)
}
