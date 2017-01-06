package main

import (
	"fmt"
)

func main() {
	var student []string
	var students [][]string
	student[0] = "Todd" // lo mismo!
	// student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
