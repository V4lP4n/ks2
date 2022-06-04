package model

type Company struct {
	Name      string
	Sphere    string
	Vacancies []Vacancy
	Jobs      []Job
	Workers   []Worker
	Id        int
	Office    RealEstate
	Currency  int
	Ownership Ownership
}
