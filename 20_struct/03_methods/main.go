package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// Estamos asociando al type el metodo fullName
// aqui no existe el this/self como en otros lenguajes
func (p person) fullName() string {
	return p.first + p.last
}

func (p person) fullNameAndOneWord(oneWord string) string {
	return p.first + p.last + " " + oneWord
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	fmt.Println(p2.age)

	fmt.Println(p1.fullNameAndOneWord("Hola"))
}
