package main

import (
	"encoding/json"
	"os"

	"github.com/cameronsima/yt-tape-go/config"
	"github.com/cameronsima/yt-tape-go/utils"
)

type Metadata struct {
	LastWidth  int
	LastHeight int
	FrameCount int
	FileSize   int
	Filename   string
}

func (m *Metadata) ToOctals(c config.Config) []string {
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	octals := utils.BytesToOctalString(b)
	octals = append(octals, c.EOFDelimiter)
	return octals
}

func NewMetadata(filename string, c config.Config) *Metadata {
	fi, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}

	filesize := int(fi.Size())
	numOctals := filesize * 3

	frameSize := c.VideoWidth * c.VideoHeight

	horizontalPixels := c.VideoWidth / c.PixelSize
	verticalPixels := c.VideoHeight / c.PixelSize
	frameCount := numOctals / (horizontalPixels * verticalPixels)
	//lastHeight := numOctals/horizontalPixels + 1
	lastWidth := numOctals % horizontalPixels

	// frameCount := numOctals / frameSize
	lastHeight := (numOctals % frameSize) / c.VideoHeight
	// lastWidth := (numOctals % frameSize) % c.VideoHeight

	return &Metadata{
		LastWidth:  lastWidth,
		LastHeight: lastHeight,
		FrameCount: frameCount + 1,
		FileSize:   filesize,
		Filename:   filename,
	}
}
