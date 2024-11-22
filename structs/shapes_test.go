package shapes

import "testing"

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("%#v expected %g got %g", shape, got, want)
		}
	}

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, want: 100.0},
		{name: "Circle", shape: Circle{10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{10.0, 10.0}, want: 50.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("Expected %.2f got %.2f", got, want)
	}
}
