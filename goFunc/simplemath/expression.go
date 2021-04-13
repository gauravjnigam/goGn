package simplemath

import (
	"errors"
)

/*
	Learning
	 - Function syntax
	 - return multiple values
	 - Blank identifier
	 - public and private functions
	 - Named return value

*/

func Add(f1 float64, f2 float64) float64 {
	return f1 + f2
}

func Sub(f1 float64, f2 float64) float64 {
	return f1 - f2
}

func Multiply(f1 float64, f2 float64) float64 {
	return f1 * f2
}

func Divide(f1, f2 float64) (result float64, err error) {
	if f2 == 0 {
		err = errors.New("Can't devide by value")
	}
	result = f1 / f2
	return
}
