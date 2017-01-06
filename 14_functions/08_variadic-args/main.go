package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	var dataA float64 = 32
	var dataB float64 = 39

	n := average(dataA, dataB, data...)

	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, value := range sf {
		total += value
	}
	return total / float64(len(sf))
}
