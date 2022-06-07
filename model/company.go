package model

type Company struct {
	Id        int
	Title     string
	Sphere    string
	Vacancies []*Vacancy
	Jobs      []*Job
	Workers   []*Worker
	Office    []*Estate
	Currency  float64
}
