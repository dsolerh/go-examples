package main

import "fmt"

//IContour interace
type IContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

//DrawContour struct
type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

//DrawContour method drawContour given the coordinates
func (contour *DrawContour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}

//DrawContour method resizeByFactor given factor
func (contour *DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
}
