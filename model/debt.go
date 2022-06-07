package model

type Debt struct {
	Lender       *Person
	Borrower     *Person
	Id           int
	Amount       float64
	Percent      int
	DailyPercent float64 //countable

}

func (d *Debt) Create(amount float64, percent int, Lender *Person, debtor *Person) {
	d.Amount = amount
	d.Percent = percent
	d.Borrower = debtor
	d.Lender = Lender
	d.DailyPercent = float64(percent) / 365
	d.Borrower.Debth = append(d.Borrower.Debth, d)
	d.Lender.OwnDebts = append(d.Lender.OwnDebts, d)
}
