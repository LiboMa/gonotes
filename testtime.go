package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	t := start.Format(time.RFC3339)

	var a uint64
	a = 3

	fmt.Println(a)
	fmt.Printf("%p", &a)
	time.Sleep(time.Second * time.Duration(a))

	fmt.Printf("time start is: ", t)
}
