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

type Animal interface {
	Eat()
	Move()
	Speek()
}

type Cow struct {
	food, locomotion, noice string
}

func NewCow() *Cow {
	cow := new(Cow)

	cow.food = "grass"
	cow.locomotion = "walk"
	cow.noice = "moo"

	return cow
}

func (cow *Cow) Eat() string {
	return cow.food
}
func (cow *Cow) Move() string {
	return cow.locomotion
}
func (cow *Cow) Speak() string {
	return cow.noice
}

type Bird struct {
	food, locomotion, noice string
}

func (bird *Bird) NewBird() {
	bird.food = "worms"
	bird.locomotion = "fly"
	bird.noice = "peep"
}

func (b *Bird) Eat() string {
	return b.food
}
func (b *Bird) Move() string {
	return b.locomotion
}
func (b *Bird) Speak() string {
	return b.noice
}

type Snake struct {
	food, locomotion, noice string
}

func (snake *Snake) NewSnake() {
	snake.food = "mice"
	snake.locomotion = "slither"
	snake.noice = "hsss"
}

func (s *Snake) Eat() string {
	return s.food
}
func (s *Snake) Move() string {
	return s.locomotion
}
func (s *Snake) Speak() string {
	return s.noice
}

const (
	NEW_ANIMAL string = "newanimal"
	QUERY      string = "query"
)

func main() {

	cowMap := make(map[string]Cow)
	birdMap := make(map[string]Bird)
	snakeMap := make(map[string]Snake)

	input_scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Usage : ")
	fmt.Println(" > newanimal name type[cow, bird, snake]")
	fmt.Println(" > query name actionp[eat, move, speak]")
	for {
		fmt.Print(" > ")
		input_scanner.Scan()
		input_str := input_scanner.Text()
		req := strings.Split(input_str, " ")
		switch req[0] {
		case NEW_ANIMAL:
			createAnimal(cowMap, birdMap, snakeMap, req)
		case QUERY:
			fmt.Println(queryAnimal(cowMap, birdMap, snakeMap, req))
		default:
			fmt.Println("Invalid input/command")
		}

	}
}

func createAnimal(cowMap map[string]Cow, birdMap map[string]Bird, snakeMap map[string]Snake, req []string) bool {
	if len(req) != 3 {
		fmt.Println("invalid command")
		return false
	}
	animal_name := req[1]
	animal_type := req[2]

	//fmt.Printf("NEW# name : %s, type : %s\n", animal_name, animal_type)

	switch animal_type {
	case "cow":
		cow := NewCow()
		cowMap[animal_name] = *cow
	case "bird":

	case "snake":
	default:
		fmt.Println("Invalid animal type.. supported animals are cow, bird, snake ")
	}

	return true
}

func queryAnimal(cowMap map[string]Cow, birdMap map[string]Bird, snakeMap map[string]Snake, req []string) string {
	if len(req) != 3 {
		fmt.Println("invalid command")
		return "value not found!!!"
	}
	animal_name := req[1]
	action := req[2]

	//fmt.Printf("QUERY#  name : %s, action : %s\n", animal_name, action)

	if val, ok := cowMap[animal_name]; ok {
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
