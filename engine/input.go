package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Mouse struct {
	Position    Vec2
	State       map[ebiten.MouseButton]bool
	OnMouseDown []func(p Vec2, button ebiten.MouseButton) error
	OnMouseUp   []func(p Vec2, button ebiten.MouseButton) error
	OnMouseMove []func(p Vec2) error
}

type Keyboard struct {
	PressedKeys []ebiten.Key
	OnKeyDown   []func(key ebiten.Key) error
	OnKeyUp     []func(key ebiten.Key) error
}

type InputManager struct {
	Mouse    *Mouse
	Keyboard *Keyboard
}

func (i *InputManager) handleMouseButton(button ebiten.MouseButton) {
	if inpututil.IsMouseButtonJustPressed(button) {
		for _, f := range i.Mouse.OnMouseDown {
			f(i.Mouse.Position, button)
		}
	}

	if inpututil.IsMouseButtonJustReleased(button) {
		for _, f := range i.Mouse.OnMouseUp {
			f(i.Mouse.Position, button)
		}
	}
}

func (i *InputManager) Update(deltaTime float64) {
	mx, my := ebiten.CursorPosition()

	lastPosition := i.Mouse.Position
	i.Mouse.Position = Vec2{float64(mx), float64(my)}

	// Handle click events
	i.handleMouseButton(ebiten.MouseButtonLeft)
	i.handleMouseButton(ebiten.MouseButtonRight)
	i.handleMouseButton(ebiten.MouseButtonMiddle)

	// Only update position on a move
	if lastPosition != i.Mouse.Position {
		for _, f := range i.Mouse.OnMouseMove {
			f(i.Mouse.Position)
		}
	}

	// Handle released keys
	for _, key := range i.Keyboard.PressedKeys {
		if inpututil.IsKeyJustReleased(key) {
			for _, f := range i.Keyboard.OnKeyUp {
				f(key)
			}
		}
	}

	// Is this correct?...
	i.Keyboard.PressedKeys = inpututil.AppendPressedKeys(i.Keyboard.PressedKeys[:0])

	// Handle newly pressed keys
	for _, key := range i.Keyboard.PressedKeys {
		if inpututil.IsKeyJustPressed(key) {
			for _, f := range i.Keyboard.OnKeyDown {
				f(key)
			}
		}
	}
}
