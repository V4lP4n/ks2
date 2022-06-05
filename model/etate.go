package model

import "errors"

type Estate struct {
	Id      int
	Name    string
	OwnerId int
	Type    string
	Cost    float64
	LoanId  int
}

func (e *Estate) Trade(seller *Person, byuer *Person, cost float64) error {

	// check if byuer have enough money
	if byuer.Currency < cost {
		return errors.New("byuer does not have enough money")
	}
	// check if seller own estate and delete it from gis ownership
	ok := false
	for i, o := range seller.Ownership.Estates {

		if o == e {
			ok = true
			seller.Ownership.RealEstates = append(seller.Ownership.RealEstates[:i], seller.Ownership.RealEstates[i+1:]...)
		}
	}
	if !ok {
		return errors.New("seller does not own estate")
	}

	// finish the deal

	byuer.Ownership.Estates = append(byuer.Ownership.Estates, e)
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
