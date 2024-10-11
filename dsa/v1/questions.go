package v1

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
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
		n := len(result) - 1
		if result[n][1] >= intervals[i][0] {
			if result[n][1] < intervals[i][1] {
				result[n][1] = intervals[i][1]
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

func CanCompleteCircuit() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	currentTotalFuel, currentFuel, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		currentTotalFuel += gas[i] - cost[i]
		currentFuel += gas[i] - cost[i]
		if currentFuel < 0 {
			start = i + 1
			currentFuel = 0
		}
	}
	if currentTotalFuel < 0 {
		fmt.Println("134. Gas Station index", -1)
		return
	}
	fmt.Println("134. Gas Station index", start)
}

func BagofTokens() {
	tokens := []int{100, 200, 300, 400}
	power := 200
	slices.Sort(tokens)
	score, i, j := 0, 0, len(tokens)-1
	maxScore := 0
	for i <= j {
		if power >= tokens[i] {
			power -= tokens[i]
			score++
			i++
			maxScore = max(maxScore, score)
		} else if score > 0 {
			score--
			power += tokens[j]
			j--
		} else {
			fmt.Println("948. Bag of Tokens", maxScore)
			return
		}
	}
	fmt.Println("948. Bag of Tokens", maxScore)
}

func BoatstoSavePeople() {
	people := []int{3, 5, 3, 4}
	limit := 5
	slices.Sort(people)
	minBoat, i, j := 0, 0, len(people)-1
	for i <= j {
		if people[i]+people[j] <= limit {
			minBoat += 1
			i++
			j--
		} else {
			minBoat += 1
			j--
		}
	}
	fmt.Println("881. Boats to Save People", minBoat)
}

func BreakAPalindrome() {
	palindrome := "abccba"
	if len(palindrome) == 1 {
		fmt.Println("1328. Break a Palindrome", "")
		return
	}
	for i := 0; i < len(palindrome)/2; i++ {
		if palindrome[i] != 'a' {
			palindrome = palindrome[:i] + "a" + palindrome[i+1:]
			fmt.Println("1328. Break a Palindrome", palindrome)
			return
		}
	}
	palindrome = palindrome[:len(palindrome)-1] + "b"
	fmt.Println("1328. Break a Palindrome", palindrome)
}

func IntegertoRoman() {
	num := 3749
	roman := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""
	for i := 0; i < len(roman); i++ {
		if num == 0 {
			break
		}
		q := num / roman[i]
		num = num % roman[i]

		result += strings.Repeat(symbol[i], q)
	}
	fmt.Println("12. Integer to Roman", result)
}

func IsSubsequence() {
	s := "b"
	t := "ahbgdc"
	i, j := 0, 0
	for j < len(t) && i < len(s) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	fmt.Println("392. Is Subsequence", i == len(s))
}

func MinimumSizeSubarraySum() {
	target, nums := 7, []int{2, 3, 1, 2, 4, 3}
	minLen := math.MaxInt32
	i, j := 0, 0
	sum := 0
	for j < len(nums) {
		sum += nums[j]
		for sum >= target {
			minLen = min(minLen, j-i+1)
			sum -= nums[i]
			i++
		}
		j++
	}
	if minLen == math.MaxInt32 {
		minLen = 0
	}
	fmt.Println("209. Minimum Size Subarray Sum", minLen)
}

func GameofLife() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	numRows := len(board)
	if numRows == 0 {
		return
	}
	numCols := len(board[0])
	toBeToggled := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			liveCount := getLiveNeighborCount(board, numRows, numCols, i, j)
			if board[i][j] == 1 {
				if liveCount < 2 || liveCount > 3 {
					// Cell dies due to underpopulation or overpopulation
					toBeToggled = append(toBeToggled, []int{i, j})
				}
			} else {
				if liveCount == 3 {
					// Cell comes to life due to reproduction
					toBeToggled = append(toBeToggled, []int{i, j})
				}
			}
		}
	}

	// Toggle cells based on the stored coordinates
	for _, cor := range toBeToggled {
		if board[cor[0]][cor[1]] == 1 {
			board[cor[0]][cor[1]] = 0
		} else {
			board[cor[0]][cor[1]] = 1
		}
	}
	fmt.Println("289. Game of Life", board)
}

