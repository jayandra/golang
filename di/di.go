package main

import (
	"fmt"
	"math"
)

// Shape is implemented by anything that can tell us its Area.
type Shape interface {
	Area() float64
}

// Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func NewRectangle() *Rectangle {
	var length float64
	var widht float64
	fmt.Println("Enter length:")
	fmt.Scanf("%f", &length)
	fmt.Println("Enter width:")
	fmt.Scanf("%f", &widht)
	return &Rectangle{
		Width:  length,
		Height: widht,
	}
}

// Circle represents a circle...
type Circle struct {
	Radius float64
}

// Area returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func NewCircle() *Circle {
	var r float64
	fmt.Println("Enter radius:")
	fmt.Scanf("%f", &r)
	return &Circle{
		Radius: r,
	}
}

// Triangle represents the dimensions of a triangle.
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of the triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func NewTriangle() *Triangle {
	var base float64
	var height float64
	fmt.Println("Enter Base:")
	fmt.Scanf("%f", &base)
	fmt.Println("Enter Height:")
	fmt.Scanf("%f", &height)
	return &Triangle{
		Base:   base,
		Height: height,
	}
}

// Because of this dependency on interface AreaComputer is not tied to any particular type
// As long as the initialized type adhers to Shape interface; it can be swapped-in on the fly
type AreaComputer struct {
	s Shape
}

func (ac AreaComputer) ComputeArea() float64 {
	return ac.s.Area()
}

func main() {
	var user_choice string
	var shape Shape
	exit_loop := false

	for !exit_loop {
		fmt.Println("\n\nEnter the shape for AreaComputer: [Circle, Triangle, Rectangle] or exit")
		fmt.Scanln(&user_choice)

		switch user_choice {
		case "Circle":
			shape = NewCircle()
		case "Triangle":
			shape = NewTriangle()
		case "Rectangle":
			shape = NewRectangle()
		case "exit":
			return
		}

		fmt.Printf("The Area is: %.2f", AreaComputer{shape}.ComputeArea())
	}
}
