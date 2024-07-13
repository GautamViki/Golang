package handler

import "fmt"

func SliceDeclaration() {
	var slice1 []string
	fmt.Println("empty slice",slice1, "slice length",len(slice1))
	slice2:= []int{1,2,3,4,5,6}
	// printing the slice
	fmt.Println("slice1",slice1)
	// adding new element in to slice
	slice2=append(slice2, 100)
	fmt.Println("new slice1",slice1)

	slice3:=make([]int,7,10)
	fmt.Println("slice3 ",slice3)
	fmt.Println("capacity ",cap(slice3))
	fmt.Println("length",len(slice3))
	slice3[0]=10
	slice3[5]=20
	fmt.Println("slice3 ",slice3)
	slice3=append(slice3, 30)
	fmt.Println("slice2 ",slice3)

	// slicing
	fmt.Println("slicing",slice3[0:2])

	// iterate on slice2
	for idx,val:=range slice2{
		fmt.Println("slice2 index ",idx,"value ",val)
	}
}
