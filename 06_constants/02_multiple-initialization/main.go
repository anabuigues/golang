package main

import "fmt"

const (
	pi       = 3.14
	language = "Go"
)

func main() {

	fmt.Println(pi)
	fmt.Println(language)
	fmt.Printf("%T \n", pi)
	fmt.Printf("%T \n", language)
}
