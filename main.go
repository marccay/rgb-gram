package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

type pixel struct {
	r, g, b, a uint32
}

type modifiedrgba struct {
	r, g, b, a uint8
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " [filename]")
		os.Exit(1)
	}
	var option string
	path := os.Args[1]
	if len(os.Args) == 2 {
		option = "all"
	} else {
		option = os.Args[2]
	}

	img := loadImage(path)
	pixels := arrayPixels(img)

	printInterface(option, pixels)
}

// opens file as image.Image through jpeg.Decode()
func loadImage(file string) image.Image {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func printInterface(option string, pixels []pixel) {
	red, green, blue := histogram(pixels)
	nmax := maxnormalize(red, green, blue)
	switch option {
	case "r", "red":
		printHistogram(red, nmax)
	case "g", "green":
		printHistogram(green, nmax)
	case "b", "blue":
		printHistogram(blue, nmax)
	case "all", "", "rgb":
		printHistogram(red, nmax)
		fmt.Println()
		printHistogram(green, nmax)
		fmt.Println()
		printHistogram(blue, nmax)
	}
}

// get RGBA values at every pixel
func arrayPixels(img image.Image) []pixel {
	bounds := img.Bounds()

	// array of pixels, with #of pixels)
	pixels := make([]pixel, bounds.Dx()*bounds.Dy())
	var i int

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[i].r = r
			pixels[i].g = g
			pixels[i].b = b
			pixels[i].a = a
			i++
		}
	}

	return pixels
}
