package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	//wgdone := make(chan bool, 1)
	//var wgdone = make(chan bool, 1)
	var wgdone = make(chan bool, 1)

	var loop = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	start := time.Now()

	for _, n := range loop {
		fmt.Println(n)
		wg.Add(1)
		go func(n int, wgdone chan bool, wg *sync.WaitGroup) {
			s := n * 2
			fmt.Println("Program 2 Execution....")
			time.Sleep(time.Second * 3)
			log.Printf("input: %d, double number: %d\n\n", n, s)
			wg.Done()
			wgdone <- true
		}(n, wgdone, &wg)
	}
	<-wgdone
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Time elapsed =>", elapsed)

}
