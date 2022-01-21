package engine

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v *Vec2) Add(other Vec2) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vec2) Subtract(other Vec2) {
	v.X -= other.X
	v.Y -= other.Y
}

func (v *Vec2) Multiply(other float64) {
	v.X *= other
	v.Y *= other
}

func (v Vec2) Dot(other Vec2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec2) Cross(other Vec2) float64 {
	return v.X*other.Y - v.Y*other.X
}

func (v Vec2) CrossF(other float64) Vec2 {
	return Vec2{-v.Y * other, v.X * other}
}

func (v Vec2) LengthSq() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.LengthSq())
}

func (v *Vec2) Normalize() {
	v.Multiply(1.0 / v.Length())
}

func (v Vec2) Normalized() Vec2 {
	return Multiply(v, 1/v.Length())
}

func Add(v, u Vec2) Vec2 {
	return Vec2{v.X + u.X, v.Y + u.Y}
}

func Subtract(v, u Vec2) Vec2 {
	return Vec2{v.X - u.X, v.Y - u.Y}
}

func Multiply(v Vec2, r float64) Vec2 {
	return Vec2{v.X * r, v.Y * r}
}
