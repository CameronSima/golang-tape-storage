package frame

import (
	"testing"

	"github.com/cameronsima/yt-tape-go/pixel"
)

func TestNewFrame(t *testing.T) {
	frame := NewFrame(10, 10)
	if frame.Width != 10 {
		t.Errorf("NewFrame() failed, got: %d, want: %d.", frame.Width, 10)
	}
	if frame.Height != 10 {
		t.Errorf("NewFrame() failed, got: %d, want: %d.", frame.Height, 10)
	}
	if len(frame.Data) != 10 {
		t.Errorf("NewFrame() failed, got: %d, want: %d.", len(frame.Data), 10)
	}
	for i := 0; i < 10; i++ {
		if len(frame.Data[i]) != 10 {
			t.Errorf("NewFrame() failed, got: %d, want: %d.", len(frame.Data[i]), 10)
		}
	}
}

func TestWritePixel(t *testing.T) {
	frame := NewFrame(100, 100)
	pixel := pixel.NewPixel("1", 10)
	frame.WritePixel(0, 0, *pixel)
	expected := [4]uint8{255, 0, 0, 255}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if frame.Data[i][j] != expected {
				t.Errorf("WritePixel() failed, got: %d, want: %d.", frame.Data[i][j], expected)
			}
		}
	}
}
