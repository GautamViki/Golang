package handler

import (
	"fmt"
	"math/rand"
)

type Employee struct {
	Name    string
	Address string
}

func NewEMp() Employee {
	return Employee{"Vikas", "UP"}
}
func Emp() {
	fmt.Println("emp details")
	fmt.Print(NewEMp())
	for i := 0; i < 6; i++ {
		fmt.Println("count", i)
	}
	i := 0
	for i < 10 {
		fmt.Println("ksdfhs", i)
		i++
	}
}

func GameRockCeasor() {
	var array []int
	[3]array={1,2,3,4}

}
