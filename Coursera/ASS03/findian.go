package main

import (
	"fmt"
	"strings"
)

/*
Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’,
and contains the character ‘a’. The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!”
for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.

The program should print “Not Found!”
for the following strings, “ihhhhhn”, “ina”, “xian”.

*/

func main() {
	var str string

	fmt.Print("Enter the string : ")
	fmt.Scan(&str)

	lowerChStr := strings.ToLower(str)

	if strings.HasPrefix(lowerChStr, "i") && strings.HasSuffix(lowerChStr, "n") && strings.Contains(lowerChStr, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
