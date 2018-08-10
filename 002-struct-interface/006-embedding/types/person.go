package types

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Address
}

func (p Person) String() string {

	if p.FirstName == "" {
		p.FirstName = "<Unknown>"
	}

	if p.LastName == "" {
		p.LastName = "<Unknown>"
	}

	return fmt.Sprintf("[%s %s]", p.FirstName, p.LastName)
}

func (p Person) Test() {
	fmt.Println("Test from Person says Hi!")
}
