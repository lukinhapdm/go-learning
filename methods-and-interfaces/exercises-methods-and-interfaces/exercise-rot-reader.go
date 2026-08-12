package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r13.r.Read(p)
	for i := 0; i < n; i++ {
		b := p[i]
		switch {
		case b >= 'a' && b <= 'z':
			p[i] = 'a' + (b-'a'+13)%26
		case b >= 'A' && b <= 'Z':
			p[i] = 'A' + (b-'A'+13)%26
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
