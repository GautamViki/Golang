package structs

import "fmt"

type person struct {
	FirstName string
	lastName  string
	Age       int
}

func NewPerson(FirstName, lastName string, age int) *person {
	return &person{
		FirstName: FirstName,
		lastName:  lastName,
		Age:       age,
	}
}

func (p *person) Walk() {
	fmt.Println("Walkinggggggggggggggggggg", p)
}

func (p *person) SetLastName(lastName string) {
	p.lastName = lastName
}

func (p *person) GetLastName() string {
	return p.lastName
}
