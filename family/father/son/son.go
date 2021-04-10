package son

import (
	"fmt"
)

func init() {
	fmt.Println("Son package is initialized")
}

type Son struct {
	Name       string
	FatherName string
}

func (s Son) Data(name string, father string) string {
	s.Name = name
	s.FatherName = father
	return "Name : " + s.Name + ", Father : " + s.FatherName
}
