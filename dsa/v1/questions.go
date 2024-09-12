package v1

import (
	"fmt"
	"math"
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
