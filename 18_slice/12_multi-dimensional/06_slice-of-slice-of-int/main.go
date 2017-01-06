package main

import (
	"fmt"
)

func main() {

	transactions := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)

	threeDimensions := make([][][]int, 0, 3)

	for i := 0; i < 3; i++ {
		twoDimensions := make([][]int, 0, 4)
		for j := 0; j < 4; j++ {
			oneDimension := make([]int, 0, 5)
			for k := 0; k < 5; k++ {
				oneDimension = append(oneDimension, k)
			}
			twoDimensions = append(twoDimensions, oneDimension)
		}
		threeDimensions = append(threeDimensions, twoDimensions)
	}

	fmt.Println(threeDimensions)
}
