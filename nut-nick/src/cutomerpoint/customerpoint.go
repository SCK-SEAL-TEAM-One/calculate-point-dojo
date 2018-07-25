package cutomerpoint

type Customers struct {
	Name   string `json:"name"`
	Level  string `json:"level"`
	Point  int    `json:"point"`
	amount int    `json:"amount"`
}

func (c *Customers) SetCustomerPoint(name, level string, point, amount int) {

	return
}
