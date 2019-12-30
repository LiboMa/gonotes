package main

import (
	"bytes"
	"fmt"
	"io"
	//"strings"
)

//import "strconv"

//type SetInt struct {
//	Na int
//	Nb string
//}
//
//func (s *SetInt) String() string {
//
//	return s.Nb
//}
//

type IntSet struct{ a, b int }

func (i *IntSet) String() string {
	//	s := Aoti()
	str := fmt.Sprintf("%d %d", i.a, i.b)
	return str
}

func main() {

	//	var ss = SetInt{1, "item1"}
	//var _ = IntSet{}.String()

	// empty struct
	var ss IntSet
	fmt.Println(ss.String())

	// emtpy interface

	var any interface{}

	any = true
	fmt.Println(any)
	any = 12.34
	fmt.Println(any)
	any = "hello"
	fmt.Println(any)

	// type assertion
	any = map[string]int{"one": 1, "two": 2}
	fmt.Println(any.(map[string]int)["one"])
	fmt.Println(any.(map[string]int)["two"])
	any = new(bytes.Buffer)

	fmt.Println(any)

	//any = new(strings.NewBuffer)

	// &bytes.Buffer must satisfy ioWriter
	var w io.Writer = new(bytes.Buffer)

	fmt.Printf("%T", w)

	//	fmt.Println(ss.String())
}
