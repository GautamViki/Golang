package handler

import "fmt"

type pattern struct{}

func NewPatternQuestions() *pattern {
	return &pattern{}
}

func (p *pattern) Print_Z() {
	num := 3
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

func (p *pattern) IsNumPrime(num int) {
	isPrime := true
	for i := 2; i*i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println("Is this prime number : ", isPrime)
}

func (p *pattern) AllFibonacci(n, first, second int) {
	if n == 0 {
		return
	}
	fmt.Print(first, " ")
	last := first + second
	first = second
	second = last
	p.AllFibonacci(n-1, first, second)
}

func (p *pattern) CountDigitOfNumber(num int) {
	
	count := 0
	for num > 0 {
		num /= 10
		count++
	}
	fmt.Println("\nNumber Of Digit : ", count)
}
