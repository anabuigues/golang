package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(42)
	fmt.Printf("%d - %b\n", 42, 42)
	fmt.Printf("%d - %b - %x\n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x\n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#x\n", 42, 42, 42)

}
