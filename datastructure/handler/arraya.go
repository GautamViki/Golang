package handler

import "fmt"

func AarrayDeclaration() {
	var firstType [4]int
	firstType[0] = 1
	firstType[1] = 2
	firstType[2] = 3
	firstType[3] = 4

	fmt.Println("firstType", firstType)

	secondType := [4]string{"abc", "sd", "erw", "efw"}
	fmt.Println("secondType", secondType)

	// length of array
	fmt.Println("length",len(secondType))

	// 2d array
	var firstType2D [2][2]int
	firstType2D[0][0]=0
	firstType2D[0][1]=1
	firstType2D[1][0]=2
	firstType2D[1][1]=3

	fmt.Println("firstType2D",firstType2D)

	secondType2D:=[2][2]int{{1,2},{3,4}}
	fmt.Println("secondType2D",secondType2D)
}
