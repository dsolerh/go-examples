package main

func main() {
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}
	var contour IContour = &DrawContour{x, y, DrawShape{}, 2}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}
