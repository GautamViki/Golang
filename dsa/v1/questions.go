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

