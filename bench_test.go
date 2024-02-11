// Copyright 2015 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import "testing"

var a = &Mat4{
	{1, 2, 4, 4},
	{5, 6, 7, 8},
	{1, 2, 3, 4},
	{5, 6, 7, 8},
}

func BenchmarkMat4ID(b *testing.B) {
	var m Mat4
	for range b.N {
		m.ID()
	}
}

func BenchmarkMat4Zero(b *testing.B) {
	var m Mat4
	for range b.N {
		m.Zero()
	}
}

func BenchmarkMat4Det(b *testing.B) {
	for range b.N {
		a.Det()
	}
}

func BenchmarkMat4Mul(b *testing.B) {
	var m Mat4
	for range b.N {
		m.Mul(a, a)
	}
}

func BenchmarkMat4Ortho(b *testing.B) {
	var m Mat4
	for range b.N {
		m.Ortho(1, 2, 3, 4, 5, 6)
	}
}

func BenchmarkMat4Frustum(b *testing.B) {
	var m Mat4
	for range b.N {
		m.Frustum(1, 2, 3, 4, 5, 6)
	}
}

func BenchmarkMat4Perspective(b *testing.B) {
	var m Mat4
	for range b.N {
		m.Perspective(1, 2, 3, 4)
	}
}

func BenchmarkMat4LookAt(b *testing.B) {
	var m Mat4
	eye := V3(1, 2, 3)
	center := V3(4, 5, 6)
	up := V3(7, 8, 9)
	for range b.N {
		m.LookAt(eye, center, up)
	}
}

func BenchmarkMat4Rot(b *testing.B) {
	var m Mat4
	axis := V3(1, 2, 3)
	for range b.N {
		m.Rot(a, 0.5, axis)
	}
}

func BenchmarkMat4T(b *testing.B) {
	var m Mat4
	for range b.N {
		m.T(a)
	}
}

func BenchmarkMat4Scale(b *testing.B) {
	var m Mat4
	v := V3(1, 2, 3)
	for range b.N {
		m.Scale(a, v)
	}
}

func BenchmarkMat4Translate(b *testing.B) {
	var m Mat4
	v := V3(1, 2, 3)
	for range b.N {
		m.Translate(a, v)
	}
}

func BenchmarkMat4Floats(b *testing.B) {
	var r *[16]float32
	for range b.N {
		r = a.Floats()
	}
	_ = r
}

func BenchmarkVec2Add(b *testing.B) {
	var r Vec2
	v := V2(1, 2)
	w := V2(3, 4)
	for range b.N {
		r = v.Add(w)
	}
	_ = r
}

func BenchmarkVec2Sub(b *testing.B) {
	var r Vec2
	v := V2(1, 2)
	w := V2(3, 4)
	for range b.N {
		r = v.Sub(w)
	}
	_ = r
}

func BenchmarkVec2Mul(b *testing.B) {
	var r Vec2
	v := V2(1, 2)
	for range b.N {
		r = v.Mul(2.5)
	}
	_ = r
}

func BenchmarkVec2Div(b *testing.B) {
	var r Vec2
	v := V2(1, 2)
	for range b.N {
		r = v.Div(2.5)
	}
	_ = r
}

func BenchmarkVec2Neg(b *testing.B) {
	var r Vec2
	v := V2(1, 2)
	for range b.N {
		r = v.Neg()
	}
	_ = r
}

func BenchmarkVec2Dot(b *testing.B) {
	var r float32
	v := V2(1, 2)
	w := V2(4, 5)
	for range b.N {
		r = v.Dot(w)
	}
	_ = r
}

func BenchmarkVec2CrossLen(b *testing.B) {
	var r float32
	v := V2(1, 2)
	w := V2(4, 5)
	for range b.N {
		r = v.CrossLen(w)
	}
	_ = r
}

func BenchmarkVec2CompMul(b *testing.B) {
	var r Vec2
	v := V2(1, 2)
	w := V2(4, 5)
	for range b.N {
		r = v.CompMul(w)
	}
	_ = r
}

func BenchmarkVec2CompDiv(b *testing.B) {
	var r Vec2
	v := V2(1, 2)
	w := V2(4, 5)
	for range b.N {
		r = v.CompDiv(w)
	}
	_ = r
}

func BenchmarkVec2SqDist(b *testing.B) {
	var r float32
	v := V2(1, 2)
	w := V2(4, 5)
	for range b.N {
		r = v.SqDist(w)
	}
	_ = r
}

func BenchmarkVec2Dist(b *testing.B) {
	var r float32
	v := V2(1, 2)
	w := V2(4, 5)
	for range b.N {
		r = v.Dist(w)
	}
	_ = r
}

