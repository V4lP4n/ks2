package model

type Company struct {
	Id       int
	Title    string
	Sphere   string
	Currency float64
}

func (c *Company) CreateDeal(estate *Estate) {

}
func (c *Company) ReturnId() int {

	return c.Id
}
func (c *Company) ConfirmCurrency(cost float64) bool {
	return c.Currency >= cost
}
func (c *Company) UpdateCurrency(sum float64) {
	c.Currency += sum
}

func (c *Company) Type() string {
	return "company"
}
