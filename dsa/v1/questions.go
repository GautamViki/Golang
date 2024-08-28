package v1

import (
	"fmt"
)

func MaximumProductSubarray() {
	nums := []int{2, 2, 0, 3}
	// O(n^2)
	// max := nums[0]
	// pro := 1
	// for i := 0; i < len(nums); i++ {
	// 	for j := i; j < len(nums); j++ {
	// 		pro *= nums[j]
	// 		if pro > max {
	// 			max = pro
	// 		}
	// 	}
	// 	pro = 1
	// }
	// fmt.Println("dbhssssssssssssssssssssssss", max)

	// O(n)
	leftPro, rightPro := 1, 1
	left, right := 0, len(nums)-1
	maxPro := nums[0]
	for left < len(nums) || right >= 0 {
		leftPro *= nums[left]
		rightPro *= nums[right]
		maxPro = max(maxPro, max(leftPro, rightPro))
		if nums[left] == 0 {
			leftPro = 1
		}
		if nums[right] == 0 {
			rightPro = 1
		}
		left++
		right--
	}
	fmt.Println("Maximum Product Subarray", maxPro)
}

func MaxSumOfSubArray() {
	nums := []int{-2, -1}
	maxSum := nums[0]
	// Brutforce method O(n^2)
	// for i := 0; i < len(nums); i++ {
	// 	sum := 0
	// 	for j := i; j < len(nums); j++ {
	// 		sum += nums[j]
	// 		maxSum = max(maxSum, sum)
	// 	}
	// }
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		maxSum = max(maxSum, sum)
		if sum < 0 {
			sum = 0
		}
	}
	fmt.Println("Maximum Subarray", maxSum)
}

func LongestSubarrayConsistingOfUniqueElements() {
	// nums := []int{1, 2, 4, 4, 5, 6, 7, 8, 3, 4, 5, 3, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4}
	nums := []int{1, 2, 2, 1}
	result := []int{}
	tempMap := make(map[int]int)
	start := 0
	for end, num := range nums {
		if idx, found := tempMap[num]; found && idx >= start {
			start = idx + 1
		}
		tempMap[num] = end
		currentLength := end - start + 1
		if currentLength > len(result) {
			result = nums[start : end+1]
		}
	}
	fmt.Println("Longest Subarray consisting of unique elements from an Array", result)
}

func LongestUniqueSubstring() {
	s := "abba"
	maxSub := make(map[string]int)
	start := 0
	var result string
	for end, char := range s {
		if idx, found := maxSub[string(char)]; found && idx >= start {
			start = idx + 1
		}
		maxSub[string(char)] = end
		currentLength := end - start + 1
		if currentLength > len(result) {
			result = s[start : end+1]
		}
	}
	fmt.Println(result, "3. Longest Substring Without Repeating Characters", len(result))
}
