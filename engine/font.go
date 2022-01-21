package engine

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Font struct {
	face font.Face
	size float64
	dpi  float64
}

func (f *Font) GetSize() float64 {
	return f.size
}

func (f *Font) GetDpi() float64 {
	return f.dpi
}

func (f *Font) Draw(dst *ebiten.Image, str string, x, y float64, clr color.Color) {
	text.Draw(dst, str, f.face, int(x), int(y), clr)
}

func (f *Font) Close() error {
	return f.face.Close()
}

func OpenFont(path string, size float64, dpi float64) (*Font, error) {
	fd, fe := os.ReadFile(path)
	if fe != nil {
		return nil, fe
	}

	tt, te := opentype.Parse(fd)
	if te != nil {
		return nil, te
	}

	fontFace, err := opentype.NewFace(tt, &opentype.FaceOptions{Size: size, DPI: dpi, Hinting: font.HintingFull})

	if err != nil {
		return nil, err
	}

	return &Font{face: fontFace, size: size, dpi: dpi}, nil
}
