package model

import "errors"

type Loan struct {
	Id           int
	Lender       Dealer
	Borrower     Dealer
	Amount       float64 // loan amount
	Term         int     //time in days
	Percent      int     // percent per year
	FullAmount   float64 // Countable
	Estate       []*Estate
	DailyPayment float64 // Countable
}

func (l *Loan) Create(lender Dealer, amount float64, term int, percent int) {
	l.Lender = lender
	l.Amount = amount
	l.Term = term
	l.Percent = percent
	l.FullAmount = (((float64(term) / 365) * float64(l.Percent)) * (l.Amount / 100)) + l.Amount
	l.DailyPayment = l.FullAmount / float64(l.Term)

}
func (l *Loan) AttachMoney(borrower Dealer) {
	l.Borrower = borrower
	l.Lender.UpdateCurrency(-l.Amount)
	l.Borrower.UpdateCurrency(l.Amount)

}
func (l *Loan) Close() {

	l = &Loan{}

}
func (l *Loan) Pay() error {
	if l.Borrower == nil {
		return errors.New("borrower does not exists")
	}
	if l.Lender == nil {
		return errors.New("lender does not exists")
	}
	//check if borrower have money if ok confirm deal
	if l.Borrower.ConfirmCurrency(l.DailyPayment) {
		//check daily payment more then amount
		if l.FullAmount > l.DailyPayment {
			l.Borrower.UpdateCurrency(-l.DailyPayment)
			l.Lender.UpdateCurrency(l.DailyPayment)
			l.FullAmount -= l.DailyPayment
		} else { //or close the loan
			l.Borrower.UpdateCurrency(-l.DailyPayment)
			l.Lender.UpdateCurrency(l.DailyPayment)
			l.Close()
		}

	} else { //else check if loan have mortgaged property
		if len(l.Estate) != 0 {
			// if do confiscate it and close Loan

			//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
			for _, e := range l.Estate {
				e.Owner = l.Lender

			}
			l.Close()

		} else { // else create debt and close loan

			d := new(Debt)
			d.Create(l.FullAmount, l.Percent, l.Lender, l.Borrower)
			l.Close()
		}
	}

	return nil
}
