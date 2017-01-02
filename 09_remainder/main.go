package main

import "fmt"

const ( // iota is reset to 0
	c0 = 0  // c0 == 0
	c1 = 1  // c1 == 1
	c2 = 2  // c2 == 2
)

func main() {
	x := 13 % 3
	fmt.Println(x)
	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	for i := 1; i < 70; i++ {
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}
