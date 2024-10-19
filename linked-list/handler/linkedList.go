package handler

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyListWithRandomPointer_138(head *Node) *Node {
	if head == nil {
		return nil
	}
	hashMap := make(map[*Node]*Node)
	currentNode := head
	var newHead *Node
	var preNode *Node
	for currentNode != nil {
		temp := Node{
			Val: currentNode.Val,
		}
		if newHead == nil {
			newHead = &temp
			preNode = newHead
		} else {
			preNode.Next = &temp
		}
		hashMap[currentNode] = preNode
		currentNode = currentNode.Next
	}
	currentNode = head
	newCuurent := newHead
	for currentNode != nil {
		if currentNode.Random == nil {
			newCuurent.Random = nil
		} else {
			newCuurent.Random = hashMap[currentNode]
		}
		newCuurent = newCuurent.Next
		currentNode = currentNode.Next
	}
	return newHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseLinkedList_II(head *ListNode, left, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := ListNode{}
	dummy.Next = head
	pre := &dummy

	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	curr := pre.Next
	for i := 1; i <= right-left; i++ {
		temp := pre.Next
		pre.Next = curr.Next
		curr.Next = curr.Next.Next
		pre.Next.Next = temp
	}
	return dummy.Next
}
