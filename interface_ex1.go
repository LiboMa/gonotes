package main

import "fmt"

func main() {
	var x interface{} = []int{1, 2, 3}
	fmt.Println(x)
}
