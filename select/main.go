package main

import (
	"fmt"
	"time"
)

func main() {
	// Example 1. Waiting on Multiple Channels
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

	// Example 2: Default Case
	// The default case in a select statement is executed if none of the channels are ready.
	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "Default Case"
	}()

	select {
	case msg := <-chan1:
		fmt.Println("Recieved ", msg)
	default:
		fmt.Println("No messaged recieved")
	}

	// Example 3: Timeouts
	// The select statement can be combined with a time.After channel to implement timeouts.
	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "Timeout"
	}()

	select {
	case msg := <-chan1:
		fmt.Print("Recieved ", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}
}
