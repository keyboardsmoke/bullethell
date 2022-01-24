package main

import (
	"bh/engine"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"time"

	"github.com/flopp/go-findfont"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	GameScreenWidth  = 1024
	GameScreenHeight = 768
)

type Game struct {
	lastUpdate   time.Time
	cam          engine.Camera
	audioContext *audio.Context
	ScreenWidth  int
	ScreenHeight int
	DebugFont    *engine.Font
	Blob         *engine.Animation
	RootFrame    *engine.Frame
}

func (g *Game) Update() error {
	now := time.Now()
	deltaTime := float64(now.Sub(g.lastUpdate).Microseconds()) / 1000.0

	// Update inputs :)
	engine.GetInputManager().Update(deltaTime)

	// Update each child frame with this
	g.RootFrame.Update(deltaTime)

	// Entity updates
	g.Blob.Update(deltaTime)

	g.lastUpdate = now

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	fps := fmt.Sprintf("FPS: %f", ebiten.CurrentFPS())
	mouseInfo := fmt.Sprintf("Mouse: %t, %t",
		engine.GetInputManager().Mouse.State[0], engine.GetInputManager().Mouse.State[1])

	g.DebugFont.Draw(screen, fps, 10, 20, color.White)
	g.DebugFont.Draw(screen, mouseInfo, 10, 40, color.White)

	// g.DebugFont.Draw(screen, "Hello World!", 20, 20, color.White)

	g.Blob.Draw(screen, 50, 50)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func (g *Game) Start() {
	g.cam = engine.Camera{Position: engine.Vec2{X: 0, Y: -240}}
	g.audioContext = audio.NewContext(44100)

	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle("BULLETHELL")

	blobAnim, err := engine.LoadAnimation("spw.png", 16)
	if err != nil {
		log.Fatal(err)
	}

	blobAnim.AddFrame(0, 200.0, nil, nil)
	blobAnim.AddFrame(1, 200.0, nil, nil)
	blobAnim.AddFrame(2, 200.0, nil, nil)

	g.Blob = blobAnim

	fontPath, err := findfont.Find("arial.ttf")
	if err != nil {
		log.Fatal(err)
	}

	dbgFont, err := engine.OpenFont(fontPath, 12.0, 72.0)
	if err != nil {
		log.Fatal(err)
	}

	g.DebugFont = dbgFont
}

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	gameInstance := Game{ScreenWidth: GameScreenWidth, ScreenHeight: GameScreenHeight}
	gameInstance.Start()

	if err := ebiten.RunGame(&gameInstance); err != nil {
		log.Fatal(err)
	}
}
