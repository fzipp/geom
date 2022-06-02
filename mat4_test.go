// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
	"testing"
)

func TestMat4NearEq(t *testing.T) {
	tests := []struct {
		a, b Mat4
		want bool
	}{
		{Mat4{
			{11, 12, 13, 14},
			{21, 22, 23, 24},
			{31, 32, 33, 34},
			{41, 42, 43, 44},
		}, Mat4{
			{11, 12, 13, 14},
			{21, 22, 23, 24},
			{31, 32, 33, 34},
			{41, 42, 43, 44},
		}, true},

		{Mat4{
			{11, 12, 13, 14},
			{21, 22, 23, 24},
			{31, 32, 33, 34},
			{41, 42, 43, 44},
		}, Mat4{
			{11.000001, 12.000001, 13.000001, 14.000001},
			{21.000001, 22.000001, 23.000001, 24.000001},
			{31.000001, 32.000001, 33.000001, 34.000001},
			{41.000001, 42.000001, 43.000001, 44.000001},
		}, true},

		{Mat4{
			{11, 12, 13, 14},
			{21, 22, 23, 24},
			{31, 32, 33, 34},
			{41, 42, 43, 44},
		}, Mat4{
			{10.999999, 11.999999, 12.999999, 13.999999},
			{20.999999, 21.999999, 22.999999, 23.999999},
			{30.999999, 31.999999, 32.999999, 33.999999},
			{40.999999, 41.999999, 42.999999, 43.999999},
		}, true},

		{Mat4{
			{11, 12, 13, 14},
			{21, 22, 23, 24},
			{31, 32, 33, 34},
			{41, 42, 43, 44},
		}, Mat4{
			{10.99999, 11.99999, 12.99999, 13.99999},
			{20.99999, 21.99999, 22.99999, 23.99999},
			{30.99999, 31.99999, 32.99999, 33.99999},
			{40.99999, 41.99999, 42.99999, 43.99999},
		}, false},
	}
	for _, tt := range tests {
		x := tt.a.nearEq(&tt.b)
		if x != tt.want {
			t.Errorf("%v.nearEq(%v) = %v, want %v", tt.a, tt.b, x, tt.want)
		}
	}
}

func TestMat4ID(t *testing.T) {
	m := Mat4{
		{11, 12, 13, 14},
		{21, 22, 23, 24},
		{31, 32, 33, 34},
		{41, 42, 43, 44},
	}
	id := Mat4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	mp := m.ID()
	if m != id {
		t.Errorf("m.ID() does not set m to the identity matrix, got instead: %v", m)
	}
	if mp != &m {
		t.Errorf("m.ID() does not return the pointer to m")
	}
}

func TestMat4Zero(t *testing.T) {
	m := Mat4{
		{11, 12, 13, 14},
		{21, 22, 23, 24},
		{31, 32, 33, 34},
		{41, 42, 43, 44},
	}
	id := Mat4{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	mp := m.Zero()
	if m != id {
		t.Errorf("m.Zero() does not set m to the zero matrix, got instead: %v", m)
	}
	if mp != &m {
		t.Errorf("m.Zero() does not return the pointer to m")
	}
}

func TestMat4Det(t *testing.T) {
	tests := []struct {
		m    Mat4
		want float32
	}{
		{id, 1},
		{Mat4{
			{-3, 2, 6, 5},
			{4, 1.5, 1, 8},
			{1, 4, 2, 4},
			{5.25, 6, -2, 8},
		}, 348.25},
	}
	for _, tt := range tests {
		if det := tt.m.Det(); det != tt.want {
			t.Errorf("%v.Det() = %g, want %g", tt.m, det, tt.want)
		}
	}
}

func TestMat4Ortho(t *testing.T) {
	tests := []struct {
		l, r, b, t, n, f float32
		want             Mat4
	}{
		{-1, 1, -1, 1, 1, -1, Mat4{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		}},
		{-2, 2, -2, 2, 2, -2, Mat4{
			{0.5, 0, 0, 0},
			{0, 0.5, 0, 0},
			{0, 0, 0.5, 0},
			{0, 0, 0, 1.0},
		}},
		{1, 2, 3, 4, 5, 6, Mat4{
			{2, 0, 0, 0},
			{0, 2, 0, 0},
			{0, 0, -2, 0},
			{-3, -7, -11, 1},
		}},
	}
	var m Mat4
	for _, tt := range tests {
		m.Ortho(tt.l, tt.r, tt.b, tt.t, tt.n, tt.f)
		if m != tt.want {
			t.Errorf("m.Ortho(%g, %g, %g, %g, %g, %g) = %v, want %v",
				tt.l, tt.r, tt.b, tt.t, tt.n, tt.f, m, tt.want)
		}
	}
}

func TestMat4Frustum(t *testing.T) {
	tests := []struct {
		l, r, b, t, n, f float32
		want             Mat4
	}{
		{-1, 1, -1, 1, 1, -1, Mat4{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, -1},
			{0, 0, -1, 0},
		}},
		{-2, 2, -2, 2, 2, -2, Mat4{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, -1},
			{0, 0, -2, 0},
		}},
		{1, 2, 3, 4, 5, 6, Mat4{
			{10, 0, 0, 0},
			{0, 10, 0, 0},
			{3, 7, -11, -1},
			{0, 0, -60, 0},
		}},
	}
	var m Mat4
	for _, tt := range tests {
		m.Frustum(tt.l, tt.r, tt.b, tt.t, tt.n, tt.f)
		if m != tt.want {
			t.Errorf("m.Frustum(%g, %g, %g, %g, %g, %g) = %v, want %v",
				tt.l, tt.r, tt.b, tt.t, tt.n, tt.f, m, tt.want)
		}
	}
}

