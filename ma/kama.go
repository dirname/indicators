package ma

import (
	"math"
	"time"
)

//KamaCalculator KAMA
type KamaCalculator struct {
	*Ticker
	Period        int32
	Count         int32
	Result        float64
	Max           float64
	Diff          float64
	SumROC        float64
	Temp          []float64
	PrevKama      float64
	PeriodROC     float64
	TrailingValue float64
	TempValue     float64
}

//KAMA KAMA
type KAMA struct {
	Value      float64
	Calculator *KamaCalculator
}

//NewKAMA new KAMA
func (t *Ticker) NewKAMA(inTimePeriod int32) *KAMA {
	max := 2.0 / 31.0
	calculator := &KamaCalculator{
		Ticker: t,
		Period: inTimePeriod,
		Max:    max,
		Diff:   (2.0 / 3.0) - max,
		Temp:   make([]float64, 0, inTimePeriod),
	}
	return &KAMA{
		Calculator: calculator,
	}
}

//calcKAMA calculate KAMA
func (c *KamaCalculator) calcKAMA() {
	c.Temp = append(c.Temp, c.Price)
	if c.Count <= c.Period && len(c.Temp) >= 2 {
		c.SumROC += math.Abs(c.Temp[c.Count-1] - c.Temp[c.Count])
	}
	if c.Count == c.Period {
		c.PrevKama = c.Temp[c.Count-1]
		c.PeriodROC = c.Price - c.Temp[0]
		c.TrailingValue = c.Temp[0]
	} else if c.Count > c.Period {
		c.Temp = c.Temp[1:]
		c.PeriodROC = c.Price - c.Temp[0]
		c.SumROC -= math.Abs(c.TrailingValue - c.Temp[0])
		c.SumROC += math.Abs(c.Price - c.Temp[len(c.Temp)-2])
		c.TrailingValue = c.Temp[0]
	}
	if c.Count >= c.Period {
		if (c.SumROC < c.PeriodROC) || ((-1e-14 < c.SumROC) && (c.SumROC < 1e-14)) {
			c.TempValue = 1.0
		} else {
			c.TempValue = math.Abs(c.PeriodROC / c.SumROC)
		}
		c.TempValue = (c.TempValue * c.Diff) + c.Max
		c.TempValue *= c.TempValue
		c.PrevKama = ((c.Price - c.PrevKama) * c.TempValue) + c.PrevKama
		c.Result = c.PrevKama
	}
	c.Count++
}

//Update Update the KAMA value of the current price
func (k *KAMA) Update(price float64, date time.Time) {
	k.Calculator.setPrice(price, date)
	k.Calculator.calcKAMA()
	k.Value = k.Calculator.Result
}

//Sum Returns the KAMA value of the current KAMA object
func (k *KAMA) Sum() float64 {
	return k.Value
}
