package main

import (
	"os"
	"testing"

	config "github.com/cameronsima/yt-tape-go/config"
)

func TestWriteVideo(t *testing.T) {
	c := config.NewConfig()
	c.VideoFps = 1
	c.PixelSize = 20
	WriteVideo("./testdata/simpletesttext.txt", "./testdata/testvideo.mp4", *c)

	if _, err := os.Stat("./testdata/testvideo.mp4"); os.IsNotExist(err) {
		t.Errorf("WriteVideo() failed, file not created.")
	}

}

func TestWriteVideoMultilineText(t *testing.T) {
	c := config.NewConfig()
	c.VideoFps = 1
	c.PixelSize = 20
	WriteVideo("./testdata/testtext.txt", "./testdata/testvideo2.mp4", *c)

	if _, err := os.Stat("./testdata/testvideo2.mp4"); os.IsNotExist(err) {
		t.Errorf("WriteVideo() failed, file not created.")
	}
}

func TestWriteVideoMp4(t *testing.T) {
	c := config.NewConfig()
	c.VideoFps = 30
	c.PixelSize = 10
	WriteVideo("./testdata/Big_Buck_Bunny_1080_10s_1MB.mp4", "./testdata/testvideo3.mp4", *c)

	if _, err := os.Stat("./testdata/testvideo3.mp4"); os.IsNotExist(err) {
		t.Errorf("WriteVideo() failed, file not created.")
	}
}
