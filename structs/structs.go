package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	a := Point{1, -3}
	a.x = 2

	b := Point{} // implicit x = 0 and y = 0

	c := Point{x: 2} // implicit y = 0

	d := &Point{1, 2} // pointer to a

	fmt.Println(a, b, c, d)

}
