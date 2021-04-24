package ma

import "time"

//temaCalculator MA
type temaCalculator struct {
	*Ticker
	*EMA
	SecondEMA   *EMA
	ThirdEMA    *EMA
	FirstCount  int32
	SecondCount int32
	ThirdCount  int32
	Period      int32
	Result      float64
}

//TEMA TEMA
type TEMA struct {
	Value      float64
	Calculator *temaCalculator
}

//NewTEMA new TEMA
func (t *Ticker) NewTEMA(inTimePeriod int32) *TEMA {
	calculator := &temaCalculator{
		Ticker:    t,
		EMA:       t.NewEMA(inTimePeriod),
		Period:    inTimePeriod,
		SecondEMA: new(Ticker).NewEMA(inTimePeriod),
		ThirdEMA:  new(Ticker).NewEMA(inTimePeriod),
	}
	return &TEMA{
		Calculator: calculator,
	}
}

//calcTEMA calculate TEMA
func (s *temaCalculator) calcTEMA() {
	s.FirstCount++
	if s.FirstCount >= s.Period {
		s.SecondEMA.Update(s.Sum(), s.Date)
		s.SecondCount++
	}
	if s.SecondCount >= s.Period {
		s.ThirdEMA.Update(s.SecondEMA.Sum(), s.Date)
		s.ThirdCount++
	}
	if s.FirstCount > (s.Period*3)-3 && s.SecondCount > (s.Period*2)-2 && s.ThirdCount > s.Period-1 {
		s.Result = s.ThirdEMA.Sum() + ((3.0 * s.Sum()) - (3.0 * s.SecondEMA.Sum()))
	}
}

//Update Update the TEMA value of the current price
func (d *TEMA) Update(price float64, date time.Time) {
	d.Calculator.Update(price, date)
	d.Calculator.calcTEMA()
	d.Value = d.Calculator.Result
}

//Sum Returns the TEMA value of the current TEMA object
func (d *TEMA) Sum() float64 {
	return d.Value
}
