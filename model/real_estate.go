package model

import "errors"

type RealEstate struct {
	Id           int
	OwnerId      int
	Type         string
	Cost         int
	UtilityBills int
}

func (r *RealEstate) Trade(seller *Person, byuer *Person, cost int) error {

	// check if byuer have enough money
	if byuer.Currency < cost {
		return errors.New("byuer does not have enough money")
	}
	// check if seller own estate and delete it from gis ownership
	ok := false
	for i, o := range seller.Ownership.RealEstates {

		if o == r.Id {
			ok = true
			seller.Ownership.RealEstates = append(seller.Ownership.RealEstates[:i], seller.Ownership.RealEstates[i+1:]...)
		}
	}
	if !ok {
		return errors.New("seller does not own estate")
	}

	// finish the deal

	byuer.Ownership.RealEstates = append(byuer.Ownership.RealEstates, r.Id)
	byuer.Currency -= cost
	seller.Currency += cost
	r.Cost = cost
	//check if

	return nil
}
