package model

type Person struct {
	//Basics
	Id         int
	FirstName  string
	SecondName string
	Ownership  Ownership
	//Career
	Qualification Qualification
	JobId         int

	// psyhic
	Happyness int8
	Greed     int8
	Tocxity   int8
	Courage   int8

	//assets
	Currency float64
	Lendings []*Loan
	Loans    []*Loan
	OwnDebts []*Debt
	Debth    []*Debt
}
