package v1

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"
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

func BestTimeToBuyAndStock_II() {
	prices := []int{7, 6, 4, 3, 1}
	maxS := 0
	for i := 1; i < len(prices); i++ {
		curMax := prices[i] - prices[i-1]
		if curMax > 0 {
			maxS += curMax
		}
	}
	fmt.Println("Best Time To Buy And Stock_II", maxS)
}

func HouseRobber() {
	nums := []int{2, 7}
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1 // Initialize with -1 to indicate uncomputed results
	}
	sum := MaxRobber(nums, 0, memo)
	fmt.Println("198. House Robber", sum)
}

func MaxRobber(nums []int, idx int, memo []int) int {
	if idx >= len(nums) {
		return 0
	}
	if memo[idx] != -1 {
		return memo[idx]
	}
	steal := nums[idx] + MaxRobber(nums, idx+2, memo)
	skip := MaxRobber(nums, idx+1, memo)
	memo[idx] = max(steal, skip)
	return memo[idx]
}

func NumberOf_1_Bits() {
	n := 11
	count := 0
	for n != 0 {
		count += int(n & 1)
		n >>= 1
	}
	fmt.Println("191. Number of 1 Bits", count)
}

func MergeIntervals() {
	intervals := [][]int{{1, 80}, {4, 6}, {8, 10}, {15, 181}}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		x := result[len(result)-1][1]
		y := intervals[i][0]
		if x >= y {
			if x < intervals[i][1] {
				result[len(result)-1][1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}
	fmt.Println("56. Merge Intervals", result)
}

func SetMatrixZeroes() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	zeroIdx := [][]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				zeroIdx = append(zeroIdx, []int{i, j})
			}
		}
	}
	for z := 0; z < len(zeroIdx); z++ {
		i, j := zeroIdx[z][0], zeroIdx[z][1]
		for idx := 0; idx < len(matrix); idx++ {
			matrix[idx][j] = 0
		}
		for idx := 0; idx < len(matrix[i]); idx++ {
			matrix[i][idx] = 0
		}
	}
	fmt.Println("73. Set Matrix Zeroes", matrix)
}

func SpiralMatrix() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	n, row, col := len(matrix)*len(matrix[0]), len(matrix), len(matrix[0])
	i, j, res := 0, 0, []int{}
	count, t := 0, 0

	for count < n {
		for j < col-t && count < n {
			res = append(res, matrix[i][j])
			count++
			j++
		}
		j--
		i++
		for i < row-t && count < n {
			res = append(res, matrix[i][j])
			i++
			count++
		}
		j--
		i--
		for j >= t && count < n {
			res = append(res, matrix[i][j])
			j--
			count++
		}
		i--
		j++
		t++
		for i >= t && count < n {
			res = append(res, matrix[i][j])
			i--
			count++
		}
		j++
		i++
	}
	fmt.Println("54. Spiral Matrix", res)
}

func Sum_3() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sort.Ints(nums)
	var result [][]int
	if len(nums) < 3 {
		fmt.Println("15. 3Sum", result)
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
	fmt.Println("15. 3Sum", result)
}

func ContainerWithMostWater() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	left, right, mMax := 0, len(height)-1, 0
	for left < right {
		var minHeight int
		if height[left] < height[right] {
			minHeight = height[left]
		} else {
			minHeight = height[right]
		}
		if mMax < minHeight*(right-left) {
			mMax = minHeight * (right - left)
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	fmt.Println("11. Container With Most Water", mMax)
}

func GroupAnagrams() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	sortStr := [][]string{}
	hash := map[string][]string{}
	for _, str := range strs {
		newStr := []rune(str)
		slices.Sort(newStr)
		hash[string(newStr)] = append(hash[string(newStr)], str)
	}
	for _, val := range hash {
		sortStr = append(sortStr, val)
	}
	fmt.Println("49. Group Anagrams", sortStr)
}

func GroupAnagrams_1() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	sortStr := [][]string{}
	hash := map[string][]string{}
	for _, str := range strs {
		word := generateWord(str)
		hash[word] = append(hash[word], str)
	}
	for _, val := range hash {
		sortStr = append(sortStr, val)
	}
	fmt.Println("49. Group Anagrams_1", sortStr)

}
func generateWord(str string) string {
	arr := [26]int{}
	for _, ch := range str {
		arr[ch-'a'] += 1
	}
	newStr := ""
	for idx, num := range arr {
		newStr += strings.Repeat(string(idx+'a'), num)
	}
	return newStr
}

