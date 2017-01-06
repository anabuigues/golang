package main

import (
	"fmt"
)


func main(){

	fmt.Println("Default values variables")
	var e int
	var f string
	var g float64
	var h bool
	var pepe * string

	fmt.Printf("value: %v \t type: %T \n", e, e)
	fmt.Printf("value: %v \t type: %T \n", f, f)
	fmt.Printf("value: %v \t type: %T \n", g, g)
	fmt.Printf("value: %v \t type: %T \n", h, h)
	fmt.Printf("value: %v \t type: %T \n", pepe, pepe)

}

