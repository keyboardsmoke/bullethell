package engine

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type Sound struct {
	filename string
	player   *audio.Player
}

type Context struct {
	audioContext *audio.Context
}

func (c *Context) LoadWAV(filename string, volume float64) (*Sound, error) {
	fd, fe := os.Open(filename)
	if fe != nil {
		return nil, fe
	}

	ss, se := wav.DecodeWithSampleRate(c.audioContext.SampleRate(), fd)
	if se != nil {
		return nil, se
	}

	mp, me := c.audioContext.NewPlayer(ss)
	if me != nil {
		return nil, me
	}

	mp.SetVolume(volume)

	return &Sound{filename: filename, player: mp}, nil
}

func (c *Context) LoadMP3(filename string, volume float64) (*Sound, error) {
	fd, fe := os.Open(filename)
	if fe != nil {
		return nil, fe
	}

	ss, se := mp3.DecodeWithSampleRate(c.audioContext.SampleRate(), fd)
	if se != nil {
		return nil, se
	}

	mp, me := c.audioContext.NewPlayer(ss)
	if me != nil {
		return nil, me
	}

	mp.SetVolume(volume)

	return &Sound{filename: filename, player: mp}, nil
}

func (c *Context) LoadOGG(filename string, volume float64) (*Sound, error) {
	fd, fe := os.Open(filename)
	if fe != nil {
		return nil, fe
	}

	ss, se := vorbis.DecodeWithSampleRate(c.audioContext.SampleRate(), fd)
	if se != nil {
		return nil, se
	}

	mp, me := c.audioContext.NewPlayer(ss)
	if me != nil {
		return nil, me
	}

	mp.SetVolume(volume)

	return &Sound{filename: filename, player: mp}, nil
}

func (s *Sound) GetPlayer() *audio.Player {
	return s.player
}

func (s *Sound) GetFilename() string {
	return s.filename
}

func CreateAudioContext(sampleRate int) *Context {
	return &Context{audioContext: audio.NewContext(sampleRate)}
}
