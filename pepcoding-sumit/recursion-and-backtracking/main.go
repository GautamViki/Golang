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
}
