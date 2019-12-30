package main

import (
	"fmt"
	"sort"
)

func main() {

	tstring := []string{"x", "axxx", "123x", "world", "hello"}

	sort.Strings(tstring)

	fmt.Println(tstring)

}
