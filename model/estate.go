package model

import "errors"

type Estate struct {
	Id    int
	Title string
	Owner Dealer
	Type  string
	Cost  float64
	Loan  *Loan
}

func (e *Estate) Trade(seller *Person, byuer *Person, cost float64) error {

	// check if byuer have enough money
	if byuer.Currency < cost {
		return errors.New("byuer does not have enough money")
	}
	// check if seller own estate and delete it from gis ownership
	ok := false
	for i, o := range seller.Estates {

		if o == e {
			ok = true
			seller.Estates = append(seller.Estates[:i], seller.Estates[i+1:]...)
		}
	}
	if !ok {
		return errors.New("seller does not own estate")
	}

	// finish the deal

	byuer.Estates = append(byuer.Estates, e)
	byuer.Currency -= cost
	seller.Currency += cost
	e.Cost = cost
	//check if

	return nil
}

// func (e *Estate) LoanTrade(seller *Person, byuer *Person, loan *Loan) error {

// 	//check if seller own estate

// 	if loan.LenderId

// 	return nil
// }
