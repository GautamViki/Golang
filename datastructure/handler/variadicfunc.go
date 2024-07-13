package handler

import "fmt"

// it's take variable number of parameter
func VariadicHandler(nums ...int) {
	var arr [10]int
	// for idx, val := range nums {
	// 	arr[idx] = val
	// }
	copy(arr[:],nums)
	fmt.Println("variadic func",arr)

	fmt.Println("nums",len(nums))
}