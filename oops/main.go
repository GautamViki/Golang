package main

import (
	"fmt"
	"oops/polymorphism"
	"oops/structs"
)

func main() {
	p := structs.NewPerson("asdsa", "sasfa", 574)
	// p.Age = 10
	// p.FirstName = "sdfsfsdf"
	// p.SetLastName("Gautamaaaaaaaaaaa")
	fmt.Println("Person", p.GetLastName())
	p.Walk()
	// var dog polymorphism.Animal = polymorphism.Dog{}
	cat := polymorphism.Cat{}
	// dog.Speak()
	polymorphism.MakeSpeak(&cat)
}
