package main

import "fmt"

func main() {

	fmt.Printf("%T \n", true)

	if true {
		fmt.Println("This ran")
	}

	if false {
		fmt.Println("This did not run")
	}
}
