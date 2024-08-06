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

func (list *LinkedList) CreateLinkedList() {
	value := 10
	node := Node{Val: value}
	list.Head = &node
	fmt.Println("first Node : ", list.Head)
}
