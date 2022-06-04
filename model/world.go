package model

type World struct {
	Companies  []Company
	Persons    []Person
	RealEstate []RealEstate
}

func (World) Create() {

}

type DefaultWorldSettings struct {
}
