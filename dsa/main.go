package main

import (
	"fmt"

	"main.go/questions"
)

func main() {
	dsa :=questions.NewDataStructure()
	fmt.Println(dsa.LongestCommonPrefix())
}
