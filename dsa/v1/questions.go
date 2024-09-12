package v1

import "fmt"

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
