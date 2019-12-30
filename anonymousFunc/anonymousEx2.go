package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	var loop = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	start := time.Now()

	for _, n := range loop {
		fmt.Println(n)
		wg.Add(1)
		go func(n int, wg *sync.WaitGroup) {
			s := n * 2
			fmt.Println("Program 2 Execution....")
			time.Sleep(time.Second * 1)
			fmt.Printf("input: %d, double number: %d\n\n", n, s)
			wg.Done()
		}(n, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Time elapsed =>", elapsed)

}
