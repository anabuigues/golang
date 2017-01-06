package main

import (
	"fmt"
)

func main() {
	var student []string //not initialize!
	var students [][]string //not initialize!
	fmt.Println(student) // []
	fmt.Println(students) // []
	fmt.Println(student == nil) // true
}
