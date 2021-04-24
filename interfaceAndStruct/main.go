package main

import (
	"fmt"
	"goGn/interfaceAndStruct/org"
)

func main() {
	fmt.Println("Hello")

	var p org.Identifiable = org.Person{}
	fmt.Println(p.ID())
}
