package main

import (
	"fmt"
	recursion "recursion-backtracking/handler"
)

func main() {
	recursion.PrintDecreasing(10)
	fmt.Println()
	recursion.PrintIncreasing(10)
	fmt.Println()
	recursion.PrintIncreasingDecreasing(10)
	fmt.Println("\nFactorial of N: ", recursion.FactorialOf_N(10))
	fmt.Println("Power of x to n: ", recursion.PowerOf_X(2, 20))
	fmt.Println("Power of x to n in Logerithemic: ", recursion.PowerOf_X_in_Log(2, 20))
	recursion.TowerOfHanoi(3, "S", "D", "H")
	recursion.DisplayArray([]int{1, 2, 3, 4, 5, 6, 7, 8}, 0)
	fmt.Println()
	recursion.DisplayArrayReverse([]int{1, 2, 3, 4, 5, 6, 7, 8}, 0)
	fmt.Println("\nMaximum in Array : ", recursion.Max_In_Array([]int{1, 2, 3, 40, 5, 6, 7, 8}, 0, 1))
	fmt.Println("First occurence in array index at: ", recursion.FirstOccurenceInArray([]int{1, 8, 3, 4, 5, 6, 7, 8}, 0, 80))
	fmt.Println("Last occurence in array index at: ", recursion.LastOccurrenceInArray([]int{1, 8, 3, 4, 5, 6, 7, 8}, 0, 4, -1))
	fmt.Println("All occurence of array index: ", recursion.OccurenceOfAllIndeces([]int{1, 8, 3, 4, 8, 6, 7, 8}, []int{}, 0, 8))
}
