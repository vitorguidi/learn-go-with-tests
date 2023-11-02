package interfaces

import (
	"math"
	"testing"
)

type Result struct {
	Shape Shape
	Want  float64
}

func Test_Areas(t *testing.T) {
	results := []Result{
		{
			Shape: Circle{Radius: 3},
			Want:  math.Pi * 9.0,
		},
		{
			Shape: Rectangle{Width: 2.0, Height: 3.0},
			Want:  6.0,
		},
	}
	for _, result := range results {
		got := result.Shape.Area()
		expected := result.Want
		if got != expected {
			t.Errorf("Expected %f, got %f", expected, got)
		}
	}
}

func Test_Perimeters(t *testing.T) {
	results := []Result{
		{
			Shape: Circle{Radius: 3.0},
			Want:  math.Pi * 6.0,
		},
		{
			Shape: Rectangle{Width: 3.0, Height: 4.0},
			Want:  14.0,
		},
	}
	for _, result := range results {
		got := result.Shape.Perimeter()
		expected := result.Want
		if got != expected {
			t.Errorf("Expected %f, got %f", expected, got)
		}
	}
}
