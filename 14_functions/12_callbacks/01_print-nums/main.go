package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {

	odd := func(n int) {
		fmt.Println(n % 3 == 0)
	}

	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})

	visit([]int{1, 2, 3, 4}, odd)
}

// callback: passing a func as an argument
