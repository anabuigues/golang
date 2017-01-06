package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "Todd" // se ha reservado 35 :)
	// student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
