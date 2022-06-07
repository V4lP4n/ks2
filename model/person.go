package model

type Person struct {
	//Basics
	Id         int
	Name       string
	Experience int
	// psyhic
	Happiness int8
	Greed     int8
	Tocxity   int8
	Courage   int8

	//assets
	Currency float64
	Lendings []*Loan
	Loans    []*Loan
	OwnDebts []*Debt
	Debth    []*Debt
	Estates  []*Estate
}

func (p *Person) MakeDeal(estate *Estate) {

}