func isValidCell(numRows, numCols, i, j int) bool {
	// Check if the cell coordinates are within the bounds of the board
	return i >= 0 && j >= 0 && i < numRows && j < numCols
}

func getLiveNeighborCount(board [][]int, numRows, numCols, i, j int) int {
	// Define the offsets for adjacent cells
	dx := []int{-1, 1, 0, 0, -1, -1, 1, 1}
	dy := []int{0, 0, -1, 1, -1, 1, -1, 1}

	liveCount := 0
	for k := 0; k < 8; k++ {
		ni, nj := i+dx[k], j+dy[k]
		if isValidCell(numRows, numCols, ni, nj) && board[ni][nj] == 1 {
			liveCount++
		}
	}
	return liveCount
}

func RansomNote() {
	ransomNote, magazine := "bg", "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"
	ransomNoteMap := map[rune]int{}
	magazineMap := map[rune]int{}
	for _, s := range ransomNote {
		ransomNoteMap[s]++
	}
	for _, s := range magazine {
		magazineMap[s]++
	}
	for key, val := range ransomNoteMap {
		if val > magazineMap[key] {
			fmt.Println("383. Ransom Note", false)
			return
		}
	}
	fmt.Println("383. Ransom Note", true)
}

func LongestConsecutiveSequence() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	numMap := map[int]struct{}{}
	for _, num := range nums {
		numMap[num] = struct{}{}
	}
	longest := 0
	for _, num := range nums {
		currentNum := num
		if _, ok := numMap[currentNum-1]; !ok {
			count := 1
			for {
				if _, ok := numMap[currentNum+count]; ok {
					count++
					continue
				}
				longest = max(longest, count)
				break
			}
		}
	}
	fmt.Println("128. Longest Consecutive Sequence", longest)
}

func SummaryRanges() {
	nums := []int{0, 1, 2, 4, 5, 7, 8, 9, 11, 14}
	result := []string{}
	if len(nums) == 0 {
		fmt.Println("228. Summary Ranges", result)
		return
	}
	init, next, val := nums[0], nums[0], ""
	for i := 1; i < len(nums); i++ {
		if nums[i] == next+1 {
			next = nums[i]
		} else {
			if init != next {
				val = fmt.Sprintf("%v->%v", init, next)
			} else {
				val = fmt.Sprintf("%v", init)
			}
			result = append(result, val)
			init = nums[i]
			next = nums[i]
		}
	}
	if init != next {
		val = fmt.Sprintf("%v->%v", init, next)
	} else {
		val = fmt.Sprintf("%v", init)
	}
	result = append(result, val)
	fmt.Println("228. Summary Ranges", result)
}

func InsertInterval() {
	intervals, newInterval := [][]int{{1, 3}, {6, 9}}, []int{2, 5}
	i := 0
	for i < len(intervals) {
		if intervals[i][1] < newInterval[0] {
			i++
		} else if intervals[i][0] > newInterval[1] {
			intervals = append(intervals[:i], append([][]int{newInterval}, intervals[i:]...)...)
			fmt.Println("57. Insert Interval", intervals)
			return
		} else {
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
			intervals = append(intervals[:i], intervals[i+1:]...)
		}
	}
	intervals = append(intervals, newInterval)
	fmt.Println("57. Insert Interval", intervals)
}

func InsertInterval_1() {
	intervals, newInterval := [][]int{{1, 3}, {6, 9}}, []int{2, 5}
	i, result := 0, [][]int{}
	for i < len(intervals) {
		if intervals[i][1] < newInterval[0] {
			result = append(result, intervals[i])
		} else if intervals[i][0] > newInterval[1] {
			break
		} else {
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
		}
		i++
	}
	result = append(result, newInterval)
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}
	fmt.Println("57. Insert Interval _1", result)
}

