package handler

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func CreateLinkedList() *Node {
	head := Node{Value: 10}
	node2 := Node{Value: 20}
	node3 := Node{Value: 30}
	node4 := Node{Value: 40}
	node5 := Node{Value: 50}

	// link all node
	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	return &head
}

func PrintAllNode(head *Node) {
	current := head
	count := 0
	for current != nil {
		count++
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println("\nTotal number of node in list", count)
}

func InsertAtFirst(val int, head *Node) *Node {
	fmt.Println("\nInserted at first node", val)
	current := head
	head = &Node{Value: val, Next: current}
	PrintAllNode(head)
	return head
}

func InsertAtLast(val int, head *Node) *Node {
	fmt.Println("\nInserted at last node", val)
	last := Node{Value: val}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &last
	PrintAllNode(head)
	return head
}

func InsertAtGivenIdx(val, idx int, head *Node) *Node {
	fmt.Println("\nInserted at given position", idx, "node", val)
	newNode := Node{Value: val}
	count := 1
	current := head
	for count < idx-1 {
		count++
		current = current.Next
	}
	if idx == 1 {
		newNode.Next = head
		head = &newNode
	} else {
		newNode.Next = current.Next
		current.Next = &newNode
	}
	PrintAllNode(head)
	return head
}

func DeleteFromFirst(head *Node) *Node {
	fmt.Println("\nDeleted from first node")
	head = head.Next
	PrintAllNode(head)
	return head
}

func DeleteFromLast(head *Node) {
	fmt.Println("\nDeleted from last node")
	curr := head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil
	PrintAllNode(head)
}

func DeleteFromGivenPosition(head *Node, idx int) {
	fmt.Println("\nDeleted from given position", idx)
	curr := head
	count := 1
	if idx == 1 {
		head = head.Next
	} else {
		for count < idx-1 {
			curr = curr.Next
			count++
		}
	}
	curr.Next = curr.Next.Next
	PrintAllNode(head)
}

func DeleteGivenKey(key int, head *Node) {
	fmt.Println("\nDeleted given key", key)
	curr := head
	pre := &Node{}
	for curr.Next != nil {
		pre = curr
		curr = curr.Next
		if curr.Value == key {
			pre.Next = curr.Next
			break
		}
	}
	PrintAllNode(head)
}

func AddNodeAsSorted(newNode Node, head *Node) {
	fmt.Println("\nInserted as sorted list Node ", newNode)
	curr := head
	pre := &Node{}
	for curr != nil && curr.Value <= newNode.Value {
		pre = curr
		curr = curr.Next
	}
	newNode.Next = curr
	pre.Next = &newNode
	PrintAllNode(head)
}

func RemoveDuplicateFromSorted(head *Node) {
	fmt.Println("\nRemove duplicate from sorted list")
	curr := head
	for curr.Next != nil && curr != nil {
		if curr.Value == curr.Next.Value {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	PrintAllNode(head)
}











