package shape

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Name() string
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Name() string {
	return "rectangle"
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Name() string {
	return "circle"
}

func PrintArea(shape Shape) {
	fmt.Printf("Area of %s is %f\n", shape.Name(), shape.Area())
}

func Describe(input Shape) {
	switch value := input.(type) {
	case Circle:
		fmt.Printf("Your circle area is: %f\n", value.Area())
	case Rectangle:
		fmt.Printf("Your rectangle area is: %f\n", value.Area())
	default:
		fmt.Println("shape is not recognized", value)
	}
}
