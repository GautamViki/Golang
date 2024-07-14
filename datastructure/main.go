package main

import (
	"datastructure/handler"
)

func main() {
	// handler.AarrayDeclaration()
	// fmt.Println("==============================================")
	// handler.SliceDeclaration()
	// fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	// handler.MapHandler()
	// fmt.Println("##############################################")
	// handler.RuneHandler()
	// fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	// handler.VariadicHandler(11)
	// // anonymous function
	// func(a, b int) {
	// 	fmt.Println("a", a, "b", b)
	// }(1, 2)

	// closures
	// sum := handler.Sum(10)
	// fmt.Println(sum(5))
	// fmt.Println(sum(10))

	// constructer
	// rect := handler.NewRectangle()
	// rect.AreaByValue()
	// rect.AreaByPointer()
	// rect.AreaByValue()
	// fmt.Println("Height",rect.Height)
	// fmt.Println("Width ",rect.Width)

	// linked list
	list:=handler.NewLinkedList()
	list.InsertNodeAtFirst(8)

}

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
//             if nums[i]+nums[j]==target{
//                 var arr []int
//                 arr[0]=i
//                 arr[1]=j
//                 return arr
//             }
// 		}
// 	}
//     return []int{}
// }