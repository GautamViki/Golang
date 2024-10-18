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
		newCuurent=newCuurent.Next
		currentNode=currentNode.Next
	}
	return newHead
}
