package shape

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	tests := []struct {
		shape Shape
		want  float64
	}{
		{&Rectangle{10, 10}, 40.0},
		{&Circle{10}, 62.83},
	}
	for _, tt := range tests {
		got := Perimeter(tt.shape)
		if math.Round(got) != math.Round(tt.want) {
			t.Errorf("got %f want %f", got, tt.want)
		}
	}
}
