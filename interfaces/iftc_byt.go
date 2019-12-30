package main

import (
	"bytes"
	"io"
)

const debug = true

// error version
//func main() {
//	var buf *bytes.Buffer
//
//	if debug {
//		buf = new(bytes.Buffer)
//	}
//	f(buf)
//}
//
//func f(out io.Writer) {
//	if out != nil {
//		out.Write([]byte("done!\n"))
//	}
//}

func main() {
	var buf io.Writer // abstract type  interface nil

	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
