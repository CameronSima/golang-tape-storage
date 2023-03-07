package utils

import (
	"strings"
	"testing"
)

func TestBytesToOctalString(t *testing.T) {
	bytes_obj := []byte("Hello World!")
	octals := BytesToOctalString(bytes_obj)
	expected := "110145154154157040127157162154144041"
	if strings.Join(octals, "") != expected {
		t.Errorf("BytesToOctalString() failed, got: %s, want: %s.", strings.Join(octals, ""), expected)
	}
}

func TestOctalStringToBytes(t *testing.T) {
	octal_string := "110145154154157040127157162154144041"
	bytes_obj := OctalStringToBytes(octal_string)
	expected := []byte("Hello World!")
	if string(bytes_obj) != string(expected) {
		t.Errorf("octal_string_to_bytes() failed, got: %s, want: %s.", bytes_obj, expected)
	}
}

// func TestReadFile(t *testing.T) {
// 	filename := "../testdata/simpletesttext.txt"
// 	for bytes_obj := range ReadFile(filename) {

// 		expected := []byte("tuscon arizona")
// 		if string(bytes_obj[:len(expected)]) != string(expected) {
// 			t.Errorf("read_file() failed, got: %s, want: %s.", bytes_obj[:len(expected)], expected)
// 		}
// 	}
// }
