package ma

import "time"

type EMA struct {
	Value      float64
	Calculator *emaCalculator
}

//emaCalculator emaCalculator calculation object
type emaCalculator struct {
	*Ticker
	Temp   float64
	Period int32
	Count  int32
	MA     float64
	K      float64
}

//calcEMA calculate emaCalculator
func (m *emaCalculator) calcEMA() {
	if m.Count < m.Period {
		m.Temp += m.Price
		m.Count++
		if m.Count == m.Period {
			m.MA = m.Temp / float64(m.Period)
		}
	} else {
		m.MA = ((m.Price - m.MA) * m.K) + m.MA
		m.Count++
	}
}

//NewEMA factory function return a new object of emaCalculator
func (t *Ticker) NewEMA(inTimePeriod int32) *EMA {
	calculator := &emaCalculator{
		Ticker: t,
		Period: inTimePeriod,
		K:      2.0 / float64(inTimePeriod+1),
	}
	return &EMA{
		Calculator: calculator,
	}
}

//Update Update the SMA value of the current price
func (m *EMA) Update(price float64, date time.Time) {
	m.Calculator.setPrice(price, date)
	m.Calculator.calcEMA()
	m.Value = m.Calculator.MA
}

//Sum Returns the SMA value of the current SMA object
func (m *EMA) Sum() float64 {
	return m.Value
}
