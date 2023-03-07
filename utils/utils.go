package utils

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

const FileBufferSize = 1024

func BytesToOctalString(bytes []byte) []string {
	result := make([]string, 0)
	for _, b := range bytes {
		if len((fmt.Sprintf("%03o", b))) != 3 {
			panic("Octal string is not 3 characters long")
		}
		result = append(result, fmt.Sprintf("%03o", b))
	}
	return result
}

func OctalStringToBytes(octal_string string) []byte {
	var result []byte
	for i := 0; i < len(octal_string); i += 3 {
		o := octal_string[i : i+3]
		b, _ := strconv.ParseInt(o, 8, 8)
		result = append(result, byte(b))
	}
	return result
}

func ReadFile(filename string) chan []byte {
	c := make(chan []byte)
	go func() {
		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		for {
			buffer := make([]byte, FileBufferSize)
			bytes_read, err := f.Read(buffer)
			if err != nil {
				if err != io.EOF {
					panic(err)
				} else {
					break
				}
			}
			c <- buffer[:bytes_read]
		}
		close(c)
	}()
	return c
}

func WriteFile(filename string, data []byte) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(data)
}

func Average(nums []uint8) uint8 {
	var sum int
	for _, n := range nums {
		sum += int(n)
	}
	return uint8(sum / len(nums))
}
