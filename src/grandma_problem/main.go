package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) Distance(point *Point) float64 {
	return math.Sqrt(math.Pow(p.x-point.x, 2) + math.Pow(p.y-point.y, 2))
}

func main() {
	var short [2]*Point
	var newPoint *Point
	points := make([]*Point, 0)
	distance := math.MaxFloat64
	line := ""
	x := 0.0
	y := 0.0

	fmt.Scanf("%f", &x)
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		line = scan.Text()
		nums := strings.Split(line[1:len(line)-1], ", ")
		x, _ = strconv.ParseFloat(nums[0], 64)
		y, _ = strconv.ParseFloat(nums[1], 64)
		newPoint = &Point{x, y}
		for _, point := range points {
			if newDist := newPoint.Distance(point); newDist < distance {
				short[1], short[0] = point, newPoint
				distance = newDist
			}
		}
		points = append(points, &Point{x, y})
	}

	fmt.Printf("(%v, %v) (%v, %v)\n", short[0].x, short[0].y, short[1].x, short[1].y)
}
