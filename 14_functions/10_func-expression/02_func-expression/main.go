package main

import "fmt"

func main() {

	// funcion anonima asignada a una variable, para ejecutarla tenemos que colocar los () al lado de la variable
	greeting := func() {
		fmt.Println("Hello world!")
	}

	greeting()
}
