package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Mouse struct {
	Position    Vec2
	State       map[ebiten.MouseButton]bool
	onMouseDown []func(p Vec2, button ebiten.MouseButton) error
	onMouseUp   []func(p Vec2, button ebiten.MouseButton) error
	onMouseMove []func(p Vec2) error
}

type Keyboard struct {
	PressedKeys []ebiten.Key
	onKeyDown   []func(key ebiten.Key) error
	onKeyUp     []func(key ebiten.Key) error
}

type InputManager struct {
	Mouse    *Mouse
	Keyboard *Keyboard
}

var (
	manager *InputManager
)

func GetInputManager() *InputManager {
	if manager == nil {
		manager = &InputManager{
			Mouse:    &Mouse{State: make(map[ebiten.MouseButton]bool)},
			Keyboard: &Keyboard{},
		}
	}

	return manager
}

func (m *Mouse) AddMouseDownCallback(f func(p Vec2, button ebiten.MouseButton) error) {
	m.onMouseDown = append(m.onMouseDown, f)
}

func (m *Mouse) AddMouseUpCallback(f func(p Vec2, button ebiten.MouseButton) error) {
	m.onMouseUp = append(m.onMouseUp, f)
}

func (m *Mouse) AddMouseMoveCallback(f func(p Vec2) error) {
	m.onMouseMove = append(m.onMouseMove, f)
}

func (k *Keyboard) AddKeyDownCallback(f func(key ebiten.Key) error) {
	k.onKeyDown = append(k.onKeyDown, f)
}

func (k *Keyboard) AddKeyUpCallback(f func(key ebiten.Key) error) {
	k.onKeyUp = append(k.onKeyUp, f)
}

func (i *InputManager) handleMouseButton(button ebiten.MouseButton) {
	if inpututil.IsMouseButtonJustPressed(button) {
		for _, f := range i.Mouse.onMouseDown {
			f(i.Mouse.Position, button)
		}
	}

	if inpututil.IsMouseButtonJustReleased(button) {
		for _, f := range i.Mouse.onMouseUp {
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
		for _, f := range i.Mouse.onMouseMove {
			f(i.Mouse.Position)
		}
	}

	// Handle released keys
	for _, key := range i.Keyboard.PressedKeys {
		if inpututil.IsKeyJustReleased(key) {
			for _, f := range i.Keyboard.onKeyUp {
				f(key)
			}
		}
	}

	// Is this correct?...
	i.Keyboard.PressedKeys = inpututil.AppendPressedKeys(i.Keyboard.PressedKeys[:0])

	// Handle newly pressed keys
	for _, key := range i.Keyboard.PressedKeys {
		if inpututil.IsKeyJustPressed(key) {
			for _, f := range i.Keyboard.onKeyDown {
				f(key)
			}
		}
	}
}
