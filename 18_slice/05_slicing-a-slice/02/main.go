package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	// from 1 to 2
	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2])
	// from 0 to 2
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2])
	// from 5 to last
	fmt.Print("[5:] ")
	fmt.Println(greeting[5:])
	//all elements
	fmt.Print("[:] ")
	fmt.Println(greeting[:])
}
