package main

import (
	"fmt"
)

//package scope
var i string = "cowboy"

func main(){

	fmt.Println("Variables")
	a := 10
	b := "Golang"
	b1 := `my hat`
	b2 := 'M'
	c := 4.17
	d := true

	fmt.Printf("value: %v \t type: %T \n", a, a)
	fmt.Printf("value: %v \t type: %T \n", b, b)
	fmt.Printf("value: %v \t type: %T \n", b1, b1)
	fmt.Printf("value: %v \t type: %T \n", b2, b2)
	fmt.Printf("value: %v \t type: %T \n", c, c)
	fmt.Printf("value: %v \t type: %T \n", d, d)

}

