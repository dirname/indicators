package rsi

import "time"

//Ticker ticker price information
type Ticker struct {
	Price float64
	Date  time.Time
}

//rsiCalculator RSI calculator
type rsiCalculator struct {
	*Ticker
	TempValue float64
	calcValue float64
	PervValue float64
	Period    int32
	Count     int32
	PervGain  float64
	PervLoss  float64
	Result    float64
}

//RSI RSI object
type RSI struct {
	Value      float64
	Calculator *rsiCalculator
}

//NewRSI factory function return a new object of RSI
func (t *Ticker) NewRSI(inTimePeriod int32) *RSI {
	calculator := rsiCalculator{Ticker: t, Period: inTimePeriod}
	return &RSI{
		Calculator: &calculator,
	}
}

//setPrice set up price
func (t *Ticker) setPrice(price float64, date time.Time) {
	t.Price = price
	t.Date = date
}

//calcRSI calculate RSI
func (r *rsiCalculator) calcRSI() {
	if r.Count == 0 {
		r.PervValue = r.Price
	}
	r.calcValue = r.Price - r.PervValue
	r.PervValue = r.Price
	if r.Count > r.Period {
		r.PervLoss *= float64(r.Period - 1)
		r.PervGain *= float64(r.Period - 1)
	}
	if r.calcValue < 0 {
		r.PervLoss -= r.calcValue
	} else {
		r.PervGain += r.calcValue
	}
	if r.Count >= r.Period {
		r.PervLoss /= float64(r.Period)
		r.PervGain /= float64(r.Period)
		r.TempValue = r.PervGain + r.PervLoss
		if !((-1e-14 < r.TempValue) && (r.TempValue < 1e-14)) {
			r.Result = 100.0 * (r.PervGain / r.TempValue)
		} else {
			r.Result = 0.0
		}
	}
	r.Count++
}

//Update Update the RSI value of the current price
func (r *RSI) Update(price float64, date time.Time) {
	r.Calculator.setPrice(price, date)
	r.Calculator.calcRSI()
	r.Value = r.Calculator.Result
}

//Sum Returns the RSI value of the current RSI object
func (r *RSI) Sum() float64 {
	return r.Value
}
