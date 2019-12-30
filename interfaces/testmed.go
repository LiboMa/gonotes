package main

import (
	"fmt"
	"time"
)

type Cat struct {
	age    int
	gender int
}

func (c *Cat) Spark() {
	c.age = 1
	fmt.Println("meao", c.age)
}

func (c *Cat) Sleep(t time.Duration) {
	c.age += 1
	fmt.Println("meao", c.age)
	time.Sleep(t)
}

func (c *Cat) Run() {
	c.age += 1
	fmt.Println("meao", c.age)
}

func main() {

	c := Cat{}
	c.Spark()
	c.Sleep(time.Second)
	c.Run()

}
