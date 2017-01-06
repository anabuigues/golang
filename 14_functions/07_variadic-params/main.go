package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	// If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.
	for _, value := range sf {
		total += value
	}
	return total / float64(len(sf))
}