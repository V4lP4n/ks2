package model

type World struct {
	Companies []Company
	Persons   []Person
}

func (World) Create() {

}

type DefaultWorldSettings struct {
}
