package main

import (
	"fmt"
	"recursion-backtracking/handler"
)

func main() {
	recursion := handler.NewRecursion()
	nums := []int{2, 4, 3, 8, 5, 2, 3, 5, 9, 6}
	firstOccurrence := recursion.FindFirstOccurrence(nums, 0, 30)
	fmt.Println("First Occurrence of a number : ", firstOccurrence)
	lastOccurrence := recursion.FindLastOccurrence(nums, 0, 61, -1)
	fmt.Println("Last Occurrence of a number : ", lastOccurrence)
	allOccurrence := recursion.FindAllOccurrence(nums, []int{}, 0, 2)
	fmt.Println("ALl Occurrence of a number : ", allOccurrence)
}
