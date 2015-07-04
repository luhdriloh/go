package main

import (
	"fmt"
	"image/color"
)

func main() {
	red := color.RGBA{255, 0, 0, 0}
	fmt.Println(red.RGBA())

	aString := "hello"
	fmt.Println(aString[2:4])

	byteValue := byte(30)
	newByteValue := byteValue + 1
	fmt.Println(byteValue, newByteValue)
	fmt.Println(newByteValue == byte(31))

	fmt.Println(3 / 2)
}
