package types

import (
	"fmt"
)

type Address struct {
	Apartment string
	Building  string
	Area      string
	Landmark  string
	ZipCode   int
	City      string
	Country   string
}

func (a Address) String() string {
	return fmt.Sprintf("Apt#%s, %s\nLandmark: %s,\n%s, %s,\n%s - %v ", a.Apartment, a.Building, a.Landmark, a.Area, a.City, a.Country, a.ZipCode)
}

func (a Address) Indian() bool {
	if a.Country == "India" {
		return true
	}
	return false
}

func (a Address) Test() {
	fmt.Println("Test from Address says Hi!")
}
