package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {

	*c += ByteCounter(len(p))
	var a io.Writer
	re := strings.NewReader(string(p))
	bre := bytes.NewReader(p)
	resp, _ := http.Get("https://www.thbex.com/otc/")

	// method 1
	// buf := new(bytes.Buffer) // new Empty Buffer Struct
	// buf.ReadFrom(resp.Body)  // Read all buffer to Struct
	// body := buf.String()     // Convert Byte Buffer to String
	// fmt.Println(body)

	// method 2
	bodySec, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodySec))

	bre.Len()
	re.Size()
	fmt.Println(re.Size(), bre.Len(), a)
	// return len(p), nil
	return re.Len(), nil
}

func HttpGetDataFromIOReader(url string) {

	// Way 1 read vi io.Reader
	fmt.Println("Method 1: read data via ioutil:")
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body, string(body))
	// Way 2 read via buffer
	fmt.Println("Method 2: read data via buffer: ")

	respSec, _ := http.Get(url)
	buf := new(bytes.Buffer)
	buf.ReadFrom(respSec.Body)
	buf.Bytes()
	fmt.Println(buf.Bytes(), buf.String())

}
func main() {

	var c ByteCounter

	c.Write([]byte("hello"))
	fmt.Println(c)

	HttpGetDataFromIOReader("https://thbex.com/otc/")

	c = 0

	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	c = 0
	var name2 = "Polly"
	fmt.Fprintf(&c, "hey, %s", name2)
	fmt.Println(c)

}
