package bollinger

import "time"

//Ticker ticker price information
type Ticker struct {
	Price float64
	Date  time.Time
}

//setPrice set up price
func (t *Ticker) setPrice(price float64, date time.Time) {
	t.Price = price
	t.Date = date
}
