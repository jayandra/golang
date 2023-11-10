package main

import (
	"testing"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}

func TestComputeArea(t *testing.T) {
	rect := Rectangle{Width: 12, Height: 6}
	circ := Circle{Radius: 10}
	tria := Triangle{Base: 12, Height: 6}

	tests := []struct {
		name string
		ac   AreaComputer
		want float64
	}{
		{name: "Rectangle", ac: AreaComputer{s: rect}, want: 72.0},
		{name: "Circle", ac: AreaComputer{s: circ}, want: 314.1592653589793},
		{name: "Triangle", ac: AreaComputer{s: tria}, want: 36.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ac.s.Area()
			if got != tt.want {
				t.Errorf("%s: got %g, want %g", tt.name, got, tt.want)
			}
		})
	}
}
