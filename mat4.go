// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
	"unsafe"
)

// A Mat4 represents a 4x4 matrix. The indices are [row][column].
type Mat4 [4][4]float32

// id is the 4x4 identity matrix.
var id = Mat4{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
}

// zero is the 4x4 zero matrix.
var zero Mat4

// Id sets m to the identity matrix and returns m.
func (m *Mat4) Id() *Mat4 {
	*m = id
	return m
}

// Zero sets all elements of m to 0 (zero matrix) and returns m.
func (m *Mat4) Zero() *Mat4 {
	*m = zero
	return m
}

// Det calculates the determinant of 4x4 matrix m.
func (m *Mat4) Det() float32 {
	return m[0][3]*m[1][2]*m[2][1]*m[3][0] - m[0][2]*m[1][3]*m[2][1]*m[3][0] -
		m[0][3]*m[1][1]*m[2][2]*m[3][0] + m[0][1]*m[1][3]*m[2][2]*m[3][0] +
		m[0][2]*m[1][1]*m[2][3]*m[3][0] - m[0][1]*m[1][2]*m[2][3]*m[3][0] -
		m[0][3]*m[1][2]*m[2][0]*m[3][1] + m[0][2]*m[1][3]*m[2][0]*m[3][1] +
		m[0][3]*m[1][0]*m[2][2]*m[3][1] - m[0][0]*m[1][3]*m[2][2]*m[3][1] -
		m[0][2]*m[1][0]*m[2][3]*m[3][1] + m[0][0]*m[1][2]*m[2][3]*m[3][1] +
		m[0][3]*m[1][1]*m[2][0]*m[3][2] - m[0][1]*m[1][3]*m[2][0]*m[3][2] -
		m[0][3]*m[1][0]*m[2][1]*m[3][2] + m[0][0]*m[1][3]*m[2][1]*m[3][2] +
		m[0][1]*m[1][0]*m[2][3]*m[3][2] - m[0][0]*m[1][1]*m[2][3]*m[3][2] -
		m[0][2]*m[1][1]*m[2][0]*m[3][3] + m[0][1]*m[1][2]*m[2][0]*m[3][3] +
		m[0][2]*m[1][0]*m[2][1]*m[3][3] - m[0][0]*m[1][2]*m[2][1]*m[3][3] -
		m[0][1]*m[1][0]*m[2][2]*m[3][3] + m[0][0]*m[1][1]*m[2][2]*m[3][3]
}

// Mul sets m to the matrix product a*b and returns m.
func (m *Mat4) Mul(a *Mat4, b *Mat4) *Mat4 {
	*m = Mat4{
		{
			a[0][0]*b[0][0] + a[1][0]*b[0][1] + a[2][0]*b[0][2] + a[3][0]*b[0][3],
			a[0][1]*b[0][0] + a[1][1]*b[0][1] + a[2][1]*b[0][2] + a[3][1]*b[0][3],
			a[0][2]*b[0][0] + a[1][2]*b[0][1] + a[2][2]*b[0][2] + a[3][2]*b[0][3],
			a[0][3]*b[0][0] + a[1][3]*b[0][1] + a[2][3]*b[0][2] + a[3][3]*b[0][3],
		},
		{
			a[0][0]*b[1][0] + a[1][0]*b[1][1] + a[2][0]*b[1][2] + a[3][0]*b[1][3],
			a[0][1]*b[1][0] + a[1][1]*b[1][1] + a[2][1]*b[1][2] + a[3][1]*b[1][3],
			a[0][2]*b[1][0] + a[1][2]*b[1][1] + a[2][2]*b[1][2] + a[3][2]*b[1][3],
			a[0][3]*b[1][0] + a[1][3]*b[1][1] + a[2][3]*b[1][2] + a[3][3]*b[1][3],
		},
		{
			a[0][0]*b[2][0] + a[1][0]*b[2][1] + a[2][0]*b[2][2] + a[3][0]*b[2][3],
			a[0][1]*b[2][0] + a[1][1]*b[2][1] + a[2][1]*b[2][2] + a[3][1]*b[2][3],
			a[0][2]*b[2][0] + a[1][2]*b[2][1] + a[2][2]*b[2][2] + a[3][2]*b[2][3],
			a[0][3]*b[2][0] + a[1][3]*b[2][1] + a[2][3]*b[2][2] + a[3][3]*b[2][3],
		},
		{
			a[0][0]*b[3][0] + a[1][0]*b[3][1] + a[2][0]*b[3][2] + a[3][0]*b[3][3],
			a[0][1]*b[3][0] + a[1][1]*b[3][1] + a[2][1]*b[3][2] + a[3][1]*b[3][3],
			a[0][2]*b[3][0] + a[1][2]*b[3][1] + a[2][2]*b[3][2] + a[3][2]*b[3][3],
			a[0][3]*b[3][0] + a[1][3]*b[3][1] + a[2][3]*b[3][2] + a[3][3]*b[3][3],
		},
	}
	return m
}

// Ortho sets m to an orthographic projection matrix with the given clipping
// planes and returns m.
func (m *Mat4) Ortho(left, right, bottom, top, near, far float32) *Mat4 {
	dx := left - right
	dy := bottom - top
	dz := near - far
	*m = Mat4{
		{-2 / dx, 0, 0, 0},
		{0, -2 / dy, 0, 0},
		{0, 0, 2 / dz, 0},
		{(left + right) / dx, (top + bottom) / dy, (far + near) / dz, 1},
	}
	return m
}

