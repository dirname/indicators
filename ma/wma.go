package ma

import "time"

//wmaCalculator simple MA
type wmaCalculator struct {
	*Ticker
	Divider       int32
	PeriodSum     float64
	PeriodSub     float64
	TrailingValue float64
	Temp          []float64
	Period        int32
	Result        float64
}

//WMA variance
type WMA struct {
	Value      float64
	Calculator *wmaCalculator
}

//NewWMA new WMA
func (t *Ticker) NewWMA(inTimePeriod int32) *WMA {
	calculator := &wmaCalculator{
		Ticker:  t,
		Period:  inTimePeriod,
		Temp:    make([]float64, 0, inTimePeriod),
		Divider: (inTimePeriod * (inTimePeriod + 1)) >> 1,
	}
	return &WMA{
		Calculator: calculator,
	}
}

//calcWMA calculate WMA
func (s *wmaCalculator) calcWMA() {
	s.Temp = append(s.Temp, s.Price)
	s.PeriodSub += s.Price
	l := int32(len(s.Temp))
	if l >= s.Period {
		s.PeriodSum += s.Price * float64(s.Period)
		s.PeriodSub -= s.TrailingValue
		s.TrailingValue = s.Temp[0]
		s.Result = s.PeriodSum / float64(s.Divider)
		s.PeriodSum -= s.PeriodSub
		s.Temp = s.Temp[1:]
	} else {
		s.PeriodSum += s.Price * float64(len(s.Temp))
	}
}

//Update Update the WMA value of the current price
func (w *WMA) Update(price float64, date time.Time) {
	w.Calculator.setPrice(price, date)
	w.Calculator.calcWMA()
	w.Value = w.Calculator.Result
}

//Sum Returns the WMA value of the current WMA object
func (w *WMA) Sum() float64 {
	return w.Value
}
