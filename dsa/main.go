package main

import (
	"fmt"

	"main.go/questions"
	v1 "main.go/v1"
)

func main() {
	dsa := questions.NewDataStructure()
	fmt.Println("LongestCommonPrefix==>", dsa.LongestCommonPrefix())
	fmt.Println("ValidParentheses==>", dsa.ValidParentheses())
	fmt.Print("MergeTwoSortedLists", dsa.MergeTwoSortedLists())
	dsa.LengthOfLastWord()
	dsa.SquareRoot()
	fmt.Println("ClimbingStairs ", dsa.ClimbingStairs(37))
	fmt.Println("ThreeConsecutiveOdds ", dsa.ThreeConsecutiveOdds())
	fmt.Println("ValidPalindrome", dsa.ValidPalindrome_125())
	dsa.MajorityElement()
	dsa.AddBinary()
	dsa.RemoveDuplicate()
	dsa.ProductExceptSelf()
	dsa.MergeTwoSortedArray()
	dsa.SingleNumber()
	dsa.ContainsDuplicate()
	dsa.RemoveDuplicateFromSortedArray()
	dsa.ExcelSheetColumnTitle()
	fmt.Println("Is Happy Number? : ", dsa.HappyNumber(0))
	fmt.Println("Reverse Integer ", dsa.ReverseInteger(1534236469))
	dsa.ThreeSum()
	dsa.ArraySetInGO()
	dsa.LongestSubstringWithoutRepeatingCharacters()
	dsa.StringToInteger()
	dsa.MedianOfTwoSortedArrays()
	dsa.KeyBoardRow_500()
	dsa.LongestPalindromicSubstring_5()
	dsa.ZigzagConversion_6("PAYPALISHIRING", 3)
	dsa.ContainerWithMostWater_11()
	fmt.Println("ThreeSumClosest_16 : ", dsa.ThreeSumClosest_16())
	dsa.BestTimetoBuySellStock()
	fmt.Println("36. Valid Sudoku : ", dsa.IsValidSudoku_36())
	dsa.MaximumSubarray_53()
	dsa.FindDuplicate_287()
	fmt.Println("BinarySearch_704 ", dsa.BinarySearch_704())
	dsa.MissingNumber_268()
	dsa.ReverseString_344()
	fmt.Println("House robber_198: ", dsa.HouseRobber_198())
	dsa.MergeIntervals_56()
	dsa.FindResultantArrayRemovingAnagrams_2273()
	dsa.GroupAnagrams_49()
	dsa.KthLargestElementInAnArray_215()
	dsa.MaximalSquare_221()
	dsa.FindPeakElement_162()
	// Array
	nums := []int{3}
	max := dsa.FindMaxInArray(nums, 0, nums[0])
	fmt.Println("Find Max In Array", max)
	dsa.MaximumProductSubarray_152()
	dsa.Kadane_Algo()
	dsa.GroupAnagrams_49_v1()
	//
	// nums = []int{1, 1, 2, 2, 2, 2, 3}
	dsa.FindAllOccurrenc()
	dsa.FindFirstandLastPositionofElementinSortedArray_36()
	dsa.FindOccurrencesofanElementinanArray_3159()
	dsa.SortColor_75()
	dsa.SubarraySumEqualsTarget()

	//
	//
	fmt.Println("\n\n\n\n")
	v1.MaximumProductSubarray()
	v1.MaxSumOfSubArray()
	v1.LongestSubarrayConsistingOfUniqueElements()
	v1.LongestUniqueSubstring()
	v1.MergeIntervals()
	v1.DutchNationalFlag()
}
