package geometry

import (
	"fmt"
	"math"
)

func main() {
	a := make([]int, 5, 10)
	a[7] = 7

	s := a[:5]
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

type Areas struct {
	circle    float64
	rectangle float64
}

func Geometry() Areas {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Println("circle area:", circle.area())
	fmt.Println("rectangle area:", rectangle.area())

	areas := Areas{
		circle:    circle.area(),
		rectangle: rectangle.area(),
	}

	return areas
}

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
