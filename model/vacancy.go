package model

type Vacancy struct {
	Id          int
	Title       string
	Salary      float64
	CompanyId   int
	EducationId int
	Status      bool
}
