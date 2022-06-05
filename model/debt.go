package model

type Debt struct {
	Id           int
	Amount       float64
	Percent      int
	DailyPercent float64 //countable
	Lender       *Person
	Debtor       *Person
}

func (d *Debt) Create(amount float64, percent int, Lender *Person, debtor *Person) {
	d.Amount = amount
	d.Percent = percent
	d.Debtor = debtor
	d.Lender = Lender
	d.DailyPercent = float64(percent) / 365
	d.Debtor.Debth = append(d.Debtor.Debth, d)
	d.Lender.OwnDebts = append(d.Lender.OwnDebts, d)
}
