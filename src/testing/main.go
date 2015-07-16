package main

import (
	"fmt"
	"image/color"
	"os"
	"strconv"
	"time"
)

func main() {
	aTimer := time.NewTimer(0)

	red := color.RGBA{255, 0, 0, 0}
	fmt.Println(red.RGBA())

	aString := "hello"
	fmt.Println(aString[2:4])

	byteValue := byte(30)
	newByteValue := byteValue + 1
	fmt.Println(byteValue, newByteValue)
	fmt.Println(newByteValue == byte(31))

	fmt.Println(3 / 2)

	one := "3"
	two := "5"

	newOne, _ := strconv.ParseInt(one, 10, 8)
	newTwo, _ := strconv.ParseInt(two, 10, 8)

	eight := newOne + newTwo

	fmt.Println(fmt.Sprintf("%d", eight))

	anotherString := "hello person how are you"

	fmt.Printf("%q", anotherString[4])

	envValues := os.Environ()

	for _, value := range envValues {
		fmt.Printf("\n%s", value)
	}

	fmt.Println("\n\n\n")

	fmt.Println(os.Hostname())

	// os.Mkdir("yo", 0777)
	// the above creates a new directory in the working file

	rightNow := time.Now()

	isItAfter := rightNow.After(time.Now())
	isItBefore := rightNow.Before(time.Now())

	fmt.Println(isItAfter)
	fmt.Println(isItBefore)

	fmt.Println(rightNow.Clock())

	time.Sleep(time.Second * 5)

	aTimer.Stop()

	fmt.Println(aTimer.C)
}
