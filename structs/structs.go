package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return 0
}

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return 0
}

type Shape interface { // All of the structs which have the following methods are considered as Shape
	Area() float64
	Perimeter() float64
}

// func Perimeter(rectangle Rectangle) float64 {
// 	return 2 * (rectangle.Width + rectangle.Height)
// }
//
// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }
