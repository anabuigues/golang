package main

import "fmt"

func main() {
	greet("Jane", "Doe")
}

func greet(firstName string, lastName string) {
	fmt.Println(firstName, lastName)
}
