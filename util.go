// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
	"strconv"
)

// lerp returns the linear interpolation between a and b by amount t.
// The amount t is usually a value between 0 and 1. If t=0 a will be returned;
// if t=1 b will be returned.
func lerp(a, b, t float32) float32 {
	return a + (b-a)*t
}

const epsilon = 1e-5

// nearEq compares two floating-point numbers for equality within an
// absolute difference tolerance of epsilon.
// This relation is not transitive, except for ε=0.
func nearEq(a, b, ε float32) bool {
	return float32(math.Abs(float64(a-b))) <= ε
}

// str converts a float32 to a string in "%g" format.
func str(f float32) string {
	return strconv.FormatFloat(float64(f), 'g', -1, 32)
}

// x is the radians<->degrees conversion factor.
const x = math.Pi / 180

// Deg converts the measurement of an angle from radians to degrees.
func Deg(rad float32) float32 {
	return rad / x
}

// Rad converts the measurement of an angle from degrees to radians.
func Rad(deg float32) float32 {
	return deg * x
}
