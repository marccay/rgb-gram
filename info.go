package main

import (
	"image"
)

func imgPxlDimensions(img image.Image) (int, int) {
	bounds := img.Bounds()
	x := bounds.Dx()
	y := bounds.Dy()
	return x, y
}
