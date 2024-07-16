package main

import (
	"datastructure/handler"
	"fmt"
)

func main() {
	// handler.AarrayDeclaration()
	// fmt.Println("==============================================")
	// handler.SliceDeclaration()
	// fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	// handler.MapHandler()
	// fmt.Println("##############################################")
	// handler.RuneHandler()
	// fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	// handler.VariadicHandler(11)
	// // anonymous function
	// func(a, b int) {
	// 	fmt.Println("a", a, "b", b)
	// }(1, 2)

	// closures
	// sum := handler.Sum(10)
	// fmt.Println(sum(5))
	// fmt.Println(sum(10))

	// constructer
	// rect := handler.NewRectangle()
	// rect.AreaByValue()
	// rect.AreaByPointer()
	// rect.AreaByValue()
	// fmt.Println("Height",rect.Height)
	// fmt.Println("Width ",rect.Width)

	// linked list
	// list := handler.NewLinkedList()
	// list.InsertNodeAtFirst(8)

	// GoRoutine Function
	// fmt.Println("Goroutine Concept")
	// go handler.GoRoutineSay("*********************")
	// handler.GoRoutineSay("#######")

	// channel concepts
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go handler.ChannelInGo1(s[:len(s)/2], c)
	go handler.ChannelInGo2(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, "Channel Concepts ", x+y)

	// Buffered Channels
	/*
		Channels can be buffered. A buffered channel has a capacity and does not block until the buffer is full.
	*/
	fmt.Println("Buffered Channels")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	// ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
}
