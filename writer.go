package main

import (
	vidio "github.com/AlexEidt/Vidio"
	config "github.com/cameronsima/yt-tape-go/config"
	"github.com/cameronsima/yt-tape-go/frame"
	pixel "github.com/cameronsima/yt-tape-go/pixel"
	utils "github.com/cameronsima/yt-tape-go/utils"
)

type VideoWriter struct {
	filepath          string
	currentWidth      int
	currentHeight     int
	currentFrameCount int
	currentFrame      *frame.Frame
	config            config.Config
	video             *vidio.VideoWriter
}

type OctalGenerator func() chan []byte

func NewVideoWriter(filepath string, config config.Config) *VideoWriter {
	options := vidio.Options{
		Codec:   "h264",
		FPS:     config.VideoFps,
		Quality: 1,
	}

	video, err := vidio.NewVideoWriter(filepath, config.VideoWidth, config.VideoHeight, &options)
	if err != nil {
		panic(err)
	}
	return &VideoWriter{
		filepath:          filepath,
		config:            config,
		currentWidth:      0,
		currentHeight:     0,
		currentFrameCount: 1,
		currentFrame:      frame.NewFrame(config.VideoWidth, config.VideoHeight),
		video:             video,
	}
}

func (w *VideoWriter) WriteData(octals []string) {
	for _, octal := range octals {
		for _, bit := range octal {
			bit_str := string(bit)

			pixel := pixel.NewPixel(bit_str, w.config.PixelSize)
			w.currentFrame.WritePixel(w.currentWidth, w.currentHeight, *pixel)
			w.currentWidth += w.config.PixelSize

			if w.currentWidth > w.config.VideoWidth-w.config.PixelSize {
				w.currentWidth = 0
				w.currentHeight += w.config.PixelSize
			}

			if w.currentHeight > w.config.VideoHeight-w.config.PixelSize {
				w.WriteFrame()
			}
		}
	}
}

func (w *VideoWriter) WriteMetaData(m *Metadata) {
	octals := m.ToOctals(w.config)
	w.WriteData(octals)
	w.WriteFrame()
}

func (w *VideoWriter) WriteFrame() {
	bytes := w.currentFrame.ToBytes()
	w.video.Write(bytes)
	w.currentFrame = frame.NewFrame(w.config.VideoWidth, w.config.VideoHeight)
	w.currentWidth = 0
	w.currentHeight = 0
	w.currentFrameCount += 1
}

func WriteVideo(file_in_path string, video_out_path string, config config.Config) *VideoWriter {
	writer := NewVideoWriter(video_out_path, config)

	// write metadata
	metadata := NewMetadata(file_in_path, config)
	writer.WriteMetaData(metadata)

	// write the rest of the frames
	for b := range utils.ReadFile(file_in_path) {
		octals := utils.BytesToOctalString(b)
		writer.WriteData(octals)
	}
	// write the last frame
	writer.WriteFrame()
	writer.video.Close()
	return writer
}
