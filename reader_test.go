package main

import (
	"os"
	"strings"
	"testing"

	"github.com/cameronsima/yt-tape-go/config"
	"github.com/cameronsima/yt-tape-go/utils"
)

func TestNewVideoReader(t *testing.T) {
	c := config.NewConfig()
	c.PixelSize = 10
	WriteVideo("./testdata/simpletesttext.txt", "./testdata/testvideo.mp4", *c)
	reader := NewVideoReader("./testdata/testvideo.mp4", "./testdata/out.txt", *c)
	reader.Read()

	if reader.Metadata.FileSize != 14 {
		t.Errorf("NewVideoReader() failed, got: %d, want: %d.", reader.Metadata.FileSize, 14)
	}

	// read file
	for b := range utils.ReadFile("./testdata/out.txt") {
		expected := "tuscon arizona"
		if string(b) != expected {
			t.Errorf("NewVideoReader() failed, got: %s, want: %s.", string(b), expected)
		}
	}
}

func TestNewVideoReader2(t *testing.T) {
	c := config.NewConfig()
	c.PixelSize = 10
	WriteVideo("./testdata/testtext.txt", "./testdata/testvideo2.mp4", *c)
	reader := NewVideoReader("./testdata/testvideo2.mp4", "./testdata/out2.txt", *c)
	reader.Read()

	// read file
	for b := range utils.ReadFile("./testdata/out2.txt") {
		expected := "Here's some test data with some numbers: 20, 530, 54,344"
		if !strings.Contains(string(b), expected) {
			t.Errorf("TestNewVideoReader2 failed, got: %s, want: %s.", string(b), expected)
		}
	}
}

func TestNewVideoReader3(t *testing.T) {
	c := config.NewConfig()
	c.PixelSize = 10
	WriteVideo("./testdata/sonnets.txt", "./testdata/testvideo3.mp4", *c)
	reader := NewVideoReader("./testdata/testvideo3.mp4", "./testdata/out3.txt", *c)
	reader.Read()

	if _, err := os.Stat("./testdata/out3.txt"); os.IsNotExist(err) {
		t.Errorf("TestNewVideoReader3 failed, file not created.")
	}
}

func TestNewVideoReader4(t *testing.T) {
	c := config.NewConfig()
	c.PixelSize = 10
	c.VideoFps = 30
	WriteVideo("./testdata/Big_Buck_Bunny_1080_10s_1MB.mp4", "./testdata/testvideo4.mp4", *c)
	reader := NewVideoReader("./testdata/testvideo4.mp4", "./testdata/out4.mp4", *c)
	reader.Read()

	if _, err := os.Stat("./testdata/out4.mp4"); os.IsNotExist(err) {
		t.Errorf("WriteVideo() failed, file not created.")
	}
}

func TestNewVideoReader5(t *testing.T) {
	c := config.NewConfig()
	c.PixelSize = 10
	c.VideoFps = 30
	WriteVideo("./testdata/jefferson-hospital-exterior-1-scaled.webp", "./testdata/testvideo5.mp4", *c)
	reader := NewVideoReader("./testdata/testvideo5.mp4", "./testdata/out5.webp", *c)
	reader.Read()

	if _, err := os.Stat("./testdata/out5.webp"); os.IsNotExist(err) {
		t.Errorf("WriteVideo() failed, file not created.")
	}
}