func BenchmarkVec2SqLen(b *testing.B) {
	var r float32
	v := V2(1, 2)
	for range b.N {
		r = v.SqLen()
	}
	_ = r
}

func BenchmarkVec2Len(b *testing.B) {
	var r float32
	v := V2(1, 2)
	for range b.N {
		r = v.Len()
	}
	_ = r
}

func BenchmarkVec2Norm(b *testing.B) {
	var r Vec2
	v := V2(1, 2)
	for range b.N {
		r = v.Norm()
	}
	_ = r
}

func BenchmarkVec2Reflect(b *testing.B) {
	var r Vec2
	v := V2(1, 2)
	normal := V2(2, 3)
	for range b.N {
		r = v.Reflect(normal)
	}
	_ = r
}

func BenchmarkVec2Lerp(b *testing.B) {
	var r Vec2
	v := V2(1, 2)
	w := V2(2, 3)
	for range b.N {
		r = v.Lerp(w, 0.6)
	}
	_ = r
}

func BenchmarkVec2Min(b *testing.B) {
	var r Vec2
	v := V2(1, 2)
	w := V2(2, 3)
	for range b.N {
		r = v.Min(w)
	}
	_ = r
}

func BenchmarkVec2Max(b *testing.B) {
	var r Vec2
	v := V2(1, 2)
	w := V2(2, 3)
	for range b.N {
		r = v.Max(w)
	}
	_ = r
}

func BenchmarkVec2Transform(b *testing.B) {
	var r Vec2
	v := V2(1, 2)
	for range b.N {
		r = v.Transform(a)
	}
	_ = r
}

func BenchmarkVec3Add(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	w := V3(4, 5, 6)
	for range b.N {
		r = v.Add(w)
	}
	_ = r
}

func BenchmarkVec3Sub(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	w := V3(4, 5, 6)
	for range b.N {
		r = v.Sub(w)
	}
	_ = r
}

func BenchmarkVec3Mul(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	for range b.N {
		r = v.Mul(2.5)
	}
	_ = r
}

func BenchmarkVec3Div(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	for range b.N {
		r = v.Div(2.5)
	}
	_ = r
}

func BenchmarkVec3Neg(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	for range b.N {
		r = v.Neg()
	}
	_ = r
}

func BenchmarkVec3Dot(b *testing.B) {
	var r float32
	v := V3(1, 2, 3)
	w := V3(4, 5, 6)
	for range b.N {
		r = v.Dot(w)
	}
	_ = r
}

func BenchmarkVec3Cross(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	w := V3(4, 5, 6)
	for range b.N {
		r = v.Cross(w)
	}
	_ = r
}

func BenchmarkVec3CompMul(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	w := V3(4, 5, 6)
	for range b.N {
		r = v.CompMul(w)
	}
	_ = r
}

func BenchmarkVec3CompDiv(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	w := V3(4, 5, 6)
	for range b.N {
		r = v.CompDiv(w)
	}
	_ = r
}

func BenchmarkVec3SqDist(b *testing.B) {
	var r float32
	v := V3(1, 2, 3)
	w := V3(4, 5, 6)
	for range b.N {
		r = v.SqDist(w)
	}
	_ = r
}

func BenchmarkVec3Dist(b *testing.B) {
	var r float32
	v := V3(1, 2, 3)
	w := V3(4, 5, 6)
	for range b.N {
		r = v.Dist(w)
	}
	_ = r
}

func BenchmarkVec3SqLen(b *testing.B) {
	var r float32
	v := V3(1, 2, 3)
	for range b.N {
		r = v.SqLen()
	}
	_ = r
}

func BenchmarkVec3Len(b *testing.B) {
	var r float32
	v := V3(1, 2, 3)
	for range b.N {
		r = v.Len()
	}
	_ = r
}

func BenchmarkVec3Norm(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	for range b.N {
		r = v.Norm()
	}
	_ = r
}

func BenchmarkVec3Reflect(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	normal := V3(2, 3, 4)
	for range b.N {
		r = v.Reflect(normal)
	}
	_ = r
}

func BenchmarkVec3Lerp(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	w := V3(2, 3, 4)
	for range b.N {
		r = v.Lerp(w, 0.6)
	}
	_ = r
}

func BenchmarkVec3Min(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	w := V3(2, 3, 4)
	for range b.N {
		r = v.Min(w)
	}
	_ = r
}

func BenchmarkVec3Max(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	w := V3(2, 3, 4)
	for range b.N {
		r = v.Max(w)
	}
	_ = r
}

func BenchmarkVec3Transform(b *testing.B) {
	var r Vec3
	v := V3(1, 2, 3)
	for range b.N {
		r = v.Transform(a)
	}
	_ = r
}
