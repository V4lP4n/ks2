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
	User      *User
	//assets
	Currency float64
	// Lendings []*Loan
	// Loans    []*Loan
	// OwnDebts []*Debt
	// Debth    []*Debt
}

func (p *Person) CreateDeal(estate *Estate) {

}
func (p *Person) ReturnId() int {

	return p.Id
}
func (p *Person) ConfirmCurrency(cost float64) bool {
	return p.Currency >= cost
}
func (p *Person) UpdateCurrency(sum float64) {
	p.Currency += sum
}

func (p *Person) Type() string {
	return "person"
}
