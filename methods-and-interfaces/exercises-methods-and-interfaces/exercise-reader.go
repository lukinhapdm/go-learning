package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Read populates the given slice with 'A' characters and returns the number of bytes written and a nil error
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
