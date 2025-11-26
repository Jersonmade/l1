package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(point *Point) float64 {
	return math.Sqrt(math.Pow(point.x - p.x, 2) + math.Pow(point.y - p.y, 2))
}

func main() {
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	fmt.Println(point1.Distance(point2))
}