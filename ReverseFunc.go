package main

import "fmt"

func main() {
	SliceA := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("Before Reverse:", SliceA)
	reverse(SliceA)
	fmt.Println("After Reverse:", SliceA)

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("i:%v, j:%v\n", i, j)
		s[i], s[j] = s[j], s[i]
	}
	// fmt.Println(s)
}