// Frustum sets m to a frustum matrix with the given clipping planes and
// returns m.
func (m *Mat4) Frustum(left, right, bottom, top, near, far float32) *Mat4 {
	dx := right - left
	dy := top - bottom
	dz := near - far
	*m = Mat4{
		{(2 * near) / dx, 0, 0, 0},
		{0, (2 * near) / dy, 0, 0},
		{(left + right) / dx, (top + bottom) / dy, (far + near) / dz, -1},
		{0, 0, (2 * far * near) / dz, 0},
	}
	return m
}

// Perspective sets m to a perspective matrix with the given vertical field of
// view angle (in radians), aspect ratio, near and far bounds of the frustum,
// and returns m.
func (m *Mat4) Perspective(fovy, aspect, near, far float32) *Mat4 {
	f := 1 / float32(math.Tan(float64(fovy/2)))
	dz := near - far
	*m = Mat4{
		{f / aspect, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, (far + near) / dz, -1},
		{0, 0, (2 * far * near) / dz, 0},
	}
	return m
}

// LookAt sets m to a viewing matrix given an eye point, a reference point
// indicating the center of the scene and an up vector, and returns m.
func (m *Mat4) LookAt(eye, center, up Vec3) *Mat4 {
	vz := eye.Sub(center).Norm()
	vx := up.Cross(vz).Norm()
	vy := vz.Cross(vx)
	*m = Mat4{
		{vx.X, vy.X, vz.X, 0},
		{vx.Y, vy.Y, vz.Y, 0},
		{vx.Z, vy.Z, vz.Z, 0},
		{-vx.Dot(eye), -vy.Dot(eye), -vz.Dot(eye), 1},
	}
	return m
}

// Rot sets m to the rotation of matrix a by the given angle in radians around
// the given axis, and returns m.
func (m *Mat4) Rot(a *Mat4, angle float32, axis Vec3) *Mat4 {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	t := 1 - c
	n := axis.Norm()
	b := Mat4{
		{n.X*n.X*t + c, n.Y*n.X*t + n.Z*s, n.Z*n.X*t - n.Y*s, 0},
		{n.X*n.Y*t - n.Z*s, n.Y*n.Y*t + c, n.Z*n.Y*t + n.X*s, 0},
		{n.X*n.Z*t + n.Y*s, n.Y*n.Z*t - n.X*s, n.Z*n.Z*t + c, 0},
		{0, 0, 0, 1},
	}
	return m.Mul(a, &b)
}

// T sets m to the transpose of matrix a and returns m.
func (m *Mat4) T(a *Mat4) *Mat4 {
	*m = Mat4{
		{a[0][0], a[1][0], a[2][0], a[3][0]},
		{a[0][1], a[1][1], a[2][1], a[3][1]},
		{a[0][2], a[1][2], a[2][2], a[3][2]},
		{a[0][3], a[1][3], a[2][3], a[3][3]},
	}
	return m
}

// Scale sets m to the scaling of matrix a by the scale factors of v and
// returns m.
func (m *Mat4) Scale(a *Mat4, v Vec3) *Mat4 {
	*m = Mat4{
		{a[0][0] * v.X, a[0][1] * v.X, a[0][2] * v.X, a[0][3] * v.X},
		{a[1][0] * v.Y, a[1][1] * v.Y, a[1][2] * v.Y, a[1][3] * v.Y},
		{a[2][0] * v.Z, a[2][1] * v.Z, a[2][2] * v.Z, a[2][3] * v.Z},
		{a[3][0], a[3][1], a[3][2], a[3][3]},
	}
	return m
}

// Translate sets m to the translation of matrix a by the vector v and
// returns m.
func (m *Mat4) Translate(a *Mat4, v Vec3) *Mat4 {
	*m = Mat4{
		{a[0][0], a[0][1], a[0][2], a[0][3]},
		{a[1][0], a[1][1], a[1][2], a[1][3]},
		{a[2][0], a[2][1], a[2][2], a[2][3]},
		{
			a[0][0]*v.X + a[1][0]*v.Y + a[2][0]*v.Z + a[3][0],
			a[0][1]*v.X + a[1][1]*v.Y + a[2][1]*v.Z + a[3][1],
			a[0][2]*v.X + a[1][2]*v.Y + a[2][2]*v.Z + a[3][2],
			a[0][3]*v.X + a[1][3]*v.Y + a[2][3]*v.Z + a[3][3],
		},
	}
	return m
}

// Floats returns a pointer to the matrix elements represented as a flat
// array of float32 numbers in row-major order. Changing an element value
// of this array will affect m and vice versa.
func (m *Mat4) Floats() *[16]float32 {
	return (*[16]float32)(unsafe.Pointer(m))
}

// nearEq returns whether m and m2 are approximately equal. This relation is
// not transitive in general. The tolerance for the floating-point components
// is ±1e-5.
func (m *Mat4) nearEq(m2 *Mat4) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if !nearEq(m[i][j], m2[i][j], epsilon) {
				return false
			}
		}
	}
	return true
}
