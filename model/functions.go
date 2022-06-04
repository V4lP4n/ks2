package model

import "errors"

func FindPerson(id int, world *World) (Person, error) {
	var pers Person
	ex := false
	for _, p := range world.Persons {

		if p.Id != id {

		} else {
			pers = p
			ex = true
		}
	}
	if !ex {
		return pers, errors.New("Person does not exists")
	}
	return pers, nil
}
