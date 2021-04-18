package org

type Identifiable interface {
	ID() string
}

type Person struct {

}

func (p Person) ID() string {
	return "Persion#01"
}