package main

import "fmt"

type Point struct {
	x, y int
}

type Rectangle struct {
	x, y, width, height int
}

func (p *Point) GetX() string {
	return fmt.Sprintf("%d", p.x)
}

func (p *Point) GetY() string {
	return fmt.Sprintf("%d", p.y)
}

func (r *Rectangle) GetX() string {
	return fmt.Sprintf("%d", r.x)
}

func (r *Rectangle) GetY() string {
	return fmt.Sprintf("%d", r.y)
}

func PrintPoints(d DotType) {
	fmt.Printf("x: %s y: %s\n", d.GetX(), d.GetY())
}

type DotType interface {
	GetX() string
	GetY() string
}

func main() {
	point_1 := &Point{3, 4}
	rectangle_1 := &Rectangle{0, 0, 10, 10}

	fmt.Println(rectangle_1.GetY())

	PrintPoints(point_1)
	PrintPoints(rectangle_1)
}
