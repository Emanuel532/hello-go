package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	areaTests := map[string]struct {
		input  Shape
		output float64
	}{
		"rectangle area": {
			input:  Rectangle{Height: 5.0, Width: 5.0},
			output: 25.0,
		},
		"circle area": {
			input:  Circle{Radius: 5.0},
			output: 78.53981633974483,
		},
		"triangle area": {
			input:  Triangle{Height: 12, Base: 6},
			output: 36.0},
	}
	for name, test := range areaTests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			checkArea(t, test.input, test.output)
		})

	}

}
