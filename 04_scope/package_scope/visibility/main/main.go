package main

import (
	"fmt"
	"github.com/golang/04_scope/packageScope/visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)

	vis.PrintVar()
}
