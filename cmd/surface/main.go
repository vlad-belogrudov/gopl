// Surface computes SVG for 3D function
package main

import (
	"fmt"
	"math"

	"github.com/vlad-belogrudov/gopl/pkg/color"
)

const (
	width, height = 1200, 640
	cells         = 200
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var zmin, zmax = math.Inf(1), math.Inf(-1)
var topColor = color.Color{Red: 255, Green: 0, Blue: 0}
var bottomColor = color.Color{Red: 0, Green: 100, Blue: 200}

func init() {
	for x := 0; x < xyrange; x++ {
		for y := 0; y < xyrange; y++ {
			z, ok := f(float64(x), float64(y))
			if !ok {
				continue
			}
			zmin = math.Min(zmin, z)
			zmax = math.Max(zmax, z)
		}
	}
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ac, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, bc, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, cc, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, dc, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			c := color.MixColors(color.MixColors(color.MixColors(ac, bc), cc), dc)
			fmt.Printf("<polygon fill='#%02X%02X%02X' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				c.Red, c.Green, c.Blue,
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, color.Color, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z, ok := f(x, y)
	if !ok {
		return 0, 0, color.Color{}, false
	}
	colorScale := (z - zmin) / float64(zmax-zmin)
	colorHigh := topColor.Scale(colorScale)
	colorLow := bottomColor.Scale(1.0 / colorScale)
	color := color.MixColors(colorHigh, colorLow)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, color, true
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y)
	if r == 0 {
		return 0, false
	}
	return math.Sin(r) / r, true
}
