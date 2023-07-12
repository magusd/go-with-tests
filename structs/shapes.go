package structs

import "math"

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Length + r.Width) * 2
}
func (c Circle) Perimeter() float64 {
	return c.Radius * math.Pi * 2
}
func (t Triangle) Perimeter() float64 {
	return 0.0
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