func TestMat4Perspective(t *testing.T) {
	tests := []struct {
		fovy, a, n, f float32
		want          Mat4
	}{
		{65, 1.5, 5, 10, Mat4{
			{0.35279062, 0, 0, 0},
			{0, 0.52918595, 0, 0},
			{0, 0, -3, -1},
			{0, 0, -20, 0},
		}},
		{70, 1.33, 1, 30, Mat4{
			{1.5868644, 0, 0, 0},
			{0, 2.1105297, 0, 0},
			{0, 0, -1.0689656, -1},
			{0, 0, -2.0689654, 0},
		}},
	}
	var m Mat4
	for _, tt := range tests {
		m.Perspective(tt.fovy, tt.a, tt.n, tt.f)
		if m != tt.want {
			t.Errorf("m.Perspective(%g, %g, %g, %g) = %v, want %v",
				tt.fovy, tt.a, tt.n, tt.f, m, tt.want)
		}
	}
}

func TestMat4LookAt(t *testing.T) {
	tests := []struct {
		eye, center, up Vec3
		want            Mat4
	}{
		{V3(1, 1, 1), V3(1, 1, 0), V3(0, 1, 0), Mat4{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{-1, -1, -1, 1},
		}},
		{V3(0, 0, 1), V3(0, 0, -1), V3(0, 1, 0), Mat4{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, -1, 1},
		}},
		{V3(20, 80, 15), V3(15, 0, 12), V3(0, -1, 0), Mat4{
			{-0.5144958, 0.8552243, 0.06233464, 0},
			{0, -0.07269406, 0.99735427, 0},
			{0.857493, 0.5131346, 0.037400786, 0},
			{-2.5724783, -18.985981, -81.596054, 1},
		}},
	}
	var m Mat4
	for _, tt := range tests {
		m.LookAt(tt.eye, tt.center, tt.up)
		if !tt.want.nearEq(&m) {
			t.Errorf("m.LookAt(%s, %s, %s) = %v, want %v",
				tt.eye, tt.center, tt.up, m, tt.want)
		}
	}
}

func TestMat4Floats(t *testing.T) {
	m := Mat4{
		{11, 12, 13, 14},
		{21, 22, 23, 24},
		{31, 32, 33, 34},
		{41, 42, 43, 44},
	}
	want := [16]float32{11, 12, 13, 14, 21, 22, 23, 24, 31, 32, 33, 34, 41, 42, 43, 44}
	f := m.Floats()
	if *f != want {
		t.Errorf("%v.Floats() = %v, want %v", m, *f, want)
	}
	f[6] = 99
	if m[1][2] != 99 {
		t.Errorf("Pointer to float32 array returned by Floats() does not point to matrix data.")
	}
}

func TestMat4Mul(t *testing.T) {
	tests := []struct {
		a, b, want Mat4
	}{
		{Mat4{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		}, Mat4{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		}, Mat4{
			{34, 44, 54, 64},
			{82, 108, 134, 160},
			{34, 44, 54, 64},
			{82, 108, 134, 160},
		}},

		{Mat4{
			{2.1, 3.2, 1.2, 0},
			{4.6, -5.3, 5.4, 8.4},
			{-9.1, 1, 0.2, -7.3},
			{-1.25, 2.2, 2.3, 6.3},
		}, Mat4{
			{-3, 2.4, 8.4, 3.3},
			{0.2, -5, 2.6, 1.2},
			{5.3, 8.1, 4.4, 2.5},
			{4.9, -1, 0, 4},
		}, Mat4{
			{-75.825, -6.66, 18.63, -20.37},
			{-47.74, 32.38, -23.48, -53.42},
			{5.225, -16.07, 56.73, 51.67},
			{0.69, 29.78, 9.68, 16.8},
		}},
	}
	for _, tt := range tests {
		var m Mat4
		mp := m.Mul(&tt.a, &tt.b)
		if !tt.want.nearEq(&m) {
			t.Errorf("%v * %v = %v, want %v", tt.a, tt.b, m, tt.want)
		}
		if mp != &m {
			t.Errorf("m.Mul(...) does not return the pointer to m")
		}
	}
}

