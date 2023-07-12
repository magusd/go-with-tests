package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, s Shape, want float64) {
		got := s.Perimeter()
		if got != want {
			t.Errorf("expected %.2f, got %.2f", want, got)
		}
	}
	t.Run("rectangular perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkPerimeter(t, rectangle, 40.0)
	})
	t.Run("circular permieter", func(t *testing.T) {
		circle := Circle{10.0}
		checkPerimeter(t, circle, math.Pi*10.0*2)
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, want: 72.0},
		{name: "Circle", shape: Circle{12}, want: math.Pi * 12.0 * 12.0},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}
	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("%#v wanted %g got %g", test, test.want, got)
			}
		})
	}
}
