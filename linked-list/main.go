package main

import "linkedlist/handler"

func main() {
	head := handler.CreateLinkedList()
	handler.PrintAllNode(head)
	head = handler.InsertAtFirst(90, head)
	head = handler.InsertAtLast(100, head)
	head = handler.InsertAtGivenIdx(5, 3, head)
	head = handler.DeleteFromFirst(head)
	handler.DeleteFromLast(head)
	handler.DeleteFromGivenPosition(head, 3)
	handler.DeleteGivenKey(5, head)
	handler.AddNodeAsSorted(handler.Node{Value: 30}, head)
	handler.AddNodeAsSorted(handler.Node{Value: 30}, head)
	handler.RemoveDuplicateFromSorted(head)
}
