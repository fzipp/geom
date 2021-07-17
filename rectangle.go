// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

// A Rectangle contains the points with Min.X <= X <= Max.X, Min.Y <= Y <= Max.Y.
// It is well-formed if Min.X <= Max.X and likewise for Y.
type Rectangle struct {
	Min Vec2
	Max Vec2
}

// Rect is shorthand for Rectangle{Min: geom.V2(x0, y0), Max: geom.V2(x1, y1)}.
func Rect(x0, y0, x1, y1 float32) Rectangle {
	return Rectangle{
		Min: Vec2{X: x0, Y: y0},
		Max: Vec2{X: x1, Y: y1},
	}
}

// RectSized creates a rectangle based on its position and size.
func RectSized(pos Vec2, size Size) Rectangle {
	return Rectangle{Min: pos, Max: Vec2{X: pos.X + size.W, Y: pos.Y + size.H}}
}

// Contains reports whether the rectangle contains point pt.
func (r Rectangle) Contains(pt Vec2) bool {
	return (pt.X >= r.Min.X) && (pt.X <= r.Max.X) &&
		(pt.Y >= r.Min.Y) && (pt.Y <= r.Max.Y)
}

// Size returns the dimensions (width and height) of the rectangle.
func (r *Rectangle) Size() Size {
	return Size{W: r.Max.X - r.Min.X, H: r.Max.Y - r.Min.Y}
}
