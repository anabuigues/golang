package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {
	x := 5
	// pass by reference! :D
	zero(&x)
	fmt.Println(x) // x is 0
}
