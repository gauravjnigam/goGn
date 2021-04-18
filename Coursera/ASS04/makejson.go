package main

import (
	"encoding/json"
	"fmt"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

func main() {
	var name string
	var address string

	userMap := make(map[string]string)

	fmt.Print("Enter the name : ")
	fmt.Scan(&name)
	fmt.Print("Enter the address : ")
	fmt.Scan(&address)

	userMap["name"] = name
	userMap["address"] = address

	jsonData, _ := json.Marshal(userMap)
	fmt.Printf("Json data : %s \n", string(jsonData))

}
