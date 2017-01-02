package main

import "fmt"

// pass by value
func zero(z int) {
	fmt.Printf("value z %d \n", z)
	z = 0
	fmt.Printf("value z %d \n", z)
}

// pass by reference
func zeroReference(z * int){
	fmt.Printf("value reference z %d \n", z)
	*z = 0
	fmt.Printf("value reference %d \n", z)

}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5

	zeroReference(&x)
	fmt.Println(x) // x is 0
}