func TestMat4Rot(t *testing.T) {
	tests := []struct {
		a    Mat4
		rad  float32
		axis Vec3
		want Mat4
	}{
		{Mat4{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		}, math.Pi / 4, V3(0, 0, 1), Mat4{
			{0.7071, 0.7071, 0, 0},
			{-0.7071, 0.7071, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		}},
		{Mat4{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{1, 2, 3, 1},
		}, math.Pi * 0.5, V3(1, 0, 0), Mat4{
			{1, 0, 0, 0},
			{0, 0, 1, 0},
			{0, -1, 0, 0},
			{1, 2, 3, 1},
		}},
	}
	for _, tt := range tests {
		var m Mat4
		mp := m.Rot(&tt.a, tt.rad, tt.axis)
		if !tt.want.nearEq(&m) {
			t.Errorf("m.Rot(%v, %g, %s) = %v, want %v", tt.a, tt.rad, tt.axis, m, tt.want)
		}
		if mp != &m {
			t.Errorf("m.Rot(...) does not return the pointer to m")
		}
	}
}

func TestMat4Scale(t *testing.T) {
	tests := []struct {
		a    Mat4
		v    Vec3
		want Mat4
	}{
		{Mat4{
			{3, 4, 1, 0},
			{8, 3, 2.5, 7},
			{-2, 2, 1.3, 0},
			{1, 5.6, 2, -4},
		}, V3(2, 0.5, -1), Mat4{
			{6, 8, 2, 0},
			{4, 1.5, 1.25, 3.5},
			{2, -2, -1.3, 0},
			{1, 5.6, 2, -4},
		}},

		{Mat4{
			{3, 4, 1, 0},
			{8, 3, 2.5, 7},
			{-2, 2, 1.3, 0},
			{1, 5.6, 2, -4},
		}, V3Unit, Mat4{
			{3, 4, 1, 0},
			{8, 3, 2.5, 7},
			{-2, 2, 1.3, 0},
			{1, 5.6, 2, -4},
		}},

		{Mat4{
			{3, 4, 1, 0},
			{8, 3, 2.5, 7},
			{-2, 2, 1.3, 0},
			{1, 5.6, 2, -4},
		}, V3Zero, Mat4{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 5.6, 2, -4},
		}},
	}
	for _, tt := range tests {
		var m Mat4
		mp := m.Scale(&tt.a, tt.v)
		if !tt.want.nearEq(&m) {
			t.Errorf("m.Scale(%v, %s) = %v, want %v", tt.a, tt.v, m, tt.want)
		}
		if mp != &m {
			t.Errorf("m.Scale(...) does not return the pointer to m")
		}
	}
}

func TestMat4T(t *testing.T) {
	tests := []struct {
		a, want Mat4
	}{
		{Mat4{
			{11, 12, 13, 14},
			{21, 22, 23, 24},
			{31, 32, 33, 34},
			{41, 42, 43, 44},
		}, Mat4{
			{11, 21, 31, 41},
			{12, 22, 32, 42},
			{13, 23, 33, 43},
			{14, 24, 34, 44},
		}},
	}
	for _, tt := range tests {
		var m Mat4
		mp := m.T(&tt.a)
		if tt.want != m {
			t.Errorf("m.T(%v) = %v, want %v", tt.a, m, tt.want)
		}
		if mp != &m {
			t.Errorf("m.T(...) does not return the pointer to m")
		}
	}
}

func TestMat4Translate(t *testing.T) {
	tests := []struct {
		a    Mat4
		v    Vec3
		want Mat4
	}{
		{Mat4{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{3, 9, 4, 1},
		}, V3(2, 0.5, -1), Mat4{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{5, 9.5, 3, 1},
		}},
	}
	for _, tt := range tests {
		var m Mat4
		mp := m.Translate(&tt.a, tt.v)
		if !tt.want.nearEq(&m) {
			t.Errorf("m.Translate(%v, %s) = %v, want %v", tt.a, tt.v, m, tt.want)
		}
		if mp != &m {
			t.Errorf("m.Translate(...) does not return the pointer to m")
		}
	}
}
