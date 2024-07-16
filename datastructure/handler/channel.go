package handler

import "fmt"

/*
Channels are used to communicate between goroutines. They provide a way for one
goroutine to send data to another goroutine. Channels can be created using the make function.
*/
var count1 = 0

func ChannelInGo1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		count1++
		fmt.Println("count1", count1)
	}
	c <- sum // send sum to c
}

var count2 = 0

func ChannelInGo2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		count2++
		fmt.Println("count2", count2)
	}
	c <- sum // send sum to c
}
