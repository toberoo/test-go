package main

import (
	"fmt"
	"time"
)

func sendHello(channel chan string) {
	channel <- " hello world"
}

// <-chan means it reads message
func readMessage(channel <-chan string) {
	message := <-channel
	fmt.Println(message)
}

// chan<- means it only sends message
func sendNumbers(channel chan<- int) {
	for i := 1; i < 11; i++ {
		channel <- i
	}
}

func workerOne(c chan bool) {
	fmt.Println("worker 1 started")
	time.Sleep(time.Second)
	fmt.Println("worker 1 finished")
	c <- true
}

func workerTwo(c chan bool) {
	fmt.Println("worker 2 started")
	time.Sleep(time.Second)
	fmt.Println("worker 2 finished")
	c <- true
}

func main() {

	// Blocks till channel is read
	channel := make(chan string)
	go sendHello(channel)
	readMessage(channel)

	// Buffered channels will block till
	// the number of items is sent
	bufferedChannel := make(chan int, 2)
	go sendNumbers(bufferedChannel)
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

	done := make(chan bool, 2)
	go workerOne(done)
	go workerTwo(done)
	<-done
	<-done
}
