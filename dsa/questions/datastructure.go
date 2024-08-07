package questions

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
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
	a := "1010"
	b := "1011"
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	result := ""
	sum := 0
	temp := 0
	for i >= 0 || j >= 0 {
		sum = carry
		if i >= 0 {
			temp, _ = strconv.Atoi(string(a[i]))
			sum += temp
		}
		if j >= 0 {
			temp, _ = strconv.Atoi(string(b[j]))
			sum += temp
		}
		result += strconv.Itoa(sum % 2)
		carry = sum / 2
		i--
		j--
	}
	if carry > 0 {
		result += strconv.Itoa(carry)
	}
	runes := []rune(result)
	slices.Reverse(runes)
	fmt.Println("Add Bunary ", string(runes))
}

func (ds *DataStructure) RemoveDuplicate() {
	arr := []int{1, 1, 1, 5, 7, 9, 1, 2, 3, 3, 3, 4, 4, 5, 5, 5, 6}
	slices := []int{}
	hash := make(map[int]struct{})
	for _, ele := range arr {
		if _, ok := hash[ele]; !ok {
			hash[ele] = struct{}{}
			slices = append(slices, ele)
		}
	}
	fmt.Println("Remove Duplicate Elements", slices)
}

func (ds *DataStructure) ProductExceptSelf() {
	arr := [4]int{1, 2, 3, 4}
	slicesTemp := make([]int, len(arr))
	multiple := 1
	for i := 0; i < len(arr); i++ {
		slicesTemp[i] = multiple
		multiple *= arr[i]
	}
	multiple = 1
	for i := len(arr) - 1; i >= 0; i-- {
		slicesTemp[i] *= multiple
		multiple *= arr[i]
	}
	fmt.Println("Product of Array Except Self", slicesTemp)
}

func (ds *DataStructure) MergeTwoSortedArray() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
	fmt.Println("Merge Two Sorted Array ", nums1)
}

func (ds *DataStructure) SingleNumber() {
	arr := []int{4, 1, 2, 1, 2}
	var uniqueEle int
	// map method
	// myMap := map[int]int{}
	// for _, val := range arr {
	// 	myMap[val] += 1
	// }
	// for key, val := range myMap {
	// 	if val == 1 {
	// 		uniqueEle = key
	// 		break
	// 	}
	// }
	for _, val := range arr {
		uniqueEle = uniqueEle ^ val
	}
	fmt.Println("Single number element ", uniqueEle)
}

func (ds *DataStructure) ContainsDuplicate() {
	// arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	arr := []int{1, 2, 3, 4}
	myMap := map[int]int{}
	var contain bool
	for _, val := range arr {
		if _, ok := myMap[val]; ok {
			contain = ok
			break
		}
		myMap[val] += 1
	}
	fmt.Println("Contains Duplicate", contain)
}

func (ds *DataStructure) RemoveDuplicateFromSortedArray() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	i, k, count := 1, 1, 1
	for i < len(nums) {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
			count++
		}
		i++
	}
	fmt.Println("Remove Duplicate From Sorted Array ", count)
}

func (ds *DataStructure) ExcelSheetColumnTitle() {
	columnNumber := 701
	// A:65, B:62
	fmt.Println('B')
	n := 0
	num := 1
	for num+26 <= columnNumber {
		n++
		num = 26 * num
	}
	// fmt.Println("value of n", n)
}

func (ds *DataStructure) HappyNumber(n int) bool {
	happyMap := make(map[int]int)
	for {
		if n == 1 {
			return true
		}
		str := strconv.Itoa(n)
		sum := 0
		for _, digit := range str {
			i, _ := strconv.Atoi(string(digit))
			sum += i * i
		}
		if _, found := happyMap[sum]; found {
			return false
		}
		happyMap[sum] = 1
		n = sum
	}
}

