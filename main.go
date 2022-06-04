package main

import (
	"fmt"
	"ks/model"
)

func main() {
	est1 := new(model.RealEstate)
	est1.Cost = 100
	est1.Id = 123

	vasya := new(model.Person)
	petya := new(model.Person)
	vasya.Id = 1
	vasya.Currency = 7
	petya.Id = 2
	petya.Currency = 120

	est1.OwnerId = vasya.Id
	vasya.Ownership.RealEstates = append(vasya.Ownership.RealEstates, est1.Id)

	fmt.Println(vasya, petya, est1)

	err := est1.Trade(vasya, petya, 110)
	fmt.Println(err)
	fmt.Println(vasya, petya, est1)

}
