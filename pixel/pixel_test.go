package pixel

import (
	"testing"

	"github.com/cameronsima/yt-tape-go/config"
)

func TestNewPixel(t *testing.T) {
	pixel := NewPixel("1", 10)
	if pixel.Color != OctToPixel["1"] {
		t.Errorf("NewPixel() failed, got: %d, want: %d.", pixel.Color, OctToPixel["1"])
	}
}

func TestGetPixelContent(t *testing.T) {
	data := []byte{
		255, 0, 0, 255, 255, 0, 0, 255,
		255, 0, 0, 255, 255, 0, 0, 255,
		0, 0, 0, 255, 0, 0, 0, 255,
		0, 0, 0, 255, 0, 0, 0, 255,
	}
	c := config.NewConfig()
	c.PixelSize = 2
	c.VideoHeight = 2
	c.VideoWidth = 2

	pixelContent := GetPixelContent(0, 0, data, *c)

	expected := [][]Color{
		{
			{255, 0, 0, 255},
			{255, 0, 0, 255},
		},
		{
			{255, 0, 0, 255},
			{255, 0, 0, 255},
		},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if pixelContent[i][j] != expected[i][j] {
				t.Errorf("GetPixelContent() failed, got: %d, want: %d.", pixelContent[i][j], expected[i][j])
			}
		}
	}
}

func TestProcessPixelContent(t *testing.T) {
	c := config.NewConfig()
	c.PixelSize = 2
	pixelContent := [][]Color{
		{
			{0, 0, 255, 255},
			{1, 2, 254, 255},
		},
		{
			{0, 0, 255, 255},
			{1, 0, 254, 255},
		},
	}
	pixel := ProcessPixelContent(pixelContent, *c)
	expected := Color{0, 0, 255, 255}
	if pixel.Color != expected {
		t.Errorf("ProcessPixelContent() failed, got: %d, want: %d.", pixel.Color, OctToPixel["3"])
	}
}
