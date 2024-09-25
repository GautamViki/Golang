package polymorphism

import (
	"fmt"
	"oops/interfaces"
)

func MakeSpeak(a interfaces.Animal) {
	a.Speakaaa()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Bhow bhow bhowwwww")
}

type Cat struct{}

func (c *Cat) Speak() {
	fmt.Println("Mew Mew mew")
}
