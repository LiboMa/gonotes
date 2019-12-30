package main

import "fmt"

func main() {

	Aarray := make([]int, 3)
	Barray := make([]int, 3)

	Aarray[0] = 1
	Barray[0] = 1

	fmt.Println(Aarray, Barray)
	fmt.Printf("%T\t%T", Aarray, Barray)

	// fmt.Println("Result: ", CompareArray(Aarray, Barray))
	fmt.Println("Result: ", equal(Aarray, Barray))

}

//func CompareArray(a, b []int) bool {
//
//	if len(a) != len(b) {
//		return false
//	}
//	if a != b {
//		return false
//	}
//
//	return true
//}

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
