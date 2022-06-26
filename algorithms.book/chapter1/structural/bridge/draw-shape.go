package main

import "fmt"

//IDrawShape interface
type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

//DrawShape struct
type DrawShape struct{}

// DrawShape struct has method draw Shape with float x and y coordinates
func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape")
}
