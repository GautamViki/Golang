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
}
