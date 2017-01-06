package main

import "fmt"

func main() {
	greet("Jane", "Doe")
}

func greet(firstName, lastName string) {
	fmt.Println(firstName, lastName)
}
