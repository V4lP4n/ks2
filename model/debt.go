package model

type Debt struct {
	Lender       Dealer
	Borrower     Dealer
	Id           int
	Amount       float64
	Percent      int
	DailyPercent float64 //countable

}

func (d *Debt) Create(amount float64, percent int, Lender Dealer, debtor Dealer) {
	d.Amount = amount
	d.Percent = percent
	d.Borrower = debtor
	d.Lender = Lender
	d.DailyPercent = float64(percent) / 365

}
