package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice - numero de elementos
	// 5 is capacity - number of elements in the underlying array - capacidad hasta que cree uno nuevo

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	//greeting[2] = "buenos dias!"
	greeting = append(greeting, "Buenos días")
	greeting = append(greeting, "Suprabadham")


	fmt.Println(greeting[0])
	fmt.Println(greeting[1])
	fmt.Println(greeting[2]) // -> empty
	fmt.Println(greeting[3])
	fmt.Println(greeting[4])
}
