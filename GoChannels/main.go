package main

import (
	"fmt"
	"gochannels/handler"
)

func main() {
	ch := handler.NewGoChannelHandler()
	chn := make(chan int, 10)
	go ch.FibonacciBufferedChannels(10, chn)
	go ch.FibonacciBufferedChannels(10, chn)
	for x := range chn {
		fmt.Println(x)
	}
}
