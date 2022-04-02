package main

import (
	"fmt"
	"math"
)

type Square struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

type Shapes [4]Shape

func (s Square) Area() float64 {
	return s.width * s.height
}

func (c Circle) Area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (s Shapes) LargestArea() {

	s[0] = Square{width: 7, height: 4}
	s[1] = Square{width: 8, height: 3}
	s[2] = Circle{radius: 5}
	s[3] = Circle{radius: 14}

	maxArea := s[0].Area()

	for i := range s {
		fmt.Printf("%.3f\n", s[i].Area())
		if maxArea < s[i].Area() {
			maxArea = s[i].Area()
		}
	}

	fmt.Printf("Max area: %.3f\n", maxArea)

}

func main() {

	var sh Shapes
	sh = Shapes{}
	sh.LargestArea()

}
