package questions

import (
	"fmt"
	"sort"
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
