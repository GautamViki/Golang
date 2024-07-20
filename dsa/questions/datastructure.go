package questions

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type DataStructure struct{}

func NewDataStructure() *DataStructure {
	return &DataStructure{}
}

func (l *DataStructure) LongestCommonPrefix() string {
	testSlice := []string{"flower", "flow", "flight"} //output ->fl
	// testSlice := []string{"dog", "racecar", "car"} //output ->""
	sort.Strings(testSlice)
	var str string
	for idx, val := range testSlice[0] {
		if (val) != rune(testSlice[len(testSlice)-1][idx]) {
			return str
		} else {
			str += string(val)
		}
	}
	return str
}

func (ds *DataStructure) ValidParentheses() bool {
	s := "([}}])"
	arr := []string{}
	if len(s)%2 == 1 {
		return false
	}
	for _, val := range s {
		if string(val) == "(" || string(val) == "{" || string(val) == "[" {
			arr = append(arr, string(val))
		} else if len(arr) > 0 && string(val) == ")" && arr[len(arr)-1] == "(" {
			arr = arr[:len(arr)-1]
		} else if len(arr) > 0 && string(val) == "]" && arr[len(arr)-1] == "[" {
			arr = arr[:len(arr)-1]
		} else if len(arr) > 0 && string(val) == "}" && arr[len(arr)-1] == "{" {
			arr = arr[:len(arr)-1]
		} else {
			return false
		}
	}
	return len(arr) == 0
}

func (ds *DataStructure) MergeTwoSortedLists() []int {
	list1 := []int{1, 9}
	list2 := []int{1, 3, 4}
	var result []int
	var i, j int
	for len(list1) > j || len(list2) > i {
		if len(list1) > j && len(list2) > i {
			if list1[j] == list2[i] {
				result = append(result, list1[j], list2[i])
				i++
				j++
			} else if list1[j] < list2[i] {
				result = append(result, list1[j])
				j++
			} else if list1[j] > list2[i] {
				result = append(result, list2[i])
				i++
			}
		} else if len(list1) > j {
			result = append(result, list1[j])
			j++
		} else if len(list2) > i {
			result = append(result, list2[i])
			i++
		}
	}
	return result
}

func (ds *DataStructure) LengthOfLastWord() {
	str := " cxx  "
	n := len(str) - 1
	count := 0
	for n >= 0 {
		if str[n] == ' ' && count == 0 {
			n--
		} else if str[n] != ' ' {
			count++
			n--
		} else if str[n] == ' ' && count != 0 {
			break
		}
	}
	fmt.Println("LengthOfLastWord ", count)
}

func (ds *DataStructure) SquareRoot() {
	x := 10
	mid := 0

	if x == 0 {
		mid = 0
	} else if x < 3 {
		mid = 1
	} else if x < 8 {
		mid = 2
	}

	n := x / 2
	i := 1
	j := n
	for i < j {
		mid = (i + j) / 2
		if mid*mid < x {
			i = mid + 1
		} else if mid*mid > x {
			j = mid - 1
		} else {
			break
		}
	}
	fmt.Println("SquareRoot ", j)
}

var cache = make(map[int]int)

func (ds *DataStructure) ClimbingStairs(n int) int {
	if n < 3 {
		return n
	}
	if val, ok := cache[n]; ok {
		return val
	}
	cache[n] = ds.ClimbingStairs(n-1) + ds.ClimbingStairs(n-2)
	return cache[n]
}

func (ds *DataStructure) ThreeConsecutiveOdds() bool {
	arr := []int{1, 2, 3}
	for i := 0; i <= len(arr)-3; i++ {
		if arr[i]%2 == 0 {
			continue
		}
		if arr[i]%2 == 1 && arr[i+1]%2 == 1 && arr[i+2]%2 == 1 {
			return true
		}
	}
	return false
}

func (ds *DataStructure) ValidPalindrome() bool {
	s := "0P0,"
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}
	s = strings.Map(f, s)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func (ds *DataStructure) MajorityElement() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	// hash := make(map[int]int)
	// for i := 0; i < len(nums); i++ {
	// 	hash[nums[i]] = hash[nums[i]] + 1
	// }
	// count := hash[nums[0]]
	// max := nums[0]
	// for val := range hash {
	// 	if hash[val] > count {
	// 		max = val
	// 	}
	// }

	var ans int
	var count int
	for _, val := range nums {
		if count == 0 {
			ans = val
		}
		if ans == val {
			count++
		} else {
			count--
		}
	}
	fmt.Println("MajorityElement", ans)
}

func (ds *DataStructure) AddBinary() {
	num1 := "111"
	num2 := "111"
	i := len(num1) - 1
	j := len(num2) - 1
	carry := '0'
	result := ""
	for i == j && i >= 0 {
		if num1[i] == num2[j] && num1[i] == '1' {
			if carry != '0' {
				result += string('1')
			} else {
				result += string('0')
			}
			carry = '1'
		} else if num1[i] == num2[j] && num1[i] == '0' {
			if carry != '0' {
				result += string('1')
			} else {
				result += string('0')
			}
			carry = '0'
		} else if num1[i] != num2[j] {
			if carry != '0' {
				result += string('0')
				carry = '1'
			} else {
				result += string('1')
				carry = '0'
			}
		}
		i--
		j--
	}
	if carry == '1' {
		result += string('1')
	}
	
	fmt.Println("Binary ", result)
}
