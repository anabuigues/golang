package main

import "fmt"

func main() {

	myGreeting := make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)

	myGreeting2 := map[string]string{}
	myGreeting2["Tim"] = "Good morning."
	myGreeting2["Jenny"] = "Bonjour."

	fmt.Println(myGreeting2)
}
