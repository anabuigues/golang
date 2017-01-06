package main

import "fmt"

func main() {
	// var x []int -> slice
	var x [58]int  //-> array
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	x[42] = 777
	fmt.Println(x[42])
}

/*

array definition
 - a numbered sequence of elements of a single type
 - does not change in size

An array is a numbered sequence of elements of a single type.
The number of elements is called the length and is never negative.
The length is part of the array's type; it must evaluate to a non-negative constant representable by a value of type int.
The length of an array a can be discovered using the built-in function len.
The elements can be addressed by integer indices 0 through len(a)-1.
Array types are always one-dimensional but may be composed to form multi-dimensional types.
not dynamic does not change in size

 */