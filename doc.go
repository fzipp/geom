// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package geom provides vector and matrix types suitable for OpenGL
programming: Vec2, Vec3 and Mat4.

	// A 4x4 zero matrix
	var a Mat4

	// Another 4x4 matrix
	b := Mat4{
		{0, 1, 2.3, 3},
		{4, 0.5, 6, 7},
		{8, -9, 10, 11},
		{12, 13, 14, 15},
	}

	// Copy elements from b to a
	a = b

	// Set a matrix element
	a[2][3] = 5

	// Multiply a and b, store the result in a.
	a.Mul(&a, &b)
*/
package geom
