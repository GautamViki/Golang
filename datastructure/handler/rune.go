package handler

import (
	"fmt"
	"unicode"
)

func RuneHandler() {
	var r rune = 'b'
	fmt.Println("rune ", r)

	// string to rune conversion
	str := "abcd"
	rune1 := []rune(str)
	fmt.Println("rune1", rune1)

	// iterating over rune
	for idx, val := range rune1 {
		fmt.Println("rune1 idx:", idx, "value:", val)
	}
	fmt.Println()
	greet := "Hello, how are you?"
	// check is letter and remove non letter
	var letter string
	for _, val := range greet {
		if unicode.IsLetter(val) {
			fmt.Println("letter", val)
			// converting to uppercase and storing in letter var
			letter += string(unicode.ToUpper(val))
		} else {
			fmt.Println("non letter", val)
		}
	}
	fmt.Println("only letter", letter)
}
