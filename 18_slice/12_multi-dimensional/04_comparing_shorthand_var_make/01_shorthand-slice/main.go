package main

import (
	"fmt"
)

func main() {
	student := []string{}
	students := [][]string{}
	student[0] = "Todd" //no se ha reservado espacio, hay que usar el append
	// student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
