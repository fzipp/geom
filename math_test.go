package geom

import (
	"testing"
)

func TestDegRad(t *testing.T) {
	tests := []struct {
		deg, rad float64
	}{
		{57.2957795, 1},
		{1, 0.0174532925},
	}
	for _, tt := range tests {
		if r := Rad(tt.deg); !nearEq(float32(r), float32(tt.rad), epsilon) {
			t.Errorf("Rad(%g) = %g, want %g", tt.deg, r, tt.rad)
		}
		if d := Deg(tt.rad); !nearEq(float32(d), float32(tt.deg), epsilon) {
			t.Errorf("Deg(%g) = %g, want %g", tt.rad, d, tt.deg)
		}
	}
}
