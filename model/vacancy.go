package model

type Vacancy struct {
	Id        int
	Title     string
	Salary    float64
	Company   *Company
	Education *Education
	Status    bool
}
