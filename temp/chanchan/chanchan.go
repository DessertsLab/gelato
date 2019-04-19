package main

import (
	"fmt"
	"time"
)

func main() {
	requestChan := make(chan chan string)
	go goroutineC(requestChan)
	go goroutineD(requestChan)

	time.Sleep(time.Second)

}

func goroutineC(requestChan chan chan string) {

	responseChan := make(chan string)

	requestChan <- responseChan

	response := <-responseChan

	fmt.Printf("Response: %v\n", response)
}

func goroutineD(requestChan chan chan string) {
	responseChan := <-requestChan
	responseChan <- "wassup!"
}
