package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixel
	cells         = 100                 // number of grids cells
	xyrange       = 30.0                // axis ranges (-xyranges..+xyranges)
	xyscale       = width / 2 / xyrange // pixel per x or y unit
	zscale        = height * 0.4        // pixel per z unit
	angle         = math.Pi / 6         // angle of x,y axes (30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30), cos(30)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			rgb := rgb(i, j)

			if !math.IsNaN(ax+ay+bx+by+cx+cy+dx+dy) && !math.IsInf(ax+ay+bx+by+cx+cy+dx+dy, 0) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: rgba(%v,%v,%v,%v)'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, rgb.R, rgb.G, rgb.B, rgb.A)
			}

		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// find point (x, y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute surface height z.
	z := f(x, y)

	// project (x,y,z) isometrically onto 2-D svg canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func rgb(i, j int) color.RGBA {
	// find point (x, y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute surface height z.
	z := f(x, y)
	if z > 0 {
		return color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	} else {
		return color.RGBA{0x00, 0x00, 0xFF, 0xFF}
	}
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