func MinimumNumberofArrowstoBurstBalloons() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	arrow := points[0][1]
	count := 1
	for i := 1; i < len(points); i++ {
		if arrow >= points[i][0] {
			if arrow > points[i][1] {
				arrow = points[i][1]
			}
		} else {
			count++
			arrow = points[i][1]
		}
	}
	fmt.Println("452. Minimum Number of Arrows to Burst Balloons", count)
}

func Candy() {
	ratings := []int{1, 0, 2}

	// Initialize a slice to store candies, starting with 1 candy for each child
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	// First pass: left to right
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Second pass: right to left
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	// Sum the total candies
	sum := 0
	for _, candy := range candies {
		sum += candy
	}
	fmt.Println("135. Candy", sum)
}

func KadaneAlgoMax(nums []int) int {
	sum := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum = max(nums[i], sum+nums[i])
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func MaximumSumCircularSubarray() {
	nums := []int{5, -3, 5}
	maxSum := nums[0]
	for i := 0; i < len(nums); i++ {
		st := nums[0]
		nums = append(nums[1:], st)
		maxSum = max(maxSum, KadaneAlgoMax(nums))
	}
	fmt.Println("918. Maximum Sum Circular Subarray", maxSum)
}
func KadaneAlgoMin(nums []int) int {
	sum, minSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum = min(sum+nums[i], nums[i])
		minSum = min(minSum, sum)
	}
	return minSum
}

func MaximumSumCircularSubarray_1() {
	nums := []int{5, -3, 5}
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	circulerSum := sum - KadaneAlgoMin(nums)
	if KadaneAlgoMax(nums) > 0 {
		fmt.Println("918. Maximum Sum Circular Subarray _1:", max(KadaneAlgoMax(nums), circulerSum))
		return
	}
	fmt.Println("918. Maximum Sum Circular Subarray", KadaneAlgoMax(nums))
}

func SimplifyPath() {
	path := "/a/./b/../../c/"
	pathToken := strings.Split(path, "/")
	// [ a . b .. .. c ]
	stack := []string{}
	for _, str := range pathToken {
		if str == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if str != "." && str != "" {
			stack = append(stack, str)
		}
	}
	result := "/" + strings.Join(stack, "/")
	fmt.Println("71. Simplify Path", result)
}

func EvaluateReversePolishNotation() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	stack := []int{}
	for _, tn := range tokens {
		if tn == "*" || tn == "/" || tn == "+" || tn == "-" {
			n := len(stack) - 1
			value := 0
			switch tn {
			case "+":
				value = stack[n-1] + stack[n]
			case "-":
				value = stack[n-1] - stack[n]
			case "*":
				value = stack[n-1] * stack[n]
			case "/":
				value = stack[n-1] / stack[n]
			}
			stack = append(stack[:n-1], value)
		} else {
			num, _ := strconv.Atoi(tn)
			stack = append(stack, num)
		}
	}
	fmt.Println("150. Evaluate Reverse Polish Notation", stack[0])
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[low] <= nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func FindMinimumInRotatedSortedArray() {
	nums := []int{11, 13, 15, 17}
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Println("153. Find Minimum in Rotated Sorted Array", nums[r])
}

func FactorialTrailingZeroes() {
	n := 5
	x := 5
	ans := 0
	for n/5 > 0 && x < 10001 {
		ans += n / x
		x *= 5
	}
	fmt.Println("172. Factorial Trailing Zeroes", ans)
}

func SingleNumber_II() {
	nums := []int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}
	res := 0
	for bit := 1; bit != 0; bit <<= 1 {
		c := 0
		for _, n := range nums {
			if n&bit != 0 {
				c++
			}
		}
		if c%3 == 1 {
			res |= bit
		}
	}
	fmt.Println("137. Single Number II", res)
}

