package handler

import "fmt"

// type listHandler struct{}

// func NewListHandler() *LinkedList {
// 	return &LinkedList{}
// }

type Node struct {
	Val  int
	Next *Node
}
type LinkedList struct {
	Head *Node
}

func (list *LinkedList) CreateLinkedList(value int) {
	node := Node{Val: value}
	if list.Head == nil {
		list.Head = &node
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &node
	}
}

func (list *LinkedList) DisplayLinkedList() {
	fmt.Println("Head Node : ",list.Head)
	fmt.Print("List Node : ")
	current := list.Head
	for current != nil {
		fmt.Print(current.Val," ")
		current = current.Next
	}
}
