package main

import (
	"fmt"
	"time"
)

func worker(done chan bool, cmd string) {

	fmt.Println("worker start..", cmd)
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true

}
func main() {

	//str_chan := make(chan string, 10)
	//int_chan := make(chan int, 10)

	// slice init 1
	// var int_slice [4]int
	// int_slice[0] = 1
	// int_slice[1] = 1
	// int_slice[2] = 1
	// int_slice[3] = 1

	// slice init 2
	//int_slice := [4]int{1,2,3,4}

	// slice init 3
	//   int_slice := make([]int, 4)
	//    int_slice[0] = 0
	//    int_slice[1] = 1
	//    int_slice[2] = 2
	//    int_slice[3] = 3

	// slice init 4
	int_slice := []int{1, 2, 3, 4}

	message := "hello message"

	done := make(chan bool, 10)

	for _, v := range int_slice {
		fmt.Println(v)
	}

	for i := 0; i < 10; i++ {
		// str_chan<- message
		//  int_chan<-i
		go worker(done, message)
	}
	<-done

	/* for j:=0;j<10;j++{
	   fmt.Println(<-str_chan)
	   fmt.Println(<-int_chan)
	   }*/
}
