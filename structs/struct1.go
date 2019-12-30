package main

import "fmt"

type User struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

func main() {
	fmt.Println("vim-go")

	u := User{1, "peter", 75.5}
	fmt.Println(u)
}
