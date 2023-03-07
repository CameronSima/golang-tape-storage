package config

type Config struct {
	VideoHeight  int
	VideoWidth   int
	VideoFps     float64
	PixelSize    int
	EOFDelimiter string
}

func NewConfig() *Config {
	return &Config{
		VideoHeight: 480,
		VideoWidth:  640,
		VideoFps:    1,
		PixelSize:   2,

		// 052 is the octal representation of the EOF delimiter ('*')
		EOFDelimiter: "052",
	}
}
