package main

import (
	"fmt"
	"sync"
)

var x = 0
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	x = x + 1
	mutex.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
