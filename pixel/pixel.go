package pixel

import (
	"github.com/cameronsima/yt-tape-go/config"
	"github.com/cameronsima/yt-tape-go/utils"
)

type Color = [4]uint8

type PixelContent [][]Color

type Pixel struct {
	Color Color
	Size  int
	Bit   string
}

var OctToPixel = map[string]Color{
	"0": {0, 0, 0, 255},
	"1": {255, 0, 0, 255},
	"2": {0, 255, 0, 255},
	"3": {0, 0, 255, 255},
	"4": {255, 255, 0, 255},
	"5": {255, 0, 255, 255},
	"6": {0, 255, 255, 255},
	"7": {255, 255, 255, 255},
}

var PixelToOct = map[Color]string{
	{0, 0, 0, 255}:       "0",
	{255, 0, 0, 255}:     "1",
	{0, 255, 0, 255}:     "2",
	{0, 0, 255, 255}:     "3",
	{255, 255, 0, 255}:   "4",
	{255, 0, 255, 255}:   "5",
	{0, 255, 255, 255}:   "6",
	{255, 255, 255, 255}: "7",
}

func NewPixel(bit string, size int) *Pixel {
	color, found := OctToPixel[bit]
	if !found {
		panic("Invalid bit")
	}
	return &Pixel{color, size, bit}
}

func GetPixelContent(x int, y int, data []byte, c config.Config) PixelContent {
	content := make(PixelContent, c.PixelSize)
	for i := 0; i < c.PixelSize; i++ {
		content[i] = make([]Color, c.PixelSize)
		for j := 0; j < c.PixelSize; j++ {
			content[i][j] = Color{
				data[(y+i)*c.VideoWidth*4+(x+j)*4],
				data[(y+i)*c.VideoWidth*4+(x+j)*4+1],
				data[(y+i)*c.VideoWidth*4+(x+j)*4+2],
				data[(y+i)*c.VideoWidth*4+(x+j)*4+3],
			}
		}
	}
	return content
}

func ProcessPixelContent(content PixelContent, c config.Config) *Pixel {
	r := make([]uint8, 0)
	g := make([]uint8, 0)
	b := make([]uint8, 0)
	a := make([]uint8, 0)
	for i := 0; i < c.PixelSize; i++ {
		for j := 0; j < c.PixelSize; j++ {
			r = append(r, content[i][j][0])
			g = append(g, content[i][j][1])
			b = append(b, content[i][j][2])
			a = append(a, content[i][j][3])
		}
	}
	avgR := utils.Average(r)
	avgG := utils.Average(g)
	avgB := utils.Average(b)
	avgA := utils.Average(a)
	avgColor := Color{avgR, avgG, avgB, avgA}
	return ProcessColor(avgColor, c)
}

func ProcessColor(color Color, c config.Config) *Pixel {
	avgColor := color
	for i, c := range avgColor {
		if c > 125 {
			avgColor[i] = 255
		} else {
			avgColor[i] = 0
		}
	}
	bit := PixelToOct[avgColor]
	return &Pixel{avgColor, c.PixelSize, bit}
}

type FrameBuffer struct {
	Width  int
	Height int
	Data   []string
}

// turn an array of bytes read from an encoded video frame, extract the pixel data, and convert it to octals
func BytesToOctals(data []byte, c config.Config) chan FrameBuffer {
	ch := make(chan FrameBuffer)

	go func() {
		//octals := make([][]string, 0)
		octal := make([]string, 0)
		for i := 0; i < c.VideoHeight/c.PixelSize; i++ {
			for j := 0; j < c.VideoWidth/c.PixelSize; j++ {
				content := GetPixelContent(j*c.PixelSize, i*c.PixelSize, data, c)
				pixel := ProcessPixelContent(content, c)
				octal = append(octal, pixel.Bit)
				if len(octal) == 3 {
					fb := FrameBuffer{Width: j, Height: i + 1, Data: octal}
					ch <- fb
					octal = make([]string, 0)
				}
			}
		}
		close(ch)
	}()
	return ch
}
