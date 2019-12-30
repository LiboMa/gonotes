package main

import "fmt"
import "io"
import "os"
import "bytes"

func main() {

	var w io.Writer

	//w.Write([]byte("hello")) // panic, nil pointer dereferences
	w = os.Stdout
	if w == nil {
		fmt.Println("w is nil")
	} else {
		fmt.Printf("%v, %p, %T\n", w, w, w)
		// envs := os.Environ()
		// fmt.Println(envs)
		w.Write([]byte("hello\n")) // it may work
	}
	os.Stdout.Write([]byte("hello"))

	fmt.Println("using byte.Buffer object")

	w = new(bytes.Buffer)
	fmt.Printf("%v, %p, %T\n", w, w, w)
}
