package main

import (
	"fmt"

	"main.go/questions"
)

func main() {
	dsa := questions.NewDataStructure()
	fmt.Println("LongestCommonPrefix==>", dsa.LongestCommonPrefix())
	fmt.Println("ValidParentheses==>", dsa.ValidParentheses())
}
