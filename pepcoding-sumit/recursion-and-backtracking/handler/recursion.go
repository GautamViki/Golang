package handler

import "fmt"

func PrintDecreasing(n int) {
	if n < 1 {
		return
	}
	fmt.Print(" ", n)
	PrintDecreasing(n - 1)
}

func PrintIncreasing(n int) {
	if n == 0 {
		return
	}
	PrintIncreasing(n - 1)
	fmt.Print(" ", n)
}

func PrintIncreasingDecreasing(n int) {
	if n == 0 {
		return
	}
	fmt.Print(" ", n)
	PrintIncreasingDecreasing(n - 1)
	fmt.Print(" ", n)
}

func FactorialOf_N(n int) int {
	if n <= 1 {
		return 1
	}
	return n * FactorialOf_N(n-1)
}

func PowerOf_X(x, n int) int {
	if n == 0 {
		return 1
	}
	return x * PowerOf_X(x, n-1)
}

func PowerOf_X_in_Log(x, n int) int {
	if n == 0 {
		return 1
	}
	powerN2 := PowerOf_X_in_Log(x, n/2)
	if n%2 == 1 {
		return x * powerN2 * powerN2
	}
	return powerN2 * powerN2
}

func TowerOfHanoi(n int, S, D, H string) {
	if n == 0 {
		return
	}
	TowerOfHanoi(n-1, S, H, D)
	fmt.Printf("disc %d Source %s ==> destination %s\n", n, S, D)
	TowerOfHanoi(n-1, H, D, S)
}

func DisplayArray(nums []int, idx int) {
	if idx == len(nums) {
		return
	}
	fmt.Print(nums[idx], " ")
	DisplayArray(nums, idx+1)
}
