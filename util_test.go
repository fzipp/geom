// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import "testing"

func TestDegRad(t *testing.T) {
	tests := []struct {
		deg, rad float32
	}{
		{57.2957795, 1},
		{1, 0.0174532925},
	}
	for _, tt := range tests {
		if r := Rad(tt.deg); !nearEq(r, tt.rad, epsilon) {
			t.Errorf("Rad(%g) = %g, want %g", tt.deg, r, tt.rad)
		}
		if d := Deg(tt.rad); !nearEq(d, tt.deg, epsilon) {
			t.Errorf("Deg(%g) = %g, want %g", tt.rad, d, tt.deg)
		}
	}
}
