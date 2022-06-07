package model

import "errors"

type Loan struct {
	Id           int
	Lender       *Person
	Borrower     *Person
	Amount       float64 // loan amount
	Term         int     //time in days
	Percent      int     // percent per year
	FullAmount   float64 // Countable
	Estate       []*Estate
	DailyPayment float64 // Countable
}

func (l *Loan) Create(lender *Person, amount float64, term int, percent int) {
	l.Lender = lender
	l.Amount = amount
	l.Term = term
	l.Percent = percent
	l.FullAmount = (((float64(term) / 365) * float64(l.Percent)) * (l.Amount / 100)) + l.Amount
	l.DailyPayment = l.FullAmount / float64(l.Term)
	l.Lender.Lendings = append(l.Lender.Lendings, l)

}
func (l *Loan) AttachMoney(borrower *Person) {
	l.Borrower = borrower
	l.Borrower.Loans = append(l.Borrower.Loans, l)
	l.Lender.Currency -= l.Amount
	l.Borrower.Currency += l.Amount

}
func (l *Loan) Close() {
	//remove lending from lender
	for i, e := range l.Lender.Lendings {
		if e == l {
			l.Lender.Lendings = append(l.Lender.Lendings[:i], l.Lender.Lendings[i+1:]...)
		}
	}
	//remove Loan from borrower
	for i, e := range l.Borrower.Loans {
		if e == l {
			l.Borrower.Loans = append(l.Borrower.Loans[:i], l.Borrower.Loans[i+1:]...)
		}
	}
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
	if l.Borrower.Currency > l.DailyPayment {
		//check daily payment more then amount
		if l.FullAmount > l.DailyPayment {
			l.Borrower.Currency -= l.DailyPayment
			l.Lender.Currency += l.DailyPayment
			l.FullAmount -= l.DailyPayment
		} else { //or close the loan
			l.Borrower.Currency -= l.FullAmount
			l.Lender.Currency += l.FullAmount
			l.Close()
		}

	} else { //else check if loan have mortgaged property
		if len(l.Estate) != 0 {
			// if do confiscate it and close Loan
			for _, e := range l.Estate {
				e.Owner = l.Lender
				l.Lender.Estates = append(l.Lender.Estates, e)
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
