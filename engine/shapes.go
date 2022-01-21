package engine

import "math"

type Line struct {
	Start Vec2
	End   Vec2
}

type Box struct {
	Min Vec2
	Max Vec2
}

type Circle struct {
	Center Vec2
	Radius float64
}

func IntersectLines(l1, l2 *Line) (Vec2, bool) {
	if l1.Start == l2.Start || l1.Start == l2.End || l1.End == l2.Start || l1.End == l2.End {
		return Vec2{}, false
	}

	x1 := l1.Start.X
	y1 := l1.Start.Y
	x2 := l1.End.X
	y2 := l1.End.Y
	x3 := l2.Start.X
	y3 := l2.Start.Y
	x4 := l2.End.X
	y4 := l2.End.Y

	denominator := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)

	if denominator == 0 {
		return Vec2{}, false
	}

	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / denominator
	u := -((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3)) / denominator

	return Vec2{X: x1 + t*(x2-x1), Y: y1 + t*(y2-y1)}, t >= 0 && t <= 1 && u >= 0 && u <= 1
}

func (l *Line) Angle() float64 {
	return math.Atan2(l.End.Y-l.Start.Y, l.End.X-l.Start.X)
}

func (b *Box) Contains(v Vec2) bool {
	return b.Min.X <= v.X && b.Max.X >= v.X && b.Min.Y <= v.Y && b.Max.Y >= v.Y
}

func (b *Box) ContainsLine(l *Line) bool {
	return b.Contains(l.Start) && b.Contains(l.End)
}

func (b *Box) ContainsBox(b2 Box) bool {
	return b.Min.X <= b2.Min.X && b.Max.X >= b2.Max.X && b.Min.Y <= b2.Min.Y && b.Max.Y >= b2.Max.Y
}

func (b *Box) ContainsCircle(c Circle) bool {
	return b.Min.X <= c.Center.X-c.Radius && b.Max.X >= c.Center.X+c.Radius && b.Min.Y <= c.Center.Y-c.Radius && b.Max.Y >= c.Center.Y+c.Radius
}

func (l *Line) Intersects(other *Line) bool {
	_, ok := IntersectLines(l, other)

	return ok
}

func (b *Box) Intersects(v Vec2) bool {
	return b.Min.X <= v.X && b.Max.X >= v.X && b.Min.Y <= v.Y && b.Max.Y >= v.Y
}

func (b *Box) IntersectsLine(l *Line) bool {
	return b.Intersects(l.Start) || b.Intersects(l.End)
}

func (b *Box) IntersectsBox(b2 Box) bool {
	return b.Min.X <= b2.Max.X && b.Max.X >= b2.Min.X && b.Min.Y <= b2.Max.Y && b.Max.Y >= b2.Min.Y
}

func (b *Box) IntersectsCircle(c Circle) bool {
	return b.Min.X <= c.Center.X && b.Max.X >= c.Center.X && b.Min.Y <= c.Center.Y && b.Max.Y >= c.Center.Y
}
