package frame

import (
	"github.com/cameronsima/yt-tape-go/pixel"
)

type Frame struct {
	Width  int
	Height int
	Data   [][]pixel.Color
}

func (f *Frame) WritePixel(x int, y int, pixel pixel.Pixel) {
	for i := 0; i < pixel.Size; i++ {
		for j := 0; j < pixel.Size; j++ {
			f.Data[y+i][x+j] = pixel.Color
		}
	}
}

func (f *Frame) ToBytes() []byte {
	bytes := make([]byte, 0)
	for i := 0; i < f.Height; i++ {
		for j := 0; j < f.Width; j++ {
			p := f.Data[i][j]

			for _, k := range p {
				bytes = append(bytes, k)
			}
		}
	}
	return bytes
}

func NewFrame(width int, height int) *Frame {
	backgroundColor := pixel.OctToPixel["0"]
	data := make([][]pixel.Color, height)

	for i := 0; i < height; i++ {
		data[i] = make([]pixel.Color, width)
		for j := 0; j < width; j++ {
			data[i][j] = backgroundColor
		}
	}
	return &Frame{
		Width:  width,
		Height: height,
		Data:   data,
	}
}
