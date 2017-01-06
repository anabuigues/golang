package main

import "fmt"

func main() {
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	fmt.Println(mySlice[0])
	mySlice[0]++ //online for numeric values
	fmt.Println(mySlice[0])
	mySlice = append(mySlice, 8)
	fmt.Println(mySlice)

	mySliceString := make([]float64, 1)
	fmt.Println(mySliceString[0])
	mySliceString[0] = 34.5
	fmt.Println(mySliceString[0])
	mySliceString[0]++
	fmt.Println(mySliceString)
}
