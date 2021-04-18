package main

import (
	"fmt"
	"interfaceAndStruct/org"
)

func main() {
	fmt.Println("Hello")

	var p org.Identifiable = org.Person{}
	fmt.Println(p.ID())
}