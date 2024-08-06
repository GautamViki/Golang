package handler

import "fmt"

type pattern struct{}

func NewPatternQuestions() *pattern {
	return &pattern{}
}

func (p *pattern) Print_Z() {
	num := 10
	for i := 1; i <= num; i++ {
		if i == 1 || i == num {
			for i := 0; i < num; i++ {
				fmt.Print("*")
			}
		} else {
			for j := 1; j <= num-i; j++ {
				fmt.Print(" ")
			}
			fmt.Print("*")
		}
		fmt.Println()
	}
}
