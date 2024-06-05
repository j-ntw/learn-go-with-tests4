package main

import "math"

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}
