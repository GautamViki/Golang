package main

import (
	"dsa/questions"
	v1 "dsa/v1"
	"fmt"
	"sync"
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
	dsa.SubarraySumEquals_K_560()
	dsa.IsomorphicString_205()
	dsa.WordPattern_290()
	dsa.TrappingRainWater_42()
	dsa.RotateImage_48()
	dsa.ReverseVowelsofString()
	dsa.RemoveDuplicatesfromSortedArray()
	dsa.ReverseWordsInString()
	dsa.FindSubsequences()
	subSequenceStr := dsa.FindSubsequencesOfString_1("ab", "")
	fmt.Println("FindSubsequencesOfString_1 ", subSequenceStr)
	subSequenceArr := dsa.FindSubsequencesOfArray_1([]int{1, 2}, []int{})
	fmt.Println("FindSubsequencesOfArray_1 ", subSequenceArr)

	sub := dsa.NumSubseq_1([]int{2, 3, 3, 4, 6, 7}, 12)
	fmt.Println("NumSubseq_1", sub)
	dsa.MinimizeMaximumPairSumInArray()
	dsa.CountTheNumberOfGoodPartitions()
	dsa.MinimumLengthOfStringAfterDeletingSimilarEnds()
	dsa.StringCompression()
	dsa.CheckIfTheSentenceIsPangram()
	//
	//
	fmt.Println("\n\n\n\n")
	v1.TwoSum_1()
	v1.ValidParentheses_20()
	v1.BestTimeToBuyAndSellStock()
	v1.ValidPalindrome()
	v1.ValidAnagram()
	v1.BinarySearch()
	v1.SingleNumber()
	v1.MajorityElement()
	v1.MissingNumber()
	v1.ReverseString()
	v1.MaximumSubarray()
	v1.BestTimeToBuyAndStock_II()
	v1.HouseRobber()
	v1.NumberOf_1_Bits()
	v1.MergeIntervals()
	v1.SetMatrixZeroes()
	v1.SpiralMatrix()
	v1.Sum_3()
	v1.ContainerWithMostWater()
	v1.GroupAnagrams()
	v1.GroupAnagrams_1()
	v1.LongestPalindromicSubstring()
	v1.MinimumWindowSubstring()
	v1.ContainsDuplicate_II()
	v1.RemoveDuplicatesfromSortedArray_II()
	v1.RotateArray()
	v1.JumpGame()
	v1.JumpGame_II()
	v1.FindIn2DSortedMatrix()
	v1.H_Index()
	v1.CanCompleteCircuit()
	v1.BagofTokens()
	v1.BoatstoSavePeople()
	v1.BreakAPalindrome()
	v1.IntegertoRoman()
	v1.IsSubsequence()
	v1.MinimumSizeSubarraySum()
	v1.GameofLife()
	v1.RansomNote()
	v1.LongestConsecutiveSequence()
	v1.SummaryRanges()
	v1.InsertInterval()
	v1.InsertInterval_1()
	v1.MinimumNumberofArrowstoBurstBalloons()
	v1.Candy()
	v1.MaximumSumCircularSubarray()
	v1.MaximumSumCircularSubarray_1()
	v1.SimplifyPath()
	v1.EvaluateReversePolishNotation()
	v1.FindMinimumInRotatedSortedArray()
	v1.FactorialTrailingZeroes()
	v1.SingleNumber_II()
	v1.WordBreak()
	v1.CoinChange()
	mergedArr := v1.MergeSort([]int{60, 90, 70, 20, 30, 80, 40, 50, 10})
	fmt.Println("Merged Sort :", mergedArr)
	fmt.Println("Subsequence of string", v1.FindSubsequences("abc", ""))
	v1.LongestIncreasingSubsequence()
	v1.Triangle()
	v1.BitwiseANDOfNumbersRange()
	sortedArr := v1.QuickSort([]int{10, 9, 2, 5, 3, 7, 101, 18}, 0, 7)
	fmt.Println("Sorted Array Quick Sort :", sortedArr)
	v1.LongestHappyString()
	v1.LongestValidParentheses()
	v1.Permute()
	v1.MultiplyStrings()
	v1.MultiplyStrings_1()
	v1.MinimumTotalDistanceTraveled()
	v1.DeleteCharacterstoMakeFancyString()
	v1.MinimumNumberofRemovalstoMountainArray()
	v1.RotateString()
	v1.CircularSentence()
	v1.StringCompressionIII()
	v1.CompressedString()

	nums = []int{10, 20, 30, 40, 50}
	numCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)
	go func(ch chan int, arr []int, wg *sync.WaitGroup) {
		defer wg.Done()
		for _, num := range arr {
			ch <- num
		}
		close(ch)
	}(numCh, nums, &wg)

	sum := 0
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for num := range ch {
			sum += num
		}
	}(numCh, &wg)
	wg.Wait()
	fmt.Println("\n111111111111111111111111111111", sum)
}
