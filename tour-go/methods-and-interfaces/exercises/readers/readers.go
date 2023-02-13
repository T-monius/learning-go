package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (i int,e error) {
	for v := range b {
		b[v] = 'A'
	}

	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
