// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
	"testing"
)

func TestVec2String(t *testing.T) {
	tests := []struct {
		v    Vec2
		want string
	}{
		{V2(-2.3, 1.1), "(-2.3, 1.1)"},
		{V2(2, 1), "(2, 1)"},
		{V2(0.5, 2), "(0.5, 2)"},
		{V2(1.414213, -34.0213), "(1.414213, -34.0213)"},
	}
	for _, tt := range tests {
		if s := tt.v.String(); s != tt.want {
			t.Errorf("(%g, %g).String() = %q, want %q", tt.v.X, tt.v.Y, s, tt.want)
		}
	}
}

func TestVec2NearEq(t *testing.T) {
	tests := []struct {
		v, w Vec2
		want bool
	}{
		{V2(4, 1), V2(4, 1), true},
		{V2(-3.2145, -2.5667), V2(-3.2145, -2.5667), true},
		{V2(2.34567, -9.87654), V2(2.345669, -9.876541), true},
		{V2(4, 1), V2(-3, 7), false},
		{V2(2.34567, -9.87654), V2(2.34567, -9.87653), false},
		{V2(2.34567, -9.87654), V2(2.34568, -9.87654), false},
	}
	for _, tt := range tests {
		if x := tt.v.NearEq(tt.w); x != tt.want {
			t.Errorf("%s.NearEq(%s) = %v, want %v", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec2Add(t *testing.T) {
	tests := []struct {
		v, w Vec2
		want Vec2
	}{
		{V2(4, 1), V2(2, 5), V2(6, 6)},
		{V2(1.2, 2.3), V2(-2.1, 0.5), V2(-0.9, 2.8)},
		{V2(12.5, 9.25), V2Zero, V2(12.5, 9.25)},
		{V2UnitX, V2UnitY, V2Unit},
	}
	for _, tt := range tests {
		if x := tt.v.Add(tt.w); !x.NearEq(tt.want) {
			t.Errorf("%s + %s = %s, want %s", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec2Sub(t *testing.T) {
	tests := []struct {
		v, w Vec2
		want Vec2
	}{
		{V2(4, 1), V2(2, 5), V2(2, -4)},
		{V2(1.2, 2.3), V2(-2.1, 0.5), V2(3.3, 1.8)},
		{V2(12.5, 9.25), V2Zero, V2(12.5, 9.25)},
	}
	for _, tt := range tests {
		if x := tt.v.Sub(tt.w); !x.NearEq(tt.want) {
			t.Errorf("%s - %s = %s, want %s", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec2Mul(t *testing.T) {
	tests := []struct {
		v    Vec2
		s    float32
		want Vec2
	}{
		{V2(4, 1), 2, V2(8, 2)},
		{V2(1.4, -2.5), 0.5, V2(0.7, -1.25)},
		{V2(12.5, 9.25), 1, V2(12.5, 9.25)},
		{V2(2.7, 1.1), 0, V2Zero},
	}
	for _, tt := range tests {
		if x := tt.v.Mul(tt.s); !x.NearEq(tt.want) {
			t.Errorf("%g * %s = %s, want %s", tt.s, tt.v, x, tt.want)
		}
	}
}

func TestVec2Div(t *testing.T) {
	tests := []struct {
		v    Vec2
		s    float32
		want Vec2
	}{
		{V2(8, 2), 2, V2(4, 1)},
		{V2(0.7, -1.25), 0.5, V2(1.4, -2.5)},
		{V2(12.5, 9.25), 1, V2(12.5, 9.25)},
	}
	for _, tt := range tests {
		if x := tt.v.Div(tt.s); !x.NearEq(tt.want) {
			t.Errorf("%s / %g = %s, want %s", tt.v, tt.s, x, tt.want)
		}
	}
}

func TestVec2Neg(t *testing.T) {
	tests := []struct {
		v, want Vec2
	}{
		{V2(4, 1), V2(-4, -1)},
		{V2(-1.2, 2.3), V2(1.2, -2.3)},
		{V2(1.2, -2.3), V2(-1.2, 2.3)},
		{V2(-12.5, -9.25), V2(12.5, 9.25)},
		{V2Zero, V2Zero},
	}
	for _, tt := range tests {
		if x := tt.v.Neg(); !x.NearEq(tt.want) {
			t.Errorf("%s.Neg() = %s, want %s", tt.v, x, tt.want)
		}
	}
}

func TestVec2Dot(t *testing.T) {
	tests := []struct {
		v, w Vec2
		want float32
	}{
		{V2(2, -3), V2(-4, 2), -14},
		{V2(4, 8), V2(0.5, 1.25), 12},
		{V2(12.5, 9.25), V2Zero, 0},
		{V2UnitX, V2UnitY, 0},
		{V2(4, 5), V2Unit, 9},
	}
	for _, tt := range tests {
		if x := tt.v.Dot(tt.w); x != tt.want {
			t.Errorf("%s.Dot(%s) = %g, want %g", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec2CrossLen(t *testing.T) {
	tests := []struct {
		v, w Vec2
		want float32
	}{
		{V2(2, -3), V2(-4, 2), -8},
		{V2(4, 8), V2(0.5, 1.25), 1},
		{V2(12.5, 9.25), V2Zero, 0},
		{V2(3.5, -4), V2(1, 0.5), 5.75},
		{V2UnitX, V2UnitY, 1},
		{V2(4, 5), V2Unit, -1},
	}
	for _, tt := range tests {
		if x := tt.v.CrossLen(tt.w); x != tt.want {
			t.Errorf("%s.CrossLen(%s) = %g, want %g", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec2CompMul(t *testing.T) {
	tests := []struct {
		v, w Vec2
		want Vec2
	}{
		{V2(4, 1), V2(2, 5), V2(8, 5)},
		{V2(1.2, 2.3), V2(-2, 0.5), V2(-2.4, 1.15)},
		{V2(2, 3), V2Unit, V2(2, 3)},
		{V2(2, 3), V2UnitX, V2(2, 0)},
		{V2(2, 3), V2UnitY, V2(0, 3)},
		{V2(2, 3), V2Zero, V2Zero},
	}
	for _, tt := range tests {
		if x := tt.v.CompMul(tt.w); !x.NearEq(tt.want) {
			t.Errorf("%s.CompMul(%s) = %s, want %s", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec2CompDiv(t *testing.T) {
	tests := []struct {
		v, w Vec2
		want Vec2
	}{
		{V2(4, 1), V2(2, 5), V2(2, 0.2)},
		{V2(1.2, 2.3), V2(-2, 0.5), V2(-0.6, 4.6)},
		{V2(2, 3), V2Unit, V2(2, 3)},
	}
	for _, tt := range tests {
		if x := tt.v.CompDiv(tt.w); !x.NearEq(tt.want) {
			t.Errorf("%s.CompDiv(%s) = %s, want %s", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec2SqDist(t *testing.T) {
	tests := []struct {
		v, w Vec2
		want float32
	}{
		{V2Zero, V2Zero, 0.0},
		{V2Zero, V2UnitX, 1.0},
		{V2UnitX, V2UnitY, 2.0},
		{V2(-2.3, 1.1), V2(-2.3, 1.1), 0.0},
		{V2(2, 1), V2(2, 3), 4.0},
		{V2(0.5, 2), V2(1.5, 2.5), 1.25},
	}
	for _, tt := range tests {
		if x := tt.v.SqDist(tt.w); !nearEq(x, tt.want, epsilon) {
			t.Errorf("%s.SqDist(%s) = %g, want %g", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec2Dist(t *testing.T) {
	tests := []struct {
		v, w Vec2
		want float32
	}{
		{V2Zero, V2Zero, 0.0},
		{V2Zero, V2UnitX, 1.0},
		{V2UnitX, V2UnitY, 1.4142135},
		{V2(-2.3, 1.1), V2(-2.3, 1.1), 0.0},
		{V2(2, 1), V2(2, 3), 2.0},
		{V2(0.5, 2), V2(1.5, 2.5), 1.118034},
	}
	for _, tt := range tests {
		if x := tt.v.Dist(tt.w); !nearEq(x, tt.want, epsilon) {
			t.Errorf("%s.Dist(%s) = %g, want %g", tt.v, tt.w, x, tt.want)
		}
	}
}

func TestVec2Len(t *testing.T) {
	tests := []struct {
		v    Vec2
		want float32
	}{
		{V2Zero, 0.0},
		{V2Unit, 1.4142135},
		{V2UnitX, 1.0},
		{V2UnitY, 1.0},
		{V2(-2.3, 1.1), 2.5495098},
		{V2(2, 1), 2.236068},
		{V2(0.5, 2), 2.0615528},
	}
	for _, tt := range tests {
		if x := tt.v.Len(); !nearEq(x, tt.want, epsilon) {
			t.Errorf("%s.Len() = %g, want %g", tt.v, x, tt.want)
		}
	}
}

func TestVec2Norm(t *testing.T) {
	tests := []struct {
		v, want Vec2
	}{
		{V2UnitX, V2UnitX},
		{V2UnitY, V2UnitY},
		{V2Unit, V2(0.70710677, 0.70710677)},
		{V2(2, 4), V2(0.4472136, 0.8944272)},
		{V2(0.0, -2.3), V2(0.0, -1.0)},
		{V2(-2.5, 3.0), V2(-0.6401844, 0.76822126)},
	}
	for _, tt := range tests {
		if x := tt.v.Norm(); !x.NearEq(tt.want) {
			t.Errorf("%s.Norm() = %s, want %s", tt.v, x, tt.want)
		}
	}
}

func TestVec2Reflect(t *testing.T) {
	tests := []struct {
		v, n, want Vec2
	}{
		{V2UnitX, V2UnitY, V2(1, 0)},
		{V2Unit, V2UnitY, V2(1, -1)},
		{V2(2, 3), V2UnitY, V2(2, -3)},
		{V2(-1.5, -3.6), V2UnitY, V2(-1.5, 3.6)},
		{V2(1, 0.5), V2Unit.Norm(), V2(-0.5, -1)},
	}
	for _, tt := range tests {
		if x := tt.v.Reflect(tt.n); !x.NearEq(tt.want) {
			t.Errorf("%s.Reflect(%s) = %s, want %s", tt.v, tt.n, x, tt.want)
		}
	}
}

func TestVec2Lerp(t *testing.T) {
	tests := []struct {
		v, w Vec2
		t    float32
		want Vec2
	}{
		{V2Zero, V2UnitX, 0.0, V2Zero},
		{V2Zero, V2UnitX, 0.25, V2(0.25, 0)},
		{V2Zero, V2UnitX, 1.0, V2UnitX},
		{V2(2, 1), V2(4, 3), 0.5, V2(3, 2)},
	}
	for _, tt := range tests {
		if x := tt.v.Lerp(tt.w, tt.t); !x.NearEq(tt.want) {
			t.Errorf("%s.Lerp(%s, %g) = %s, want %s", tt.v, tt.w, tt.t, x, tt.want)
		}
	}
}

func TestVec2Angle(t *testing.T) {
	tests := []struct {
		v    Vec2
		want float32
	}{
		{V2(1, 0), 0},
		{V2(1, 1), math.Pi * 1 / 4},
		{V2(0, 1), math.Pi * 2 / 4},
		{V2(-1, 1), math.Pi * 3 / 4},
		{V2(-1, 0), math.Pi * 4 / 4},
		{V2(-1, -1), math.Pi * 5 / 4},
		{V2(0, -1), math.Pi * 6 / 4},
		{V2(1, -1), math.Pi * 7 / 4},
	}
	for _, tt := range tests {
		if x := tt.v.angle(); x != tt.want {
			t.Errorf("%s.Angle() = %g, want %g", tt.v, x, tt.want)
		}
	}
}

func TestVec2MinMax(t *testing.T) {
	tests := []struct {
		v, w, min, max Vec2
	}{
		{V2(2, 1), V2(4, 3), V2(2, 1), V2(4, 3)},
		{V2(2, 1), V2(4, -3), V2(2, -3), V2(4, 1)},
		{V2(5, 3.2), V2(3.2, 1.4), V2(3.2, 1.4), V2(5, 3.2)},
		{V2(0, 6), V2(2, 3), V2(0, 3), V2(2, 6)},
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

func TestVec2Transform(t *testing.T) {
	var rot, trans, scale Mat4
	rot.Id().Rot(&rot, math.Pi/2, V3UnitZ)
	trans.Id().Translate(&trans, V3(2.5, 3, 0))
	scale.Id().Scale(&scale, V3(2, 3, 0))

	tests := []struct {
		v    Vec2
		m    *Mat4
		want Vec2
	}{
		{V2(1, 0), &rot, V2(0, 1)},
		{V2(1, 2), &trans, V2(3.5, 5)},
		{V2(1.5, -3), &scale, V2(3, -9)},
	}
	for _, tt := range tests {
		if x := tt.v.Transform(tt.m); !x.NearEq(tt.want) {
			t.Errorf("%s.Transform(%v) = %s, want %s", tt.v, *tt.m, x, tt.want)
		}
	}
}

func TestVec2Z(t *testing.T) {
	tests := []struct {
		v    Vec2
		z    float32
		want Vec3
	}{
		{V2(0, -2.3), 3.1, V3(0, -2.3, 3.1)},
		{V2(-2.5, 3), -1.4, V3(-2.5, 3, -1.4)},
	}
	for _, tt := range tests {
		if x := tt.v.Z(tt.z); x != tt.want {
			t.Errorf("%s.Z(%g) = %s, want %s", tt.v, tt.z, x, tt.want)
		}
	}
}
