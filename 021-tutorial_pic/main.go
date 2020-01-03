package main

import (
	"image"
	"image/color"

	"code.google.com/go-tour/pic"
)

//Image function and variable
type Image struct {
	Min, Max Point
}

//Point is pixel Absolute position
type Point struct {
	X, Y int
}

//ColorModel color model
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

//Bounds color
func (i Image) Bounds() image.Rectangle {
	// return image.Rect(0, 0, i.Size().Y, i.Size().X)
	return image.Rect(0, 0, 255, 255)
}

//At piexe color
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(255), uint8(255)}
}

// Size pic size
// func (i Image) Size() image.Point {
// 	return image.Point{i.Max.X - i.Min.X, i.Max.Y - i.Min.Y}
// }

func main() {
	m := Image{}
	pic.ShowImage(m)
}
