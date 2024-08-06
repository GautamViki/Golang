package main

import (
	"linkedlist/handler"
)

func main() {
	list := handler.LinkedList{}
	list.CreateLinkedList(10)
	list.CreateLinkedList(20)
	list.CreateLinkedList(30)
	list.CreateLinkedList(40)
	list.CreateLinkedList(50)

	list.DisplayLinkedList()
}
