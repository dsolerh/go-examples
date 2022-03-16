/*
Implement a painting program
It should support

- Circle whit location (x, y), color and radius
- Rectangle with location  (x, y), width, height and color

Each type should implement a `Draw(d Device)` method.

Implement an `ImageCanvas` struct which holds a slice of drawable items and has a
`Draw(w io.Writer)` that writes a PNG to w (using `image/png`).
*/
package main

import (
	"image/color"
	"io"
	"log"
	"os"
)

var (
	Red   = color.RGBA{0xFF, 0, 0, 0xFF}
	Green = color.RGBA{0, 0xFF, 0, 0xFF}
	Blue  = color.RGBA{0, 0, 0xFF, 0xFF}
)

type Circle struct {
}

func NewCircle(x, y, r int, c color.Color) *Circle {
	return nil
}

type Rectangle struct {
}

func NewRectangle(x, y, h, w int, c color.Color) *Rectangle {
	return nil
}

type Device interface {
	Set(int, int, color.Color)
}

type ImageCanvas struct {
}

func NewImageCanvas(height, width int) (*ImageCanvas, error) {
	return nil, nil
}

type Drawer interface {
}

func (ic *ImageCanvas) Add(d Drawer) {
}

func (ic *ImageCanvas) Draw(w io.Writer) error {
	return nil
}

func main() {
	ic, err := NewImageCanvas(200, 300)
	if err != nil {
		log.Fatal(err)
	}

	ic.Add(NewCircle(100, 100, 80, Green))
	ic.Add(NewCircle(60, 60, 10, Blue))
	ic.Add(NewCircle(140, 60, 10, Blue))
	ic.Add(NewRectangle(100, 130, 80, 10, Blue))

	f, err := os.Create("face.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err = ic.Draw(f); err != nil {
		log.Fatal(err)
	}

}
