package main

import (
	"fmt"

	"main.go/questions"
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
	fmt.Println("ValidPalindrome", dsa.ValidPalindrome())
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
}
