package model

type Deal struct {
	Id     int
	Cost   float64
	Seller *Person
	Buyer  *Person
	Estate Dealer
}
