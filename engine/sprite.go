package engine

import (
	"errors"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Filename        string
	Image           *ebiten.Image
	CellSize        uint
	NumberOfRows    uint
	NumberOfColumns uint
	NumberOfFrames  uint
}

func (s *Sprite) DrawFrame(dst *ebiten.Image, x, y int, frame uint) {
	sx := frame % s.NumberOfColumns * s.CellSize
	sy := frame / s.NumberOfColumns * s.CellSize
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y)) // ???
	rect := image.Rect(int(sx), int(sy), int(sx+s.CellSize), int(sy+s.CellSize))
	dst.DrawImage(s.Image.SubImage(rect).(*ebiten.Image), op)
}

func (s *Sprite) Dispose() {
	s.Image.Dispose()
}

func LoadSprite(filename string, cellSize uint) (*Sprite, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(fd)
	if err != nil {
		log.Fatal(err)
	}

	imageWidth := uint(img.Bounds().Dx())
	imageHeight := uint(img.Bounds().Dy())

	log.Default().Printf("imageWidth: %d, imageHeight: %d", imageWidth, imageHeight)
	log.Default().Printf("cellSize: %d", cellSize)

	if (imageWidth%cellSize) != 0 || (imageHeight%cellSize) != 0 {
		return nil, errors.New("image width or height is not a multiple of cell size")
	}

	var cols uint = imageWidth / cellSize
	var rows uint = imageHeight / cellSize
	var maxFrames uint = cols * rows

	ebiImage := ebiten.NewImageFromImage(img)

	return &Sprite{
		Filename:        filename,
		Image:           ebiImage,
		CellSize:        cellSize,
		NumberOfRows:    rows,
		NumberOfColumns: cols,
		NumberOfFrames:  maxFrames,
	}, nil
}
