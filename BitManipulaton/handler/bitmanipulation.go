package handler

import (
	"fmt"
	"slices"
	"strconv"
)

type bitmanipulaton struct{}

func NewBitManipulationHandler() *bitmanipulaton {
	return &bitmanipulaton{}
}

func (bit *bitmanipulaton) ConvertDecimalToBinary() {
	num := 854
	var result string
	for num > 0 {
		rem := num % 2
		result += strconv.Itoa(rem)
		num /= 2
	}
	resultRune := []rune(result)
	slices.Reverse(resultRune)
	fmt.Println("Binary is ", string(resultRune))
}

func (bit *bitmanipulaton) ConvertBinaryToDecimal() {
	str := "1101010110"
	var result int
	powerTwo := 1
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '1' {
			result += powerTwo
		}
		powerTwo *= 2
	}
	fmt.Println("Decimal is ", result)
}

func (bit *bitmanipulaton) SwapTwoNumber() {
	num1, num2 := 10, 20
	fmt.Println("Before swapped both number ", num1, num2)
	num1 = num1 ^ num2
	fmt.Println("num1", num1)
	num2 = num1 ^ num2
	fmt.Println("num2", num2)
	num1 = num1 ^ num2
	fmt.Println("num1", num1)
	fmt.Println("After swapped both number ", num1, num2)
}

func (bit *bitmanipulaton) CheckNthBitSet() {
	num, idx := 13, 2
	var isSet bool
	if num>>idx&1 == 1 {
		isSet = true
	}
	fmt.Println("Is Bit set? :", isSet)
}

func (bit *bitmanipulaton) ConvertNthBitToSet() {
	num, idx := 17, 2
	if num|(1<<idx) > num {
		fmt.Println("Setted successfully")
	}else{
		fmt.Println("Already have setted")
	}
}
