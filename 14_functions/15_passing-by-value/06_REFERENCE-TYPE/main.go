package main

import "fmt"

func main() {
	// make([]T, length, capacity)
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Todd]
}

func changeMe(z []string) {
	z[0] = "Todd"
	fmt.Println(z) // [Todd]
}
