package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for range ticker.C {
			fmt.Println("tick")
		}
	}()

	time.Sleep(2 * time.Second)
	ticker.Stop()
	fmt.Println("tock")
}
