package main

import (
	mathfun "goGn/goFunc/simplemath"
	"log"

	"fmt"
)

func main() {
	fmt.Println("Learning func in GO.. lets GO...")

	fmt.Println("##calling anonymous func")
	anonymousFuncEx()

	fmt.Println("## calling funcFromFunc")
	funcFromFunc()

	fmt.Println(mathfun.Add(10.0, 30.0))

	val, err := mathfun.Divide(10, 0)
	if err != nil {
		log.Fatalf("Divide error !!!")
	}
	fmt.Println(val)

}

// anonymous function
func anonymousFuncEx() {
	a := func(str string) string {
		return str + " : called!!"
	}

	a("gnFun1")
}

// return function from function

func funcFromFunc() {
	addExp := mathExpression()
	println(addExp(10.0, 20.0))
}

func mathExpression() func(float64, float64) float64 {
	return mathfun.Add
}
