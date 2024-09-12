package linkedlist

type ListNode struct {
	Value int
	Next  *ListNode
}

func CreateLinkedList(val int) {
	head := &ListNode{}
	currentNode := head
	currentNode.Value = val
	currentNode.Next = nil
}
