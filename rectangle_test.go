package geom

import "testing"

func TestRect(t *testing.T) {
	tests := []struct {
		x0, y0, x1, y1 float32
		want           Rectangle
	}{
		{
			0, 0, 0, 0,
			Rectangle{V2(0, 0), V2(0, 0)},
		},
		{
			0, 0, 320, 200,
			Rectangle{V2(0, 0), V2(320, 200)},
		},
		{
			25.2, -7.6, 102.3, 98,
			Rectangle{V2(25.2, -7.6), V2(102.3, 98)},
		},
	}
	for _, tt := range tests {
		rect := Rect(tt.x0, tt.y0, tt.x1, tt.y1)
		if !rectangleNearEq(rect, tt.want) {
			t.Errorf("Rect(%g, %g, %g, %g) was: %v, want: %v",
				tt.x0, tt.y0, tt.x1, tt.y1, rect, tt.want)
		}
	}
}

func TestRectSized(t *testing.T) {
	tests := []struct {
		pos  Vec2
		size Size
		want Rectangle
	}{
		{
			V2(0, 0),
			Size{W: 20, H: 30},
			Rectangle{V2(0, 0), V2(20, 30)},
		},
		{
			V2(12, 25),
			Size{W: 20, H: 30},
			Rectangle{V2(12, 25), V2(32, 55)},
		},
		{
			V2(-3.4, 5.2),
			Size{W: 4.15, H: -0.3},
			Rectangle{V2(-3.4, 5.2), V2(0.75, 4.9)},
		},
	}
	for _, tt := range tests {
		if rect := RectSized(tt.pos, tt.size); !rectangleNearEq(rect, tt.want) {
			t.Errorf("RectSized(%v, %+v) was: %v, want: %v",
				tt.pos, tt.size, rect, tt.want)
		}
	}
}

func TestRectangleSize(t *testing.T) {
	tests := []struct {
		rect Rectangle
		want Size
	}{
		{Rectangle{V2(0, 0), V2(0, 0)}, Size{W: 0, H: 0}},
		{Rectangle{V2(0, 0), V2(15, 28)}, Size{W: 15, H: 28}},
		{Rectangle{V2(2.5, 4), V2(5.2, 6.1)}, Size{W: 2.7, H: 2.1}},
		{Rectangle{V2(-250, -320), V2(110, 230)}, Size{W: 360, H: 550}},
	}
	for _, tt := range tests {
		if size := tt.rect.Size(); !sizeNearEq(size, tt.want) {
			t.Errorf("%v.Size() was: %v, want: %v", tt.rect, size, tt.want)
		}
	}
}

func rectangleNearEq(a, b Rectangle) bool {
	return a.Min.NearEq(b.Min) && a.Max.NearEq(b.Max)
}

func sizeNearEq(a, b Size) bool {
	return nearEq(a.W, b.W, epsilon) && nearEq(a.H, b.H, epsilon)
}
