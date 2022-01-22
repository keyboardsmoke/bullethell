package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationFrame struct {
	FrameTime float64
	FrameNum  uint
	Start     func()
	Complete  func()
}

type Animation struct {
	Sprite           *Sprite
	Frames           []AnimationFrame
	CurrentFrame     uint
	CurrentFrameTime float64
}

func (a *Animation) GetCurrentFrame() *AnimationFrame {
	return &a.Frames[a.CurrentFrame]
}

func (a *Animation) Update(deltaTime float64) {
	if len(a.Frames) == 0 {
		return // no frames to update
	}

	if (a.CurrentFrameTime + deltaTime) >= a.GetCurrentFrame().FrameTime {
		if a.GetCurrentFrame().Complete != nil {
			a.GetCurrentFrame().Complete()
		}

		a.CurrentFrame = (a.CurrentFrame + 1) % uint(len(a.Frames))
		a.CurrentFrameTime = 0

		if a.GetCurrentFrame().Start != nil {
			a.GetCurrentFrame().Start()
		}
	} else {
		a.CurrentFrameTime += deltaTime
	}
}

func (a *Animation) Reset() {
	a.CurrentFrame = 0
	a.CurrentFrameTime = 0
}

func (a *Animation) AddFrame(frameNumber uint, frameTime float64, startFn func(), completeFn func()) {
	a.Frames = append(a.Frames, AnimationFrame{
		FrameTime: frameTime,
		FrameNum:  frameNumber,
		Start:     startFn,
		Complete:  completeFn,
	})
}

func (a *Animation) Draw(dst *ebiten.Image, x, y int) {
	a.Sprite.DrawFrame(dst, x, y, a.GetCurrentFrame().FrameNum)
}

func LoadAnimation(filename string, cellSize uint) (*Animation, error) {
	sprite, err := LoadSprite(filename, cellSize)
	if err != nil {
		return nil, err
	}

	return &Animation{
		Sprite:           sprite,
		Frames:           []AnimationFrame{},
		CurrentFrame:     0,
		CurrentFrameTime: 0,
	}, nil
}
