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
}
