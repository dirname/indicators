package ma

import (
	"time"
)

//Ticker ticker price information
type Ticker struct {
	Price float64
	Date  time.Time
}

//smaCalculator simple MA
type smaCalculator struct {
	*Ticker
	Temp        []float64
	Period      int32
	PeriodTotal float64
	MA          float64
}

//SMA simple MA
type SMA struct {
	Value      float64
	Calculator *smaCalculator
}

//NewSMA factory function return a new object of SMA
func (t *Ticker) NewSMA(inTimePeriod int32) *SMA {
	calculator := &smaCalculator{
		Ticker: t,
		Period: inTimePeriod,
		Temp:   make([]float64, 0, inTimePeriod),
	}
	return &SMA{
		Calculator: calculator,
	}
}

//calcSMA calculate SMA
func (s *smaCalculator) calcSMA() {
	s.Temp = append(s.Temp, s.Price)
	s.PeriodTotal += s.Price
	l := int32(len(s.Temp))
	if l >= s.Period {
		s.MA = s.PeriodTotal / float64(s.Period)
		s.PeriodTotal -= s.Temp[0]
		s.Temp = s.Temp[1:]
	}
}

//Update Update the SMA value of the current price
func (m *SMA) Update(price float64, date time.Time) {
	m.Calculator.setPrice(price, date)
	m.Calculator.calcSMA()
	m.Value = m.Calculator.MA
}

//Sum Returns the SMA value of the current SMA object
func (m *SMA) Sum() float64 {
	return m.Value
}

//setPrice set up price
func (t *Ticker) setPrice(price float64, date time.Time) {
	t.Price = price
	t.Date = date
}
