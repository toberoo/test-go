package main

import (
	"fmt"
	"time"
)

func sendHello(channel chan string) {
	channel <- " hello world"
}

func sendBye(channel chan string) {
	channel <- "good bye cruel world"
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

// 2 second pause
func workerOne(c chan bool) {
	fmt.Println("worker 1 started")
	time.Sleep(2 * time.Second)
	fmt.Println("worker 1 finished")
	c <- true
}

// 1 second pause
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

	// We can use select to handle different messages
	// in a loop depending on which channel.
	hChan := make(chan string)
	gChan := make(chan string)

	go sendHello(hChan)
	go sendBye(gChan)

	// We can use select on channels to handle different
	// channels sending us back info. I can see this being really
	// powerful.
	for i := 0; i < 4; i++ {
		select {
		case msg1 := <-hChan:
			fmt.Println(msg1, " recieved")
		case msg2 := <-gChan:
			fmt.Println(msg2, " recieved begrudingly")
		default:
			fmt.Println("No messages")
		}
	}

	// We can use timeouts to end stuff earlier
	// with the select statement. It's kinda
	// cheeky since time.After is a goRoutine lol.
	d := make(chan bool)
	select {
	case <-d:
		fmt.Println("Worker happened before timeout")
	case <-time.After(1 * time.Second):
		fmt.Println("We timed out")
	}

	// We can also normal iterate over channel
	// We have to identify the number
	nChan := make(chan int, 3)
	nChan <- 1
	nChan <- 2
	nChan <- 3
	// We have to close before iteration
	close(nChan)
	for n := range nChan {
		fmt.Println(n)
	}
}
