package main

import (
	"image/color"
)

var (
	colors = []color.RGBA{
		color.RGBA{240, 255, 240, 0},
		color.RGBA{68, 36, 52, 0},
		color.RGBA{48, 52, 109, 0},
		color.RGBA{78, 74, 78, 0},
		color.RGBA{133, 76, 48, 0},
		color.RGBA{52, 101, 36, 0},
		color.RGBA{208, 70, 72, 0},
		color.RGBA{117, 113, 97, 0},
		color.RGBA{89, 125, 206, 0},
		color.RGBA{210, 125, 44, 0},
		color.RGBA{133, 149, 161, 0},
		color.RGBA{109, 170, 44, 0},
		color.RGBA{210, 170, 153, 0},
		color.RGBA{109, 194, 202, 0},
		color.RGBA{218, 212, 94, 0},
		color.RGBA{222, 238, 214, 0},
	}
)

func GetColor(byteValue byte) color.RGBA {
	return colors[HexToInt(byteValue)]
}

func HexAdd(byteOne, byteTwo byte) int {
	return HexToInt(byteOne) + HexToInt(byteTwo)
}

func HexToInt(byteValue byte) int {
	byteFor0 := byte(48)
	byteFora := byte(97)

	for i := 0; i < 10; i++ {
		if newVal := byteFor0 + byte(i); newVal == byteValue {
			return i
		}
	}

	for i := 0; i < 6; i++ {
		if newVal := byteFora + byte(i); newVal == byteValue {
			return 10 + i
		}
	}

	return -1
}
