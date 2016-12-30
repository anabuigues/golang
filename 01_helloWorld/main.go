package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(42)
	fmt.Printf("%d - %b\n", 42, 42)
	fmt.Printf("%d - %b - %x\n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x\n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#x\n", 42, 42, 42)
	fmt.Println(math.Pi)

	for i := 0; i < 200; i++  {
		fmt.Printf("%d \t %b \t %x %q \n", i, i, i, i)
	}

}
