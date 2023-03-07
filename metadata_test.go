package main

import (
	"testing"

	"github.com/cameronsima/yt-tape-go/config"
)

func TestFileSize(t *testing.T) {
	c := config.NewConfig()
	c.PixelSize = 20
	size := NewMetadata("./testdata/simpletesttext.txt", *c).FileSize

	expected := 14
	if size != expected {
		t.Errorf("FileSize() failed, got: %d, want: %d.", size, expected)
	}
}

func TestFrameCount(t *testing.T) {
	c := config.NewConfig()
	c.PixelSize = 20
	count := NewMetadata("./testdata/simpletesttext.txt", *c).FrameCount

	expected := 1
	if count != expected {
		t.Errorf("FrameCount() failed, got: %d, want: %d.", count, expected)
	}
}

func TestFrameCount2(t *testing.T) {
	c := config.NewConfig()
	c.PixelSize = 10
	count := NewMetadata("./testdata/Big_Buck_Bunny_1080_10s_1MB.mp4", *c).FrameCount

	expected := 1023
	if count != expected {
		t.Errorf("FrameCount() failed, got: %d, want: %d.", count, expected)
	}
}

func TestLastWidthAndHeight(t *testing.T) {
	c := config.NewConfig()
	c.PixelSize = 20
	m := NewMetadata("./testdata/simpletesttext.txt", *c)

	expectedWidth := 10
	if m.LastWidth != expectedWidth {
		t.Errorf("LastWidth() failed, got: %d, want: %d.", m.LastWidth, expectedWidth)
	}

	expectedHeight := 2
	if m.LastHeight != expectedHeight {
		t.Errorf("LastHeight() failed, got: %d, want: %d.", m.LastHeight, expectedHeight)
	}
}

func TestLastWidthAndHeight2(t *testing.T) {
	c := config.NewConfig()
	c.PixelSize = 10
	m := NewMetadata("./testdata/testtext.txt", *c)

	expectedWidth := 2
	if m.LastWidth != expectedWidth {
		t.Errorf("LastWidth() failed, got: %d, want: %d.", m.LastWidth, expectedWidth)
	}

	expectedHeight := 5
	if m.LastHeight != expectedHeight {
		t.Errorf("LastHeight() failed, got: %d, want: %d.", m.LastHeight, expectedHeight)
	}
}

func TestDelimeter(t *testing.T) {
	c := config.NewConfig()
	c.PixelSize = 20
	m := NewMetadata("./testdata/simpletesttext.txt", *c)
	octals := m.ToOctals(*c)
	lastOctal := octals[len(octals)-1]
	expected := c.EOFDelimiter

	if lastOctal != expected {
		t.Errorf("Delimeter() failed, got: %s, want: %s.", lastOctal, expected)
	}

}
