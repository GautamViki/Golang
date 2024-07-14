package handler

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) InsertNodeAtFirst(val int) {
	newNode1 := Node{Value: val}
	newNode1.Next = l.Head
	l.Head = &newNode1

	newNode2 := Node{Value: 10}
	newNode2.Next = l.Head
	l.Head = &newNode2
	fmt.Println(newNode2)
}
