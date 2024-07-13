package handler

import "fmt"

type Rectangle struct {
	Height int
	Width  int
}

func NewRectangle() *Rectangle {
	fmt.Println("in constructer")
	return &Rectangle{}
}

func (r Rectangle) AreaByValue() {
	r.Height = 10
	r.Width = 20
	fmt.Println("Area by Value of Rectangle", r.Height*r.Width)
}

func (r *Rectangle)AreaByPointer(){
	r.Height = 5
	r.Width = 20
	fmt.Println("Area by pointer of Rectangle", r.Height*r.Width)
}
