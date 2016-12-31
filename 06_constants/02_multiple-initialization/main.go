package main

import "fmt"

const (
	//untyped
	pi       = 3.14
	language = "Go"
	// typed
	name string = "Ana"
)

func main() {

	fmt.Println(pi)
	fmt.Println(language)
	fmt.Println(name)
	fmt.Printf("%T \n", pi)
	fmt.Printf("%T \n", language)
	fmt.Printf("%T \n", name)
}
