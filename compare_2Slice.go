package main

import (
	"bytes"
	"fmt"
)

func main() {

	SliceA := []int{1, 2, 3, 4}
	SliceB := []int{1, 2, 3, 4}

	SliceC := []byte("1 2 3 4")
	SliceD := []byte("1 2 3 4")

	fmt.Println(SliceA, SliceB, SliceC, SliceD)

	// Compare two byte arrays

	fmt.Printf("SliceA(%v) == SliceB(%v), Test Result: %v \n", SliceA, SliceB, equal(SliceA, SliceB))
	fmt.Printf("SliceC(%v) == SliceD(%v), Test Result: %v \n", SliceC, SliceD, bytes.Equal(SliceC, SliceD))

}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
