package main

import (
	"fmt"
	"math"
)

func folds(n int) []bool {
	size := int(math.Pow(2.0, float64(n)) - 1)
	foldsArray := make([]bool, size)

	for i := 0; i <= size/2; i++ {
		if i == 0 {
			foldsArray[i] = true
			continue
		}
		foldsArray[i] = true
		valAdded := 0
		for j := i - 1; j >= 0; j-- {
			foldsArray[i+1+valAdded] = !foldsArray[j]
			valAdded++
		}
		i += valAdded
	}
	return foldsArray
}

func recurseFolds(n int) []bool {
	if n == 1 {
		return []bool{true}
	}

	oldArray := recurseFolds(n - 1)
	newArray := []bool{}
	newArray = append(newArray, oldArray...)
	newArray = append(newArray, true)

	for i := 0; i < len(oldArray)/2; i++ {
		oldArray[i], oldArray[len(oldArray)-1-i] = !oldArray[len(oldArray)-1-i], !oldArray[i]
	}

	if len(oldArray)%2 != 0 {
		oldArray[len(oldArray)/2] = !oldArray[len(oldArray)/2]
	}

	newArray = append(newArray, oldArray...)
	return newArray
}

func main() {

	turns := recurseFolds(4)
	turns2 := folds(4)

	for _, turn := range turns {
		if turn {
			fmt.Printf("left ")
		} else {
			fmt.Printf("right ")
		}
	}

	fmt.Printf("\n")

	for _, turn := range turns2 {
		if turn {
			fmt.Printf("left ")
		} else {
			fmt.Printf("right ")
		}
	}
}
