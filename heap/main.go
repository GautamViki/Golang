package main

import (
	"fmt"
	"heap/handler"
)

func main() {
	heap := handler.NewMinHeap()
	heap.Insert(10)
	heap.Insert(20)
	heap.Insert(15)
	heap.Insert(5)
	heap.Insert(30)
	heap.Insert(25)

	fmt.Println("Min-Heap array:")
	heap.Display()

	min, _ := heap.ExtractMin()
	fmt.Println("Extracted Min:", min)
	fmt.Println("Min-Heap after extraction:")
	heap.Display()
}
