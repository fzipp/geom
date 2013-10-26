// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
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

func TestVec3Dot(t *testing.T) {
	tests := []struct {
		v, w Vec3
		want float32
	}{
		{V3(2, -3, 4), V3(-4, 2, 1), -10},
		{V3(4, 8, 3), V3(0.5, 1.25, -4.5), -1.5},
		{V3(12.5, 9.25, -2.5), V3Zero, 0},
		{V3UnitX, V3UnitY, 0},
		{V3(4, 5, 2), V3Unit, 11},
	}
	for _, tt := range tests {
		if x := tt.v.Dot(tt.w); x != tt.want {
			t.Errorf("%s.Dot(%s) = %g, want %g", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec3Cross(t *testing.T) {
	tests := []struct {
		v, w, want Vec3
	}{
		{V3(2, -3, 4), V3(-4, 2, 1), V3(-11, -18, -8)},
		{V3(4, 8, 3), V3(0.5, 1.25, -4.5), V3(-39.75, 19.5, 1)},
		{V3(12.5, 9.25, -2.5), V3Zero, V3Zero},
		{V3(3.5, -4, 9.6), V3(1, 0.5, 2), V3(-12.8, 2.6, 5.75)},
		{V3UnitX, V3UnitY, V3UnitZ},
		{V3(4, 5, 3), V3Unit, V3(2, -1, -1)},
	}
	for _, tt := range tests {
		if x := tt.v.Cross(tt.w); !x.NearEq(tt.want) {
			t.Errorf("%s.Cross(%s) = %s, want %s", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec3CompMul(t *testing.T) {
	tests := []struct {
		v, w Vec3
		want Vec3
	}{
		{V3(4, 1, 8), V3(2, 5, 3), V3(8, 5, 24)},
		{V3(1.2, 2.3, -2.1), V3(-2, 0.5, 3), V3(-2.4, 1.15, -6.3)},
		{V3(2, 3, 4), V3Unit, V3(2, 3, 4)},
		{V3(2, 3, 4), V3UnitX, V3(2, 0, 0)},
		{V3(2, 3, 4), V3UnitY, V3(0, 3, 0)},
		{V3(2, 3, 4), V3UnitZ, V3(0, 0, 4)},
		{V3(2, 3, 4), V3Zero, V3Zero},
	}
	for _, tt := range tests {
		if x := tt.v.CompMul(tt.w); !x.NearEq(tt.want) {
			t.Errorf("%s.CompMul(%s) = %s, want %s", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec3CompDiv(t *testing.T) {
	tests := []struct {
		v, w Vec3
		want Vec3
	}{
		{V3(4, 1, 8), V3(2, 5, -4), V3(2, 0.2, -2)},
		{V3(1.2, 2.3, -2.7), V3(-2, 0.5, 3), V3(-0.6, 4.6, -0.9)},
		{V3(2, 3, 4), V3Unit, V3(2, 3, 4)},
	}
	for _, tt := range tests {
		if x := tt.v.CompDiv(tt.w); !x.NearEq(tt.want) {
			t.Errorf("%s.CompDiv(%s) = %s, want %s", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec3SqDist(t *testing.T) {
	tests := []struct {
		v, w Vec3
		want float32
	}{
		{V3Zero, V3Zero, 0.0},
		{V3Zero, V3UnitX, 1.0},
		{V3UnitX, V3UnitY, 2.0},
		{V3(-2.3, 1.1, -8.4), V3(-2.3, 1.1, -8.4), 0.0},
		{V3(2, 1, 5), V3(2, 3, 1), 20},
		{V3(0.5, 2, 2.5), V3(1.5, 2.5, 3), 1.5},
	}
	for _, tt := range tests {
		if x := tt.v.SqDist(tt.w); !nearEq(x, tt.want, epsilon) {
			t.Errorf("%s.SqDist(%s) = %g, want %g", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec3Dist(t *testing.T) {
	tests := []struct {
		v, w Vec3
		want float32
	}{
		{V3Zero, V3Zero, 0.0},
		{V3Zero, V3UnitX, 1.0},
		{V3UnitX, V3UnitY, 1.4142135},
		{V3(-2.3, 1.1, -8.4), V3(-2.3, 1.1, -8.4), 0.0},
		{V3(2, 1, 5), V3(2, 3, 1), 4.472136},
		{V3(0.5, 2, 2.5), V3(1.5, 2.5, 3), 1.2247449},
	}
	for _, tt := range tests {
		if x := tt.v.Dist(tt.w); !nearEq(x, tt.want, epsilon) {
			t.Errorf("%s.Dist(%s) = %g, want %g", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec3Len(t *testing.T) {
	tests := []struct {
		v    Vec3
		want float32
	}{
		{V3Zero, 0.0},
		{V3Unit, 1.732050},
		{V3UnitX, 1.0},
		{V3UnitY, 1.0},
		{V3UnitZ, 1.0},
		{V3(-2.3, 1.1, 4.5), 5.17204},
		{V3(2, 1, 5), 5.477226},
		{V3(0.5, 2, -1), 2.291288},
	}
	for _, tt := range tests {
		if x := tt.v.Len(); !nearEq(x, tt.want, epsilon) {
			t.Errorf("%s.Len() = %g, want %g", tt.v, x, tt.want)
		}
	}
}

func TestVec3Norm(t *testing.T) {
	tests := []struct {
		v, want Vec3
	}{
		{V3UnitX, V3UnitX},
		{V3UnitY, V3UnitY},
		{V3UnitZ, V3UnitZ},
		{V3Unit, V3(0.57735026, 0.57735026, 0.57735026)},
		{V3(2, 4, 1), V3(0.43643576, 0.8728715, 0.21821788)},
		{V3(0.0, -2.3, 0.0), V3(0.0, -1.0, 0.0)},
		{V3(0.0, 0.0, 3.9), V3(0.0, 0.0, 1.0)},
		{V3(6.0, 0.0, 0.0), V3(1.0, 0.0, 0.0)},
		{V3(-2.5, 3.0, -2.0), V3(-0.5698029, 0.68376344, -0.45584232)},
	}
	for _, tt := range tests {
		if x := tt.v.Norm(); !x.NearEq(tt.want) {
			t.Errorf("%s.Norm() = %s, want %s", tt.v, x, tt.want)
		}
	}
}

func TestVec3Reflect(t *testing.T) {
	tests := []struct {
		v, n, want Vec3
	}{
		{V3UnitX, V3UnitY, V3(1, 0, 0)},
		{V3Unit, V3UnitY, V3(1, -1, 1)},
		{V3(2, 3, 1), V3UnitY, V3(2, -3, 1)},
		{V3(-1.5, -3.6, 1.2), V3UnitY, V3(-1.5, 3.6, 1.2)},
		{V3(1, 0.5, 2), V3Unit.Norm(), V3(-1.3333333, -1.8333333, -0.33333325)},
	}
	for _, tt := range tests {
		if x := tt.v.Reflect(tt.n); !x.NearEq(tt.want) {
			t.Errorf("%s.Reflect(%s) = %s, want %s", tt.v, tt.n, x, tt.want)
		}
	}
}

func TestVec3Lerp(t *testing.T) {
	tests := []struct {
		v, w Vec3
		t    float32
		want Vec3
	}{
		{V3Zero, V3UnitX, 0.0, V3Zero},
		{V3Zero, V3UnitX, 0.25, V3(0.25, 0, 0)},
		{V3Zero, V3UnitX, 1.0, V3UnitX},
		{V3(2, 1, 4), V3(4, 3, 8), 0.5, V3(3, 2, 6)},
	}
	for _, tt := range tests {
		if x := tt.v.Lerp(tt.w, tt.t); !x.NearEq(tt.want) {
			t.Errorf("%s.Lerp(%s, %g) = %s, want %s", tt.v, tt.w, tt.t, x, tt.want)
		}
	}
}

func TestVec3MinMax(t *testing.T) {
	tests := []struct {
		v, w, min, max Vec3
	}{
		{V3(2, 1, -4.3), V3(4, 3, 9.9), V3(2, 1, -4.3), V3(4, 3, 9.9)},
		{V3(2, 1, 8), V3(4, -3, -1), V3(2, -3, -1), V3(4, 1, 8)},
		{V3(5, 3.2, 0), V3(3.2, 1.4, 11), V3(3.2, 1.4, 0), V3(5, 3.2, 11)},
		{V3(0, 6, 321.6), V3(2, 3, 4), V3(0, 3, 4), V3(2, 6, 321.6)},
	}
	for _, tt := range tests {
		if x := tt.v.Min(tt.w); !x.NearEq(tt.min) {
			t.Errorf("%s.Min(%s) = %s, want %s", tt.v, tt.w, x, tt.min)
		}
		if x := tt.v.Max(tt.w); !x.NearEq(tt.max) {
			t.Errorf("%s.Max(%s) = %s, want %s", tt.v, tt.w, x, tt.max)
		}
	}
}

func TestVec3Transform(t *testing.T) {
	var rot, trans, scale Mat4
	rot.Id().Rot(&rot, math.Pi/2, V3UnitZ)
	trans.Id().Translate(&trans, V3(2.5, 3, -1))
	scale.Id().Scale(&scale, V3(2, 3, -4))

	tests := []struct {
		v    Vec3
		m    *Mat4
		want Vec3
	}{
		{V3(1, 0, 2), &rot, V3(0, 1, 2)},
		{V3(1, 2, 3), &trans, V3(3.5, 5, 2)},
		{V3(1.5, -3, -1), &scale, V3(3, -9, 4)},
	}
	for _, tt := range tests {
		if x := tt.v.Transform(tt.m); !x.NearEq(tt.want) {
			t.Errorf("%s.Transform(%v) = %s, want %s", tt.v, *tt.m, x, tt.want)
		}
	}
}
