package ma

import (
	"time"
)

//trimaCalculator TRIMA
type trimaCalculator struct {
	*Ticker
	Period       int32
	Count        int32
	Factor       float64
	Idx          float64
	MiddleIdx    int32
	TrailingIdx  int32
	Temp         []float64
	NumeratorSub float64
	NumeratorAdd float64
	Numerator    float64
	SubDone      bool
	AddDone      bool
	LastValue    float64
	Result       float64
}

//TRIMA TRIMA
type TRIMA struct {
	Value      float64
	Calculator *trimaCalculator
}

//NewTRIMA new TRIMA
func (t *Ticker) NewTRIMA(inTimePeriod int32) *TRIMA {
	calculator := &trimaCalculator{
		Ticker: t,
		Period: inTimePeriod,
		Idx:    float64(inTimePeriod >> 1),
	}
	return &TRIMA{
		Calculator: calculator,
	}
}

//calcTRIMA calculate TRIMA
func (t *trimaCalculator) calcTRIMA() {
	t.Temp = append(t.Temp, t.Price)
	if t.Period&1 == 1 && !t.SubDone && !t.AddDone {
		t.Factor = (t.Idx + 1.0) * (t.Idx + 1.0)
		t.Factor = 1.0 / t.Factor
		t.MiddleIdx = t.TrailingIdx + int32(t.Idx)
		t.Count = t.MiddleIdx + int32(t.Idx)
	} else if t.Period&1 != 1 {
		t.Factor = (t.Idx) * (t.Idx + 1)
		t.Factor = 1.0 / t.Factor
		t.MiddleIdx = t.TrailingIdx + int32(t.Idx) - 1
		t.Count = t.MiddleIdx + int32(t.Idx)
	}
	if int32(len(t.Temp)) == t.MiddleIdx+1 && !t.SubDone {
		for idx := t.MiddleIdx; idx >= t.TrailingIdx; idx-- {
			t.NumeratorSub += t.Temp[idx]
			t.Numerator += t.NumeratorSub
		}
		t.SubDone = true
	}
	if int32(len(t.Temp)) == t.Count+1 && !t.AddDone {
		t.MiddleIdx++
		for idx := t.MiddleIdx; idx <= t.Count; idx++ {
			t.NumeratorAdd += t.Temp[idx]
			t.Numerator += t.NumeratorAdd
		}
		t.AddDone = true
		t.LastValue = t.Temp[t.TrailingIdx]
		t.Result = t.Numerator * t.Factor
		t.Count++
		t.TrailingIdx++
	}
	if t.AddDone && t.SubDone && int32(len(t.Temp)) > t.Count {
		t.Temp = t.Temp[t.TrailingIdx:]
		t.MiddleIdx -= t.TrailingIdx
		t.Count -= t.TrailingIdx
		t.TrailingIdx = 0
		t.Numerator -= t.NumeratorSub
		t.NumeratorSub -= t.LastValue
		t.LastValue = t.Temp[t.MiddleIdx]
		t.MiddleIdx++
		t.NumeratorSub += t.LastValue
		if t.Period&1 == 1 {
			t.Numerator += t.NumeratorAdd
			t.NumeratorAdd -= t.LastValue
		} else {
			t.NumeratorAdd -= t.LastValue
			t.Numerator += t.NumeratorAdd
		}
		t.LastValue = t.Temp[t.Count]
		t.Count++
		t.NumeratorAdd += t.LastValue
		t.Numerator += t.LastValue
		t.LastValue = t.Temp[t.TrailingIdx]
		t.TrailingIdx++
		t.Result = t.Numerator * t.Factor
	}
}

//Update Update the TRIMA value of the current price
func (d *TRIMA) Update(price float64, date time.Time) {
	d.Calculator.setPrice(price, date)
	d.Calculator.calcTRIMA()
	d.Value = d.Calculator.Result
}

//Sum Returns the TRIMA value of the current TRIMA object
func (d *TRIMA) Sum() float64 {
	return d.Value
}
