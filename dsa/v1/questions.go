package v1

import (
	"fmt"
	"math"
	"unicode"
)

func TwoSum_1() {
	nums, target := []int{3, 2, 4}, 6
	hashMap := make(map[int]int)
	for idx, num := range nums {
		currentTarget := target - num
		if i, ok := hashMap[currentTarget]; ok {
			fmt.Println("1. Two Sum ", []int{i, idx})
			return
		}
		hashMap[num] = idx
	}
	fmt.Println("1. Two Sum ", []int{})
}

func ValidParentheses_20() {
	s := "()"
	stack := []rune{}
	if len(s)%2 == 1 {
		fmt.Println("20. Valid Parentheses ", false)
		return
	}
	for _, str := range s {
		n := len(stack)
		if str == '(' || str == '[' || str == '{' {
			stack = append(stack, str)
		} else if n > 0 && str == ')' && stack[n-1] == '(' {
			stack = stack[:n-1]
		} else if n > 0 && str == '}' && stack[n-1] == '{' {
			stack = stack[:n-1]
		} else if n > 0 && str == ']' && stack[n-1] == '[' {
			stack = stack[:n-1]
		} else {
			fmt.Println("20. Valid Parentheses ", false)
			return
		}
	}
	fmt.Println("20. Valid Parentheses ", len(stack) == 0)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoSortedLists_21(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else if list2.Val <= list1.Val {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}
	return head.Next
}

func BestTimeToBuyAndSellStock() {
	prices := []int{7, 6, 4, 3, 1}
	min := prices[0]
	max := math.MinInt32
	i := 0
	for i < len(prices) {
		currentMax := prices[i] - min
		if max < currentMax {
			max = currentMax
		}
		if min > prices[i] {
			min = prices[i]
		}
		i++
	}
	fmt.Println("121. Best Time to Buy and Sell Stock", max)
}

func ValidPalindrome() {
	s := "2%2"
	i, j := 0, len(s)-1
	for i < j {
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			i++
		} else if !unicode.IsLetter(rune(s[j])) && !unicode.IsNumber(rune(s[j])) {
			j--
		} else {
			if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
				fmt.Println("125. Valid Palindrome", false)
				return
			}
			i++
			j--
		}
	}
	fmt.Println("125. Valid Palindrome", true)
}

func ValidAnagram() {
	s, t := "anaqram", "nagaram"
	if len(s) != len(t) {
		fmt.Println("242. Valid Anagram", false)
		return
	}
	sSet := make(map[rune]int)
	tSet := make(map[rune]int)
	for idx, char := range s {
		sSet[char] += 1
		tSet[rune(t[idx])] += 1
	}
	if len(sSet) != len(tSet) {
		fmt.Println("242. Valid Anagram", false)
		return
	}
	for key, val := range sSet {
		if val != tSet[key] {
			fmt.Println("242. Valid Anagram", false)
			return
		}
	}
	fmt.Println("242. Valid Anagram", true)
}

func BinarySearch() {
	nums, target := []int{5}, 5
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high + low) / 2
		if nums[mid] == target {
			fmt.Println("704. Binary Search", mid)
			return
		} else if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		}
	}
	fmt.Println("704. Binary Search", -1)
}

func SingleNumber() {
	arr := []int{2, 2, 3, 1, 1}
	var uniqueEle int
	for _, val := range arr {
		uniqueEle ^= val
	}
	fmt.Println("136. Single Number", uniqueEle)
}

func MajorityElement() {
	nums := []int{3, 2, 3}
	res, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			res = num
		}
		if res == num {
			count++
		} else {
			count--
		}
	}
	fmt.Println("169. Majority Element ", res)
}

func MissingNumber() {
	nums := []int{3, 0, 1}
	// n, sum := len(nums), 0
	// totalSum := (n + 1) * n / 2
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println("268. Missing Number", totalSum-sum)
	xor := 0
	for i := 0; i <= len(nums); i++ {
		xor ^= i
	}
	for _, num := range nums {
		xor ^= num
	}
	fmt.Println("268. Missing Number", xor)
}

func ReverseString() {
	s := []string{"h", "e", "l", "l", "o"}
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	fmt.Println("344. Reverse String", s)
}

func MaximumSubarray() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	sum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum = max(nums[i], nums[i]+sum)
		maxSum = max(maxSum, sum)
	}
	fmt.Println("53. Maximum Subarray", maxSum)
}