func LongestPalindromicSubstring() {
	str := "babadaba"
	maxPal, maxStr := 0, ""
	var memo [1001][1001]bool
	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {
			if checkPalindrome(str, i, j, &memo) && j-i+1 > maxPal {
				maxPal = j - i + 1
				maxStr = str[i : j+1]
			}
		}
	}
	fmt.Println("5. Longest Palindromic Substring", maxStr, maxPal)
}
func checkPalindrome(str string, i, j int, memo *[1001][1001]bool) bool {
	if memo[i][j] {
		return memo[i][j]
	}
	if i >= j {
		memo[i][j] = true
		return memo[i][j]
	}

	if str[i] == str[j] {
		memo[i][j] = checkPalindrome(str, i+1, j-1, memo)
		return memo[i][j]
	}
	return false
}

func MinimumWindowSubstring() {
	s, t := "a", "b"
	mapOfT := make(map[rune]int)
	minLen := math.MaxInt32
	for _, str := range t {
		mapOfT[str]++
	}
	i, j, count, start := 0, 0, len(t), 0
	for j < len(s) {
		if mapOfT[rune(s[j])] > 0 && count > 0 {
			mapOfT[rune(s[j])]--
			count--
		} else {
			mapOfT[rune(s[j])]--
		}
		for count == 0 {
			if minLen > j-i+1 {
				minLen = j - i + 1
				start = i
			}
			mapOfT[rune(s[i])]++
			if mapOfT[rune(s[i])] > 0 {
				count++
			}
			i++
		}
		j++
	}
	if minLen == math.MaxInt32 {
		fmt.Println("76. Minimum Window Substring", "")
		return
	} else {
		fmt.Println("76. Minimum Window Substring", s[start:start+minLen])
	}
}

func ContainsDuplicate_II() {
	nums, k := []int{1, 0, 1, 1}, 1
	i, j := 0, 0
	set := make(map[int]struct{})
	for j < len(nums) {
		if k < j-i {
			delete(set, nums[i])
			i++
		}
		if _, ok := set[nums[j]]; ok && j-i <= k {
			fmt.Println("219. Contains Duplicate II", true)
			return
		}
		set[nums[j]] = struct{}{}
		j++

	}
	fmt.Println("219. Contains Duplicate II", false)
}

func RemoveDuplicatesfromSortedArray_II() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	i, j, count, max := 1, 0, 1, 1
	for i < len(nums) {
		if nums[i] == nums[j] && count < 2 {
			nums[j+1] = nums[i]
			max++
			j++
			count++
		} else if nums[i] != nums[j] {
			nums[j+1] = nums[i]
			max++
			j++
			count = 1
		}
		i++
	}
	fmt.Println("80. Remove Duplicates from Sorted Array II", max, nums)
}

func RotateArray() {
	nums, k := []int{1, 2, 3, 4, 5, 6, 7}, 4
	k = k % len(nums)
	n := len(nums) - k
	reverse(nums, 0, n-1)
	reverse(nums, n, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
	fmt.Println("189. Rotate Array", nums)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func JumpGame() {
	nums := []int{2, 0, 0}
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev == 0 {
			fmt.Println("55. Jump Game", false)
			return
		}
		prev = max(prev-1, nums[i])
	}
	fmt.Println("55. Jump Game", true)
}

func JumpGame_II() {
	nums := []int{2, 3, 1, 1, 4}
	totalJump, coverage, lastJumpIdx := 0, 0, 0
	if len(nums) == 1 {
		fmt.Println("45. Jump Game II", 0)
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		coverage = max(coverage, i+nums[i])
		if i == lastJumpIdx {
			lastJumpIdx = coverage
			totalJump++
		}
	}
	fmt.Println("45. Jump Game II", totalJump)
}

func FindIn2DSortedMatrix() {
	matrix := [][]int{
		{10, 20, 30, 40},
		{12, 25, 33, 44},
		{15, 29, 37, 47},
		{27, 33, 42, 49},
	}
	target := 42
	row, col := 0, len(matrix[0])-1
	for col >= 0 && row < len(matrix) {
		if matrix[row][col] == target {
			fmt.Println("Find In 2D Sorted Matrix on index", row, col)
			return
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	fmt.Println("Target element not found")
}

func H_Index() {
	citations := []int{4, 4, 0, 0}
	sort.Slice(citations, func(i, j int) bool { return citations[i] > citations[j] })
	fmt.Println(citations)
	nonZeroes := 0
	for i := 0; i < len(citations); i++ {
		if citations[i] == 0 {
			continue
		}
		if citations[i] < nonZeroes+1 {
			fmt.Println("274. H-Index", nonZeroes)
			return
		}
		nonZeroes++
	}
	fmt.Println("274. H-Index", nonZeroes)
}

func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    fuelLeft, globalFuelLeft, start := 0, 0, 0
    for i := 0; i < n; i++ {
        globalFuelLeft += gas[i] - cost[i]
        fuelLeft += gas[i] - cost[i]
        if fuelLeft < 0 {
            start = i + 1
            fuelLeft = 0
        }
    }

    if globalFuelLeft < 0 {
        return -1
    }
    return start
}