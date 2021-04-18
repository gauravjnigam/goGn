package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice,
and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file,
your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

const (
	MAX_LENGTH = 20
)

type Name struct {
	fName string
	lName string
}

func (name *Name) Set(first string, last string) {
	name.fName = first
	name.lName = last

	if len(first) > MAX_LENGTH {
		name.fName = first[:MAX_LENGTH]
	}

	if len(last) > MAX_LENGTH {
		name.lName = last[:MAX_LENGTH]
	}
}

func (name *Name) printName() string {
	return fmt.Sprintf("FirstName : %s, LastName : %s", name.fName, name.lName)
}

func main() {
	var fileName string
	fmt.Print("Enter the file name : ")
	fmt.Scan(&fileName)

	var nameSlice []Name

	fh, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file ", err)
	}

	scanner := bufio.NewScanner(fh)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		splittedLine := strings.Split(line, " ")
		name := new(Name)
		name.Set(splittedLine[0], splittedLine[1])
		nameSlice = append(nameSlice, *name)

	}

	printNameSlice(nameSlice)

}

func printNameSlice(n []Name) {
	for _, name := range n {
		fmt.Printf("%v \n", name.printName())
	}

}
