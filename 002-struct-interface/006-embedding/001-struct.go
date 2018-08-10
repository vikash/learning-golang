package main

import (
	"fmt"
	t "github.com/zopnow/learning-golang/002-struct-interface/006-embedding/types"
)

func main() {

	// Normal Struct
	p := t.P{FirstName: "Mahatma"}
	fmt.Println(p)

	// Stringer interface
	p2 := t.Person{FirstName: "Mahatma"}
	fmt.Println(p2)

	p3 := t.Person{"John", "Doe", t.Address{}}
	fmt.Println(p3)

	// We can directly access fields of embedded struct
	p3.Apartment = "123"
	p3.Building = "Awesome Building"

	// We can also access as a property of the embedded struct
	p3.Address.Country = "India"
	p3.Address.ZipCode = 560102
	p3.Address.Landmark = "Next to Tajmahal"
	p3.Address.Area = "HSR Layout"
	fmt.Println(p3.Address)

	// We can also call embedded struct methods
	fmt.Println("Indian? ", p3.Indian())

	// If both define same method - embedded type method won't be promoted.
	p3.Test()
	p3.Address.Test()
}
