package model

type Job struct {
	Id      int
	Salary  float64
	Vacancy *Vacancy
	Person  *Person
	Company *Company
}
