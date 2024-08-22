package main

import (
	"fmt"
	"pattern-problem/handler"
)

func main() {
	fmt.Println("pattern")
	pattern := handler.NewPatternQuestions()
	pattern.Print_Z()
	pattern.IsNumPrime(11)
	fmt.Print("All fibnocci number : ")
	pattern.AllFibonacci(11,0,1)
	pattern.CountDigitOfNumber(12345)
	fmt.Print("Print digit : ")
	pattern.PrintDigit(10000)
}
