// Surface computes SVG for 3D function
package surface

import (
	"fmt"
	"io"
	"math"

	"github.com/vlad-belogrudov/gopl/pkg/color"
)

type SurfaceParams struct {
	Function    func(x, y float64) (float64, bool)
	XYRange     float64
	Cells       int
	Width       int
	Height      int
	TopColor    color.Color
	BottomColor color.Color
}

func DefaultSurfaceParams() SurfaceParams {
	return SurfaceParams{
		Function: func(x float64, y float64) (float64, bool) {
			r := math.Hypot(x, y)
			if r == 0 {
				return 0, false
			}
			return math.Sin(r) / r, true
		},
		XYRange:     30.0,
		Cells:       200,
		Width:       1200,
		Height:      640,
		TopColor:    color.Color{Red: 255, Green: 0, Blue: 0},
		BottomColor: color.Color{Red: 0, Green: 0, Blue: 255},
	}
}

func Surface(out io.Writer, params SurfaceParams) {
	zmin, zmax := CalculateMinMax(params.Function, params.XYRange)
	corner := CornerFunction{
		SurfaceParams: params,
		XYScale:       float64(params.Width) / 2 / params.XYRange,
		ZScale:        float64(params.Height) * 0.4,
		ZMin:          zmin,
		ZMax:          zmax,
	}

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", params.Width, params.Height)

	for i := 0; i < params.Cells; i++ {
		for j := 0; j < params.Cells; j++ {
			ax, ay, ac, ok := corner.Calculate(i+1, j)
			if !ok {
				continue
			}
			bx, by, bc, ok := corner.Calculate(i, j)
			if !ok {
				continue
			}
			cx, cy, cc, ok := corner.Calculate(i, j+1)
			if !ok {
				continue
			}
			dx, dy, dc, ok := corner.Calculate(i+1, j+1)
			if !ok {
				continue
			}
			c := color.MixColors(ac, bc, cc, dc)
			fmt.Fprintf(out, "<polygon fill='#%02X%02X%02X' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				c.Red, c.Green, c.Blue,
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Fprintln(out, "</svg>")
}

type CornerFunction struct {
	SurfaceParams
	XYScale float64
	ZScale  float64
	ZMin    float64
	ZMax    float64
}

const angle = math.Pi / 6

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func (c *CornerFunction) Calculate(i, j int) (float64, float64, color.Color, bool) {
	x := c.XYRange * (float64(i)/float64(c.Cells) - 0.5)
	y := c.XYRange * (float64(j)/float64(c.Cells) - 0.5)
	z, ok := c.Function(x, y)
	if !ok {
		return 0, 0, color.Color{}, false
	}
	colorScale := (z - c.ZMin) / (c.ZMax - c.ZMin)
	colorHigh := c.TopColor.Scale(colorScale)
	colorLow := c.BottomColor.Scale(1.0 / colorScale)
	color := color.MixColors(colorHigh, colorLow)
	sx := float64(c.Width)/2 + (x-y)*cos30*c.XYScale
	sy := float64(c.Height)/2 + (x+y)*sin30*c.XYScale - z*c.ZScale
	return sx, sy, color, true
}

func CalculateMinMax(f func(x, y float64) (float64, bool), xyrange float64) (float64, float64) {
	zmin, zmax := math.Inf(1), math.Inf(-1)
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			z, ok := f(float64(x)/100*xyrange, float64(y)/100*xyrange)
			if !ok {
				continue
			}
			zmin = math.Min(zmin, z)
			zmax = math.Max(zmax, z)
		}
	}
	return zmin, zmax
}
