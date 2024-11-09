package main

import "fmt"

func main() {
	myAddress := CreateAddress().
		SetStreet("Millitery Zone").
		SetNumber(8).
		SetCity("Ilorin").
		SetCountry("Nigeria")

	fmt.Println("My Address: ", myAddress)
}

type Address struct {
	street  string
	number  int32
	city    string
	country string
}

func CreateAddress() *Address {
	return &Address{}
}

func (a *Address) SetStreet(streetName string) *Address {
	a.street = streetName
	return a
}

func (a *Address) SetNumber(addressNumber int32) *Address {
	a.number = addressNumber
	return a
}

func (a *Address) SetCity(city string) *Address {
	a.city = city
	return a
}

func (a *Address) SetCountry(country string) *Address {
	a.country = country
	return a
}
