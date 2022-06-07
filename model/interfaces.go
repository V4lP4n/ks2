package model

type Entity interface{}
type Dealer interface {
	MakeDeal(*Estate)
}
