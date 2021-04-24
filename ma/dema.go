package ma

import (
	"time"
)

//demaCalculator MA
type demaCalculator struct {
	*Ticker
	*EMA
	SecondEMA   *EMA
	FirstCount  int32
	SecondCount int32
	Period      int32
	Result      float64
}

//DEMA DEMA
type DEMA struct {
	Value      float64
	Calculator *demaCalculator
}

//NewDEMA new DEMA
func (t *Ticker) NewDEMA(inTimePeriod int32) *DEMA {
	calculator := &demaCalculator{
		Ticker:    t,
		EMA:       t.NewEMA(inTimePeriod),
		Period:    inTimePeriod,
		SecondEMA: new(Ticker).NewEMA(inTimePeriod),
	}
	return &DEMA{
		Calculator: calculator,
	}
}

//calcDEMA calculate DEMA
func (s *demaCalculator) calcDEMA() {
	s.FirstCount++
	if s.FirstCount >= s.Period {
		s.SecondEMA.Update(s.Sum(), s.Date)
		s.SecondCount++
	}
	if s.FirstCount > s.Period*2-2 && s.SecondCount > s.Period-1 {
		s.Result = (2.0 * s.Sum()) - s.SecondEMA.Sum()
	}
}

//Update Update the DEMA value of the current price
func (d *DEMA) Update(price float64, date time.Time) {
	d.Calculator.Update(price, date)
	d.Calculator.calcDEMA()
	d.Value = d.Calculator.Result
}

//Sum Returns the DEMA value of the current DEMA object
func (d *DEMA) Sum() float64 {
	return d.Value
}
