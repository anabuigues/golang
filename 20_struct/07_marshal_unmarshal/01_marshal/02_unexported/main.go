package main

import (
	"encoding/json"
	"fmt"
)

// no las exporta porque estan en minuscula
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
