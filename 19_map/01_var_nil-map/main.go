package main

import "fmt"

func main() {

	var myGreeting map[string]string
	fmt.Println(myGreeting) // map[]
	fmt.Println(myGreeting == nil) // true

	var myGreeting1 map[string]string
	fmt.Println(myGreeting1) //map[]
	fmt.Println(myGreeting1 == nil) //true -> not initialized

	var myGreeting2 = make(map[string]string)
	fmt.Println(myGreeting2) //map[]
	fmt.Println(myGreeting2 == nil) //false -> initialized

	myGreeting3 := make(map[string]string)
	fmt.Println(myGreeting3) //map[]
	fmt.Println(myGreeting3 == nil) // false -> initialized

	myGreeting4 := map[string]string{}
	fmt.Println(myGreeting4) // map
	fmt.Println(myGreeting4 == nil) // false -> initialized
}

// add these lines:
/*
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
*/
// and you will get this:
// panic: assignment to entry in nil map

/*

- key / value storage
- a “dictionary”
- an unordered group of elements of one type, called the element type
	- indexed by a set of unique keys of another type, called the key type.
	- The value of an uninitialized map is nil.
- reference type

 */