func (ds *DataStructure) ReverseInteger(n int) int {
	var result int
	p := 10
	signed := false
	if n < 0 {
		signed = true
		n *= (-1)
	}
	for n > 0 {
		r := n % 10
		result = result*p + r
		n = n / 10
	}
	if signed {
		result = result * (-1)
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

func (ds *DataStructure) ThreeSum() {
	nums := []int{0, 0, 0}
	sort.Ints(nums)
	fmt.Println(nums)
	var result [][]int
	if len(nums) < 3 {
		return
	}
	for i := 0; i < len(nums)-2; i++ {
		x := nums[i]
		j, k := i+1, len(nums)-1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < k {
			tempResult := x + nums[j] + nums[k]
			if tempResult == 0 {
				result = append(result, []int{x, nums[j], nums[k]})
				for j < len(nums)-1 && nums[j] == nums[j+1] {
					j++
				}
				for k > 0 && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if tempResult > 0 {
				k--
			} else if tempResult < 0 {
				j++
			}
		}
	}
	fmt.Println("3 Sum : ", result)
}

func (ds *DataStructure) ArraySetInGO() {
	set := make(map[[3]int]struct{})
	set[[3]int{1, 2, 3}] = struct{}{}
	set[[3]int{2, 3, 4}] = struct{}{}
	set[[3]int{3, 4, 5}] = struct{}{}
	if _, found := set[[3]int{1, 2, 31}]; found {
		fmt.Println("Found in Set : ", found)
	} else {
		fmt.Println("Found in Set : ", found)
	}
	delete(set, [3]int{1, 2, 3})
	fmt.Println("set : ", set)
}

// func (ds *DataStructure) SliceSetInGo() {
// 	sliceSet := make(map[string]struct{})

// }

// func sliceToString(slice []int) string {
// 	return strings.Trim(strings.Replace(fmt.Sprint(slice), " ", ",", -1), "[]")
// }

func (ds *DataStructure) LongestSubstringWithoutRepeatingCharacters() {
	s := "dvdf"
	var res int
	sub := make([]string, 0)
	for _, char := range s {
		if !strings.Contains(strings.Join(sub, ""), string(char)) {
			sub = append(sub, string(char))
			res = max(res, len(sub))
		} else {
			cut := strings.Index(strings.Join(sub, ""), string(char))
			sub = append(sub[cut+1:], string(char))
		}
	}
	fmt.Println("Longest Substring Without Repeating Characters", res)
}

func (ds *DataStructure) StringToInteger() {
	s := "21474836460"
	signed := false
	var result string
	count := 0
	for idx, char := range s {
		if unicode.IsDigit(char) {
			fmt.Println("digit ", string(char))
			result += string(char)
		} else if count == idx {
			if string(char) == " " {
				count++
			} else if string(char) == "-" {
				signed = true
			} else if string(char) != "+" {
				break
			}
		} else {
			break
		}
	}
	if signed {
		result = "-" + result
	}
	num, _ := strconv.Atoi(result)
	if num > math.MaxInt32 {
		num = math.MaxInt32
	} else if num < math.MinInt32 {
		num = math.MinInt32
	}

	fmt.Println("8 .String To Integer", num)
}

func (ds *DataStructure) MedianOfTwoSortedArrays() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 7}
	n := len(nums1)
	m := len(nums2)
	var result float64
	if n > 0 && n%2 == 0 {
		result = float64(nums1[n/2]+nums1[n/2-1]) / 2
	}
	if m > 0 && m%2 == 0 {
		result += float64(nums2[m/2]+nums2[m/2-1]) / 2
	}
	if n > 0 && n%2 == 1 {
		result += float64(nums1[n/2])
	}
	if m > 0 && m%2 == 1 {
		result += float64(nums2[m/2])
	}
	if n > 0 && m > 0 {
		result = result / 2
	}

	roundedFloat := math.Round(result*100000) / 100000
	fmt.Println("8888888888888888888888888888888888 ", roundedFloat)
}

// func (ds *DataStructure)ReverseBits(){
// 	num :=uint32(1111110001111111111111100)
// 	var answer uint32
//     for i:=0;i<32;i++{
//         answer=answer | (num&1)<<(31-i)
//         num=num>>1
//     }
//     fmt.Println("000000000000000000000000000",answer)
// }

func (ds *DataStructure) KeyBoardRow_500() {
	strArr := []string{"omk"}
	row1, row2, row3 := "qwertyuiop", "asdfghjkl", "zxcvbnm"
	result := []string{}
	for _, word := range strArr {
		i, j := 0, len(word)-1
		found := true
		for i <= j {
			if strings.Contains(row1, strings.ToLower(string(word[i]))) &&
				strings.Contains(row1, strings.ToLower(string(word[j]))) {
				i++
				j--
			} else {
				found = false
				break
			}
		}
		if !found {
			found = true
			i, j := 0, len(word)-1
			for i <= j {
				if strings.Contains(row2, strings.ToLower(string(word[i]))) &&
					strings.Contains(row2, strings.ToLower(string(word[j]))) {
					i++
					j--
				} else {
					found = false
					break
				}
			}
		}
		if !found {
			found = true
			i, j := 0, len(word)-1
			for i <= j {
				if strings.Contains(row3, strings.ToLower(string(word[i]))) &&
					strings.Contains(row3, strings.ToLower(string(word[j]))) {
					i++
					j--
				} else {
					found = false
					break
				}
			}
		}
		if found {
			result = append(result, word)
		}
	}
	fmt.Println("Keyboard Row : ", result)
}
