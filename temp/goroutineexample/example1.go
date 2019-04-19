package main

import (
	"fmt"
	"time"
)

var ci = make(chan int) // unbuffered
func a() {
	//cj := make(chan int, 0)        // buffered
	//cs := make(chan *os.File, 100) // buffered
	// Unbuffered channels combine communication—the exchange of a value—with synchronization—guaranteeing that two calculations (goroutines) are in a known state.
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("in there")
		ci <- 1
	}()
}

func main() {
	//fmt.Println("hello")
	a()
	fmt.Println("out ...")
	<-ci
}
