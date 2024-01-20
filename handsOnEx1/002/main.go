package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

type SecretAgent struct {
	licenceToKill bool
	Person
}

type Human interface {
	Speak() string
}

func (p Person) Speak() string {
	return p.FirstName + " " + p.LastName
}

func (sa SecretAgent) Speak() string {

	if sa.licenceToKill != true {
		return sa.FirstName + " " + sa.LastName + " Does not have a licence to kill"
	} else {
		return sa.FirstName + " " + sa.LastName + " has  a licence to kill"
	}
}

func Vomit(h Human) {
	fmt.Printf("%T\n", h)
	fmt.Println(h)
	switch v := h.(type) {
	case Person:
		fmt.Println(v.FirstName)
	case SecretAgent:
		fmt.Println(v.FirstName)
	default:
		fmt.Println("unknown")
	}

}
