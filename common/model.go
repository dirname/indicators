package common

import "time"

//Ticker ticker price information
type Ticker struct {
	Price float64
	Date  time.Time
}
