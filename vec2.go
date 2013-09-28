// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
)

// A Vec2 represents a vector with coordinates X and Y in 2-dimensional
// euclidian space.
type Vec2 struct {
	X, Y float32
}

// A Size represents the dimensions of a rectangle.
type Size struct {
	// Width and height
	W, H float32
}

var (
	// V2Zero is the zero vector (0,0).
	V2Zero = Vec2{0, 0}
	// V2Unit is the unit vector (1,1).
	V2Unit = Vec2{1, 1}
	// V2UnitX is the x-axis unit vector (1,0).
	V2UnitX = Vec2{1, 0}
	// V2UnitY is the y-axis unit vector (0,1).
	V2UnitY = Vec2{0, 1}
)

// V2 is shorthand for Vec2{X: x, Y: y}.
func V2(x, y float32) Vec2 {
	return Vec2{x, y}
}

// Add returns the vector v+w.
func (v Vec2) Add(w Vec2) Vec2 {
	return Vec2{v.X + w.X, v.Y + w.Y}
}

// Sub returns the vector v-w.
func (v Vec2) Sub(w Vec2) Vec2 {
	return Vec2{v.X - w.X, v.Y - w.Y}
}

// Mul returns the vector v*s.
func (v Vec2) Mul(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// Div returns the vector v/s.
func (v Vec2) Div(s float32) Vec2 {
	return Vec2{v.X / s, v.Y / s}
}

// Neg returns the negated vector of v.
func (v Vec2) Neg() Vec2 {
	return v.Mul(-1)
}

// Dot returns the dot (a.k.a. scalar) product of v and w.
func (v Vec2) Dot(w Vec2) float32 {
	return v.X*w.X + v.Y*w.Y
}

// CrossLen returns the length that the cross product of v and w would have
// in 3-dimensional euclidian space. This is effectively the Z component
// of the 3D cross product vector.
func (v Vec2) CrossLen(w Vec2) float32 {
	return v.X*w.Y - v.Y*w.X
}

// CompMul returns the component-wise multiplication of two vectors.
func (v Vec2) CompMul(w Vec2) Vec2 {
	return Vec2{v.X * w.X, v.Y * w.Y}
}

// CompDiv returns the component-wise division of two vectors.
func (v Vec2) CompDiv(w Vec2) Vec2 {
	return Vec2{v.X / w.X, v.Y / w.Y}
}

// SqDist returns the square of the euclidian distance between two vectors.
func (v Vec2) SqDist(w Vec2) float32 {
	return v.Sub(w).SqLen()
}

// Dist returns the euclidian distance between two vectors.
func (v Vec2) Dist(w Vec2) float32 {
	return v.Sub(w).Len()
}

// SqLen returns the square of the length (euclidian norm) of a vector.
func (v Vec2) SqLen() float32 {
	return v.Dot(v)
}

// Len returns the length (euclidian norm) of a vector.
func (v Vec2) Len() float32 {
	return float32(math.Sqrt(float64(v.SqLen())))
}

// Norm returns the normalized vector of a vector.
func (v Vec2) Norm() Vec2 {
	return v.Div(v.Len())
}

// Reflect returns the reflection vector of v given a normal n.
func (v Vec2) Reflect(n Vec2) Vec2 {
	return v.Sub(n.Mul(2 * v.Dot(n)))
}

// Lerp returns the linear interpolation between v and w by amount t.
// The amount t is usually a value between 0 and 1. If t=0 v will be
// returned; if t=1 w will be returned.
func (v Vec2) Lerp(w Vec2, t float32) Vec2 {
	// return v.Add(w.Sub(v).Mul(t))
	return Vec2{lerp(v.X, w.X, t), lerp(v.Y, w.Y, t)}
}

// Angle returns the angle (counterclockwise) of vector v with the x axis in
// radians. The result is in the interval [0,2π).
func (v Vec2) Angle() float32 {
	a := math.Atan2(float64(v.Y), float64(v.X))
	if a < 0 {
		a += 2 * math.Pi
	}
	return float32(a)
}

// Z returns a Vec3 based on v with the additional coordinate z.
func (v Vec2) Z(z float32) Vec3 {
	return Vec3{v.X, v.Y, z}
}

// NearEq returns whether v and w are approximately equal. This relation is not
// transitive in general. The tolerance for the floating-point components is
// ±1e-5.
func (v Vec2) NearEq(w Vec2) bool {
	return nearEq(v.X, w.X, epsilon) && nearEq(v.Y, w.Y, epsilon)
}

// String returns a string representation of v like "(3.25, -1.5)".
func (v Vec2) String() string {
	return "(" + str(v.X) + ", " + str(v.Y) + ")"
}
