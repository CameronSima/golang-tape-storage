package main

import (
	"encoding/json"
	"os"
	"strings"

	vidio "github.com/AlexEidt/Vidio"
	"github.com/cameronsima/yt-tape-go/config"
	"github.com/cameronsima/yt-tape-go/pixel"
	"github.com/cameronsima/yt-tape-go/utils"
)

type VideoReader struct {
	filepath          string
	outfile           *os.File
	config            config.Config
	Metadata          Metadata
	video             *vidio.Video
	currentFrameCount int
}

func (r *VideoReader) Read() {
	bitsRead := 0
	for r.video.Read() {
		frame := r.video.FrameBuffer()

		if r.currentFrameCount == 1 {
			r.Metadata = r.getMetaData(frame)
		} else if bitsRead >= r.Metadata.FileSize*3 {

		} else {
			bitCount, eofReached := r.readFrame(frame, bitsRead)
			if eofReached {
				return
			}
			bitsRead += bitCount

		}
		r.currentFrameCount++
	}
	r.outfile.Close()
	r.video.Close()
}

func (r *VideoReader) readFrame(frameBytes []byte, totalBitsRead int) (int, bool) {
	octals := pixel.BytesToOctals(frameBytes, r.config)
	bitsRead := 0

	for o := range octals {
		bitsRead += 1
		//fmt.Println(r.Metadata.FrameCount, r.currentFrameCount-1, o.Height, r.Metadata.LastHeight, o.Width, r.Metadata.LastWidth)

		// EOF
		// if r.Metadata.FrameCount >= r.currentFrameCount-1 &&
		// 	o.Height >= r.Metadata.LastHeight &&
		// 	o.Width >= r.Metadata.LastWidth {
		// 	return
		// }

		octal := strings.Join(o.Data, "")
		r.outfile.Write(utils.OctalStringToBytes(octal))

		if r.Metadata.FrameCount == r.currentFrameCount-1 {
			if totalBitsRead+bitsRead >= r.Metadata.FileSize {
				return bitsRead, true
			}
		}

		// if r.Metadata.FrameCount == r.currentFrameCount-1 {
		// 	octal = strings.TrimSuffix(octal, "000")
		// }

	}
	return bitsRead, false
}

// Read frame bytes up to the EOF delimiter
func (r *VideoReader) getMetaData(frameBytes []byte) Metadata {
	octals := pixel.BytesToOctals(frameBytes, r.config)
	bits := make([]string, 0)
	m := Metadata{}

	for o := range octals {
		octal := strings.Join(o.Data, "")

		// reached EOF (*) delimiter for metadata
		if octal == r.config.EOFDelimiter {
			b := utils.OctalStringToBytes(strings.Join(bits, ""))

			err := json.Unmarshal(b, &m)
			if err != nil {
				panic(err)
			}
			break
		}
		bits = append(bits, o.Data...)
	}
	return m
}

func NewVideoReader(filepathIn string, filepathOut string, config config.Config) *VideoReader {
	video, err := vidio.NewVideo(filepathIn)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filepathOut)
	if err != nil {
		panic(err)
	}

	return &VideoReader{
		filepath:          filepathIn,
		outfile:           f,
		config:            config,
		video:             video,
		currentFrameCount: 1,
	}
}
