package model

type Entity interface{}
type Dealer interface {
	CreateDeal(*Estate)
	ReturnId() int
	ConfirmCurrency(float64) bool
	UpdateCurrency(float64)

	Type() string
}

type Initer interface {
	Init()
}
