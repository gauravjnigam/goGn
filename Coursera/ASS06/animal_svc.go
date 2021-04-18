package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak.
 The user can issue a request to find out one of three things about an animal:
  1) the food that it eats,
  2) its method of locomotion, and
  3) the sound it makes when it speaks.
*/

type Animal struct {
	food, locomotion, noice string
}

func (animal *Animal) Eat() string {
	return animal.food
}

func (animal *Animal) Move() string {
	return animal.locomotion
}

func (animal *Animal) Speak() string {
	return animal.noice
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"grass", "fly", "peep"}
	snake := Animal{"grass", "slither", "hsss"}

	animalMap := make(map[string]Animal)
	animalMap["cow"] = cow
	animalMap["bird"] = bird
	animalMap["snake"] = snake

	input_scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" > ")
		input_scanner.Scan()
		input_str := input_scanner.Text()
		req := strings.Split(input_str, " ")
		fmt.Println(findAnimal(animalMap, req[0], req[1]))

	}
}

func findAnimal(animalMap map[string]Animal, name string, action string) string {
	if val, ok := animalMap[name]; ok {
		switch {
		case strings.Compare("eat", action) == 0:
			return val.Eat()
		case strings.Compare("move", action) == 0:
			return val.Move()
		case strings.Compare("speak", action) == 0:
			return val.Speak()
		}
	}

	return "value not found"
}
