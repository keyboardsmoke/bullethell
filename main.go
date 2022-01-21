package main

import (
	"bh/engine"
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
	newTime      int64
	oldTime      int64
	deltaTime    float64
	cam          engine.Camera
	audioContext *audio.Context
	ScreenWidth  int
	ScreenHeight int
	DebugFont    *engine.Font
	Blob         *engine.Animation
	// Blob         *engine.Sprite

}

func (g *Game) Update() error {
	g.newTime = time.Now().UnixNano()
	g.deltaTime = float64(((g.newTime - g.oldTime) / 1000000)) * 0.001
	g.oldTime = g.newTime

	g.Blob.Update(g.deltaTime)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DebugFont.Draw(screen, "Hello World!", 20, 20, color.White)

	// g.Blob.DrawFrame(screen, 50, 50, 0)
	// g.Blob.DrawFrame(screen, 50, 100, g.Blob.NumberOfFrames-1)
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

	blobAnim.AddFrame(0, 100.0, nil, nil)
	blobAnim.AddFrame(1, 100.0, nil, nil)

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
