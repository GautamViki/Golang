package main

import (
	"fmt"
	"time"
)

func main() {
	// Waiting on Multiple Channels
	chan1 := make(chan string) // creating channel
	chan2 := make(chan string)

	// this is called clouser function
	go func() {
		time.Sleep(3 * time.Second)
		chan1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}
}
