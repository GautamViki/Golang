package v1

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func MaximumProductSubarray() {
	nums := []int{2, 2, 0, 3}
	// O(n^2)
	// max := nums[0]
	// pro := 1
	// for i := 0; i < len(nums); i++ {
	// 	for j := i; j < len(nums); j++ {
	// 		pro *= nums[j]
	// 		if pro > max {
	// 			max = pro
	// 		}
	// 	}
	// 	pro = 1
	// }
	// fmt.Println("dbhssssssssssssssssssssssss", max)

	// O(n)
	leftPro, rightPro := 1, 1
	left, right := 0, len(nums)-1
	maxPro := nums[0]
	for left < len(nums) || right >= 0 {
		leftPro *= nums[left]
		rightPro *= nums[right]
		maxPro = max(maxPro, max(leftPro, rightPro))
		if nums[left] == 0 {
			leftPro = 1
		}
		if nums[right] == 0 {
			rightPro = 1
		}
		left++
		right--
	}
	fmt.Println("Maximum Product Subarray", maxPro)
}

func MaxSumOfSubArray() {
	nums := []int{-2, -1}
	maxSum := nums[0]
	// Brutforce method O(n^2)
	// for i := 0; i < len(nums); i++ {
	// 	sum := 0
	// 	for j := i; j < len(nums); j++ {
	// 		sum += nums[j]
	// 		maxSum = max(maxSum, sum)
	// 	}
	// }
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		maxSum = max(maxSum, sum)
		if sum < 0 {
			sum = 0
		}
	}
	fmt.Println("Maximum Subarray", maxSum)
}

func LongestSubarrayConsistingOfUniqueElements() {
	// nums := []int{1, 2, 4, 4, 5, 6, 7, 8, 3, 4, 5, 3, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4}
	nums := []int{1, 2, 2, 1}
	result := []int{}
	tempMap := make(map[int]int)
	start := 0
	for end, num := range nums {
		if idx, found := tempMap[num]; found && idx >= start {
			start = idx + 1
		}
		tempMap[num] = end
		currentLength := end - start + 1
		if currentLength > len(result) {
			result = nums[start : end+1]
		}
	}
	fmt.Println("Longest Subarray consisting of unique elements from an Array", result)
}

func LongestUniqueSubstring() {
	s := "abba"
	maxSub := make(map[string]int)
	start := 0
	var result string
	for end, char := range s {
		if idx, found := maxSub[string(char)]; found && idx >= start {
			start = idx + 1
		}
		maxSub[string(char)] = end
		currentLength := end - start + 1
		if currentLength > len(result) {
			result = s[start : end+1]
		}
	}
	// var maxSubStr string
	// var length int
	// for i := 0; i < len(s); i++ {
	// 	var subStr string
	// 	for j := i; j < len(s); j++ {
	// 		if !strings.Contains(subStr, s[j:j+1]) {
	// 			subStr = s[i : j+1]
	// 		} else {
	// 			break
	// 		}
	// 		if len(subStr) > len(maxSubStr) {
	// 			maxSubStr = subStr
	// 			length = len(subStr)
	// 		}
	// 	}
	// }
	// fmt.Println(length,"ffffffffffffffffffffffff", maxSubStr)
	fmt.Println(result, "3. Longest Substring Without Repeating Characters\n", len(result))
}

func MergeIntervals() {
	intervals := [][]int{{1, 5}, {4, 5}}
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		lenth := len(result) - 1
		if result[lenth][1] >= intervals[i][0] {
			if result[lenth][1] < intervals[i][1] {
				result[lenth][1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}
	fmt.Println("Merge Intervals", result)
}

func DutchNationalFlag() {
	nums := []int{0, 1, 2, 0, 1, 2}
	i, mid, j := 0, 0, len(nums)-1
	for mid <= j {
		if nums[mid] == 0 {
			nums[i], nums[mid] = nums[mid], nums[i]
			i++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else if nums[mid] == 2 {
			nums[mid], nums[j] = nums[j], nums[mid]
			j--
		}
	}
	fmt.Println("Dutch National Flag", nums)
}

func SubarrayWithGivenSum() {
	nums, target := []int{1, 2, 3, 4}, 5
	start, i, sum := 0, 0, 0
	result := []int{-1, -1}
	diff := int(math.Abs(float64(target - math.MinInt32)))
	for i < len(nums) {
		sum = sum + nums[i]
		absDifference := int(math.Abs(float64(target - sum)))
		if sum > target {
			if diff > absDifference {
				diff = absDifference
			} else {
				sum -= nums[start]
				start++
			}
		} else if sum == target {
			result = []int{start, i}
			break
		}
		if diff > absDifference {
			diff = absDifference
		}
		i++
	}
	fmt.Println("Subarray with Given Sum", result)
}
func Three_Sum() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	slices.Sort(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, h := i+1, len(nums)-1
		for l < h {
			sum := nums[i] + nums[l] + nums[h]
			if sum < 0 {
				l++
			} else if sum > 0 {
				h--
			} else {
				for l < h && nums[l] == nums[l+1] {
					l++
				}
				for l < h && nums[h] == nums[h-1] {
					h--
				}
				result = append(result, []int{nums[i], nums[l], nums[h]})
				l++
				h--
			}
		}
	}
	fmt.Println("15. 3Sum", result)
}

func ThreeSumClosest() {
	nums := []int{-4, 2, 2, 3, 3, 3}
	slices.Sort(nums)
	target := 0
	diff := math.MaxInt32
	ans := 0
	for i := 0; i < len(nums)-2; i++ {
		l, h := i+1, len(nums)-1
		sum := 0
		for l < h {
			sum = nums[i] + nums[l] + nums[h]
			temp := math.Abs(float64(sum - target))
			if diff > int(temp) {
				ans = sum
				diff = int(temp)
			}
			if sum > target {
				h--
			} else if sum < target {
				l++
			} else {
				break
			}
		}
		if sum == target {
			ans = sum
			break
		}
	}
	fmt.Println("16. 3Sum Closest ", ans)
}
