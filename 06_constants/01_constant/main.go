package main

import "fmt"

const p = "death & taxes"

func main() {

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Printf("%T \n", p)
	fmt.Printf("%T \n", q)

}

// a CONSTANT is a simple unchanging value
