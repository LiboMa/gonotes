package main

import "crypto/sha256"
import "fmt"

func main() {
	show_string_up := "HelloWorld"
	show_string_low := "helloworld"
	c1 := sha256.Sum256([]byte(show_string_up))
	c2 := sha256.Sum256([]byte(show_string_low))
	fmt.Println(len(c1))
	fmt.Println(len(c2))
	fmt.Printf("%x\n%x\n", c1, c2)
}
