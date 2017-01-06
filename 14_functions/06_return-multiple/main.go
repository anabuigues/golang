package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe "))
	fmt.Println(greet2("Jane ", "Doe "))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

func greet2(fname, lname string) (firstName string, lastName string) {
	firstName = fmt.Sprint(fname, lname)
	lastName = fmt.Sprint(lname, fname)

	return
}
