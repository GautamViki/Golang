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

func DisplayArrayReverse(nums []int, idx int) {
	if idx == len(nums) {
		return
	}
	DisplayArrayReverse(nums, idx+1)
	fmt.Print(nums[idx], " ")
}

func Max_In_Array(nums []int, idx int, maxNum int) int {
	if idx == len(nums) {
		return maxNum
	}
	maxNum = max(maxNum, nums[idx])
	return Max_In_Array(nums, idx+1, maxNum)
}

func FirstOccurenceInArray(nums []int, idx, target int) int {
	if idx == len(nums) {
		return -1
	}
	if target == nums[idx] {
		return idx
	}
	return FirstOccurenceInArray(nums, idx+1, target)
}

func LastOccurrenceInArray(nums []int, idx, target, res int) int {
	if idx == len(nums) {
		return res
	}
	if target == nums[idx] {
		res = idx
	}
	return LastOccurrenceInArray(nums, idx+1, target, res)
}

func OccurenceOfAllIndeces(nums []int, res []int, idx, target int) []int {
	if idx == len(nums) {
		return res
	}
	if nums[idx] == target {
		res = append(res, idx)
	}
	return OccurenceOfAllIndeces(nums, res, idx+1, target)
}

func FindSubSequence(str string, idx int, current string) {
	if idx == len(str) {
		if current != "" {
			fmt.Print( current," ")
		}
		return
	}
	FindSubSequence(str, idx+1, current+string(str[idx]))
	FindSubSequence(str, idx+1, current)
}
