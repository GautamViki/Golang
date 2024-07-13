package handler

func Sum(num int) func(int) int {
	sum := num
	return func(val int) int {
		sum += val
		return sum
	}
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
