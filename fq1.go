package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	v := 1
	v2 := incr(&v)
	fmt.Printf("%p, %p\n", &v, &v2)
	fmt.Println(v, v2)
}
