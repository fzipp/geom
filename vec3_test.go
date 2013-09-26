// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"testing"
)

func TestVec3String(t *testing.T) {
	tests := []struct {
		v    Vec3
		want string
	}{
		{V3(-2.3, 1.1, 12.72), "(-2.3, 1.1, 12.72)"},
		{V3(2, 1, 4), "(2, 1, 4)"},
		{V3(0.5, 2, -1), "(0.5, 2, -1)"},
		{V3(1.414213, -34.0213, 651.0284), "(1.414213, -34.0213, 651.0284)"},
	}
	for _, tt := range tests {
		if s := tt.v.String(); s != tt.want {
			t.Errorf("(%g, %g, %g).String() = %q, want %q", tt.v.X, tt.v.Y, tt.v.Z, s, tt.want)
		}
	}
}

func TestVec3NearEq(t *testing.T) {
	tests := []struct {
		v, w Vec3
		want bool
	}{
		{V3(4, 1, 8), V3(4, 1, 8), true},
		{V3(-3.2145, -2.5667, -6.3487), V3(-3.2145, -2.5667, -6.3487), true},
		{V3(2.34567, -9.87654, 7.97433), V3(2.345669, -9.876541, 7.974329), true},
		{V3(4, 1, 6), V3(4, 1, 5), false},
		{V3(4, 1, 6), V3(4, 7, 6), false},
		{V3(4, 1, 6), V3(-3, 1, 6), false},
		{V3(2.34567, -9.87654, 5.43553), V3(2.34567, -9.87653, 5.43553), false},
		{V3(2.34567, -9.87654, 5.43553), V3(2.34568, -9.87654, 5.43553), false},
		{V3(2.34567, -9.87654, 5.43553), V3(2.34567, -9.87654, 5.43554), false},
	}
	for _, tt := range tests {
		if x := tt.v.NearEq(tt.w); x != tt.want {
			t.Errorf("%s.NearEq(%s) = %v, want %v", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec3Add(t *testing.T) {
	tests := []struct {
		v, w Vec3
		want Vec3
	}{
		{V3(4, 1, 8), V3(2, 5, 3), V3(6, 6, 11)},
		{V3(1.2, 2.3, -2.7), V3(-2.1, 0.5, -1.3), V3(-0.9, 2.8, -4)},
		{V3(12.5, 9.25, 44.2), V3Zero, V3(12.5, 9.25, 44.2)},
		{V3UnitX, V3UnitY, V3(1, 1, 0)},
		{V3UnitX, V3UnitZ, V3(1, 0, 1)},
		{V3UnitY, V3UnitZ, V3(0, 1, 1)},
	}
	for _, tt := range tests {
		if x := tt.v.Add(tt.w); !x.NearEq(tt.want) {
			t.Errorf("%s + %s = %s, want %s", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec3Sub(t *testing.T) {
	tests := []struct {
		v, w Vec3
		want Vec3
	}{
		{V3(4, 1, 8), V3(2, 5, 3), V3(2, -4, 5)},
		{V3(1.2, 2.3, -2.7), V3(-2.1, 0.5, -1.3), V3(3.3, 1.8, -1.4)},
		{V3(12.5, 9.25, 44.2), V3Zero, V3(12.5, 9.25, 44.2)},
	}
	for _, tt := range tests {
		if x := tt.v.Sub(tt.w); !x.NearEq(tt.want) {
			t.Errorf("%s - %s = %s, want %s", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec3Mul(t *testing.T) {
	tests := []struct {
		v    Vec3
		s    float32
		want Vec3
	}{
		{V3(4, 1, 8), 2, V3(8, 2, 16)},
		{V3(1.4, -2.5, 7), 0.5, V3(0.7, -1.25, 3.5)},
		{V3(12.5, 9.25, 56.142), 1, V3(12.5, 9.25, 56.142)},
		{V3(2.7, 1.1, 34.6), 0, V3Zero},
	}
	for _, tt := range tests {
		if x := tt.v.Mul(tt.s); !x.NearEq(tt.want) {
			t.Errorf("%g * %s = %s, want %s", tt.s, tt.v, x, tt.want)
		}
	}
}

func TestVec3Div(t *testing.T) {
	tests := []struct {
		v    Vec3
		s    float32
		want Vec3
	}{
		{V3(8, 2, 16), 2, V3(4, 1, 8)},
		{V3(0.7, -1.25, 3.5), 0.5, V3(1.4, -2.5, 7)},
		{V3(12.5, 9.25, 56.142), 1, V3(12.5, 9.25, 56.142)},
	}
	for _, tt := range tests {
		if x := tt.v.Div(tt.s); !x.NearEq(tt.want) {
			t.Errorf("%s / %g = %s, want %s", tt.v, tt.s, x, tt.want)
		}
	}
}

func TestVec3Neg(t *testing.T) {
	tests := []struct {
		v, want Vec3
	}{
		{V3(4, 1, 8), V3(-4, -1, -8)},
		{V3(-1.2, 2.3, -2.7), V3(1.2, -2.3, 2.7)},
		{V3(1.2, -2.3, 2.7), V3(-1.2, 2.3, -2.7)},
		{V3(-12.5, -9.25, 56.142), V3(12.5, 9.25, -56.142)},
		{V3Zero, V3Zero},
	}
	for _, tt := range tests {
		if x := tt.v.Neg(); !x.NearEq(tt.want) {
			t.Errorf("%s.Neg() = %s, want %s", tt.v, x, tt.want)
		}
	}
}
