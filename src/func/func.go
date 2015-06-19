package main

import "fmt"

func main() {
	x := 34
	y := 66

	z := add_int(x, y)

	fmt.Println("z =", z)

	a := 6.21
	b := 9.59

	c := add_double(a, b)

	fmt.Println("c =", c)

}

func add_int(value_one int, value_two int) int {
	return value_one + value_two
}

func add_double(value_one float64, value_two float64) float64 {
	return value_one + value_two
}
