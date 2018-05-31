package main

import (
	"fmt"
)

func maxnormalize(red []uint32, green []uint32, blue []uint32) uint32 {
	redmax := returnmax(red)
	greenmax := returnmax(green)
	bluemax := returnmax(blue)

	max := redmax
	if greenmax > max {
		max = greenmax
	}
	if bluemax > max {
		max = bluemax
	}

	return max
}

func returnmax(arr []uint32) uint32 {
	var max uint32
	for _, x := range arr {
		if x > max {
			max = x
		}
	}
	return max
}

func printHistogram(arr []uint32, max uint32) {
	fmt.Println(" ------------------------------------------------------------------")
	var y uint32 = 20
	for y = 20; y > 0; y-- {
		var printStr string
		fmt.Printf(" |")
		for i := 0; i < len(arr); i++ {
			if arr[i]*100 >= y*5*max {
				printStr += " "
			} else {
				printStr += "0"
			}
			if i == len(arr)-1 {
				printStr += "|"
			}
		}
		fmt.Println(printStr)
	}
	fmt.Println(" ------------------------------------------------------------------")
}

func histogram(pixels []pixel) ([]uint32, []uint32, []uint32) {
	red := make([]uint32, 64)
	green := make([]uint32, 64)
	blue := make([]uint32, 64)
	for _, pxl := range pixels {
		r := pxl.r / 1024
		red[r]++
		g := pxl.g / 1024
		green[g]++
		b := pxl.b / 1024
		blue[b]++
	}
	return red, green, blue
}
