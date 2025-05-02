package main

import (
	"fmt"
	"log"
	"io"
)

type MySlowReader struct {
	Contents string
	pos int
}

func (m *MySlowReader) Read(p []byte) (n int, err error) {
    if m.pos >= len(m.Contents) {
        return 0, io.EOF
    }
    if len(p) == 0 {
        return 0, nil // Handle empty buffer per io.Reader contract
    }
    n = copy(p, m.Contents[m.pos:m.pos+1]) // Assign to named return `n`
    if n > 0 {
        m.pos++
    }
    return n, nil
}

func main() {
	mySlowReaderInstance := &MySlowReader{
		Contents: "test",
	}

	out,err := io.ReadAll(mySlowReaderInstance)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test %s\n", out) 
}
