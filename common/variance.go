package common

import (
	"time"
)

//varianceCalculator variance
type varianceCalculator struct {
	*Ticker
	Temp        []float64
	TempValue   float64
	Period      int32
	PeriodTotal float64
	Result      float64
}

//Variance variance
type Variance struct {
	Value      float64
	Calculator *varianceCalculator
}

//NewVariance new variance
func (t *Ticker) NewVariance(inTimePeriod int32) *Variance {
	calculator := &varianceCalculator{
		Ticker: t,
		Period: inTimePeriod,
		Temp:   make([]float64, 0, inTimePeriod),
	}
	return &Variance{
		Calculator: calculator,
	}
}

//calcMFI calculate variance
func (s *varianceCalculator) calcVariance() {
	s.Temp = append(s.Temp, s.Price)
	s.PeriodTotal += s.Price
	s.TempValue += s.Price * s.Price
	l := int32(len(s.Temp))
	if l >= s.Period {
		meanValue := s.PeriodTotal / float64(s.Period)
		calcValue := s.TempValue / float64(s.Period)
		s.PeriodTotal -= s.Temp[0]
		s.TempValue -= s.Temp[0] * s.Temp[0]
		s.Result = calcValue - meanValue*meanValue
		s.Temp = s.Temp[1:]
	}
}

//setPrice set up price
func (t *Ticker) setPrice(price float64, date time.Time) {
	t.Price = price
	t.Date = date
}

//Update Update the variance value of the current price
func (v *Variance) Update(price float64, date time.Time) {
	v.Calculator.setPrice(price, date)
	v.Calculator.calcVariance()
	v.Value = v.Calculator.Result
}

//Sum Returns the variance value of the current variance object
func (v *Variance) Sum() float64 {
	return v.Value
}
