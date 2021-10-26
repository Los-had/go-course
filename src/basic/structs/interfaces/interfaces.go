package main

import (
	"fmt"
)

type shape interface {
	area() float64
}

type square struct {
	Height float64
	Width  float64
}

type rectangle struct {
	Height float64
	Width  float64
}

func (s square) area() float64 {
	return s.Width * s.Height
}

func (r rectangle) area() float64 {
	return r.Width * r.Height
}

func main() {
	s1 := square{4, 4}
	r2 := rectangle{4, 4}
	shapes := []shape{s1, r2}

	fmt.Println(shapes[1].area())
	fmt.Println(shapes[0].area())
}