package main

import "math"

type Shape interface {
	Area() float64
}
 
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * math.Pi * c.Radius
}
