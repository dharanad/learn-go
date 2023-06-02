package shape

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Length)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 2 * math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func Perimeter(s Shape) float64 {
	return s.Perimeter()
}
