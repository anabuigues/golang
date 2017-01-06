package main

import "fmt"

func main() {

	xs := []int{1, 3, 5, 7, 9, 11}

	for i, value := range xs {
		fmt.Println(i, " - ", value)
	}

	//add element
	xs = append(xs, 13)

	fmt.Println(xs)

}
