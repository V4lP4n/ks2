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

func (e *Estate) Trade(seller Dealer, byuer Dealer, cost float64) error {

	// check if byuer have enough money
	if !byuer.ConfirmCurrency(cost) {
		return errors.New("byuer does not have enough money")
	}
	// check if seller own estate and delete it from gis ownership
	if seller.ReturnId() != e.Owner.ReturnId() {
		return errors.New("seller does not own estate")
	}
	// finish the deal
	e.Owner = byuer
	byuer.UpdateCurrency(-cost)
	seller.UpdateCurrency(cost)
	e.Cost = cost

	return nil
}
