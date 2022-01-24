package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Frame struct {
	Name          string
	Area          Box
	OnDraw        func(f *Frame) error
	OnKeyDown     func(target *Frame, key ebiten.Key) error
	OnKeyUp       func(target *Frame, key ebiten.Key) error
	OnResize      func(target *Frame, width, height int) error
	OnMouseDown   func(target *Frame, p Vec2, button ebiten.MouseButton) error
	OnMouseUp     func(target *Frame, p Vec2, button ebiten.MouseButton) error
	OnMouseMove   func(target *Frame, p Vec2) error
	OnMouseScroll func(target *Frame, delta float64) error
	OnMouseEnter  func(target *Frame, p Vec2) error
	OnMouseLeave  func(target *Frame, p Vec2) error
	SubFrames     []*Frame
}

var rootFrame Frame = Frame{
	Name:          "root",
	Area:          Box{Vec2{0, 0}, Vec2{0, 0}},
	OnDraw:        nil,
	OnKeyDown:     nil,
	OnKeyUp:       nil,
	OnResize:      nil,
	OnMouseDown:   nil,
	OnMouseUp:     nil,
	OnMouseMove:   nil,
	OnMouseScroll: nil,
	OnMouseEnter:  nil,
	OnMouseLeave:  nil,
	SubFrames:     []*Frame{},
}

func GetRootFrame() *Frame {
	return &rootFrame
}

func (f *Frame) AddSubFrame(frame *Frame) {
	f.SubFrames = append(f.SubFrames, frame)
}

func (f *Frame) Contains(p Vec2) bool {
	return f.Area.Contains(p)
}

func (f *Frame) ContainsFrame(f2 *Frame) bool {
	return f.Area.ContainsBox(f2.Area)
}

func (f *Frame) Update(deltaTime float64) {
	//
}

func (f *Frame) Draw() error {
	// We actually have to draw the frame children in reverse order before rendering us last
	// That is so that the parent is drawn on top of the children
	for i := len(f.SubFrames) - 1; i >= 0; i-- {
		if err := f.SubFrames[i].Draw(); err != nil {
			return err
		}
	}

	if f.OnDraw != nil {
		err := f.OnDraw(f)
		if err != nil {
			return err
		}
	}

	return nil
}