func WordBreak() bool {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	length := len(s)

	table := make(map[string]bool)
	table[""] = true // base case

	// Base case, because all input strings same as words themselves can be broken
	for _, word := range wordDict {
		table[word] = true
	}

	// In the recursion, we reduced the suffix at each level
	// So, in bottom up, we start from least suffix and build up to the original string
	endIndex := length - 1
	for startIndex := length - 1; startIndex >= 0; startIndex-- {

		subString := s[startIndex : endIndex+1]
		subStrLength := endIndex - startIndex + 1

		for _, word := range wordDict {
			wordLength := len(word)

			// Can the word be used as prefix in the substring
			if subStrLength >= wordLength {
				prefix := subString[:wordLength]
				suffix := subString[wordLength:]

				// With table[prefix], we check if the prefix is a valid word
				// We could have also done prefix == word, but string comparision is not O(1) in Go
				if table[prefix] && table[suffix] {
					table[subString] = true

					// We have found a substring, that can successfuly be broken into words
					break
				}
			}
		}
	}

	return table[s]
}

func CoinChange() {
	coins := []int{1, 2, 5}
	ammount := 11

	dp := make([]int, ammount+1)
	for i := range dp {
		dp[i] = -1
	}
	count := coinChangeRecursion_DP(coins, ammount, dp)
	fmt.Println("322. Coin Change by dp", count)
	minCoin := CoinChange_Tabular(coins, ammount)
	if minCoin == math.MaxInt32 {
		minCoin = -1
	}
	fmt.Println("322. Coin Change by tabular ", minCoin)
}

func coinChangeRecursion_DP(coins []int, ammount int, dp []int) int {
	if ammount == 0 {
		return 0
	}
	if ammount < 0 {
		return math.MaxInt32
	}
	if dp[ammount] != -1 {
		return dp[ammount]
	}
	minCoin := math.MaxInt32
	for _, coin := range coins {
		ans := coinChangeRecursion_DP(coins, ammount-coin, dp)
		if ans != math.MaxInt32 {
			minCoin = min(minCoin, ans+1)
		}
	}
	dp[ammount] = minCoin
	return minCoin
}

func CoinChange_Tabular(coins []int, amount int) int {
	table := make([]int, amount+1)
	for i := range table {
		table[i] = math.MaxInt32
	}
	table[0] = 0
	for i := 0; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && table[i-coins[j]] != math.MaxInt32 {
				table[i] = min(table[i], 1+table[i-coins[j]])
			}
		}
	}
	return table[amount]
}

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	leftNum := MergeSort(nums[:mid])
	rightNum := MergeSort(nums[mid:])
	return merge(leftNum, rightNum)
}

func merge(leftNum, rightNum []int) []int {
	mergerArr := make([]int, 0, len(leftNum)+len(rightNum))
	i, j := 0, 0
	for i < len(leftNum) && j < len(rightNum) {
		if leftNum[i] < rightNum[j] {
			mergerArr = append(mergerArr, leftNum[i])
			i++
		} else {
			mergerArr = append(mergerArr, rightNum[j])
			j++
		}
	}
	mergerArr = append(mergerArr, leftNum[i:]...)
	mergerArr = append(mergerArr, rightNum[j:]...)
	return mergerArr
}

var resultStr []string

func FindSubsequences(str string, currentStr string) []string {
	if len(str) == 0 {
		resultStr = append(resultStr, currentStr)
		return resultStr
	}
	resultStr = FindSubsequences(str[1:], currentStr+string(str[0]))
	resultStr = FindSubsequences(str[1:], currentStr)
	return resultStr
}

var dp = make([][]int, 0)

func LongestIncreasingSubsequence() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	dp = make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums)+1)
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			dp[i][j] = -1
		}
	}
	maxLen := LongestIncreasingSubsequenceRecursion(nums, 0, -1)
	fmt.Println("300. Longest Increasing Subsequence", maxLen)
}

func LongestIncreasingSubsequenceRecursion(nums []int, idx, pre int) int {
	if idx == len(nums) {
		return 0
	}
	if pre != -1 && dp[idx][pre] != -1 {
		return dp[idx][pre]
	}
	take := 0
	if pre == -1 || nums[idx] > nums[pre] {
		take = 1 + LongestIncreasingSubsequenceRecursion(nums, idx+1, idx)
	}
	skip := LongestIncreasingSubsequenceRecursion(nums, idx+1, pre)
	if pre != -1 {
		dp[idx][pre] = max(take, skip)
	}
	return max(take, skip)
}







