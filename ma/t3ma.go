package ma

import (
	"time"
)

//kamaCalculator T3MA
type t3maCalculator struct {
	*Ticker
	Period    int32
	Factor    float64
	Count     int32
	Result    float64
	K         float64
	OneMinusK float64
	E1        float64
	E2        float64
	E3        float64
	E4        float64
	E5        float64
	E6        float64
	C1        float64
	C2        float64
	C3        float64
	C4        float64
	E1Done    bool
	E2Done    bool
	E3Done    bool
	E4Done    bool
	E5Done    bool
	E6Done    bool
	Temp      []float64
	TempValue float64
}

//T3MA T3MA
type T3MA struct {
	Value      float64
	Calculator *t3maCalculator
}

//NewT3MA new T3MA
func (t *Ticker) NewT3MA(inTimePeriod int32, factor float64) *T3MA {
	k := 2.0 / (float64(inTimePeriod) + 1.0)
	calculator := &t3maCalculator{
		Ticker:    t,
		Period:    inTimePeriod,
		Factor:    factor,
		K:         k,
		OneMinusK: 1.0 - k,
		Temp:      make([]float64, 0, inTimePeriod-1),
	}
	return &T3MA{
		Calculator: calculator,
	}
}

//calcT3MA calculate T3MA
func (c *t3maCalculator) calcT3MA() {
	if c.Count == 0 {
		c.TempValue = c.Price
		c.Count++
		return
	}
	c.Temp = append(c.Temp, c.Price)
	if int32(len(c.Temp)) == c.Period-1 {
		switch {
		case !c.E1Done:
			for _, v := range c.Temp {
				c.TempValue += v
			}
			c.E1 = c.TempValue / float64(c.Period)
			c.TempValue = c.E1
			c.E1Done = true
			c.Temp = c.Temp[0:0]
		case !c.E2Done:
			for _, v := range c.Temp {
				c.E1 = (c.K * v) + (c.OneMinusK * c.E1)
				c.TempValue += c.E1
			}
			c.E2 = c.TempValue / float64(c.Period)
			c.TempValue = c.E2
			c.E2Done = true
			c.Temp = c.Temp[0:0]
		case !c.E3Done:
			for _, v := range c.Temp {
				c.E1 = (c.K * v) + (c.OneMinusK * c.E1)
				c.E2 = (c.K * c.E1) + (c.OneMinusK * c.E2)
				c.TempValue += c.E2
			}
			c.E3 = c.TempValue / float64(c.Period)
			c.TempValue = c.E3
			c.E3Done = true
			c.Temp = c.Temp[0:0]
		case !c.E4Done:
			for _, v := range c.Temp {
				c.E1 = (c.K * v) + (c.OneMinusK * c.E1)
				c.E2 = (c.K * c.E1) + (c.OneMinusK * c.E2)
				c.E3 = (c.K * c.E2) + (c.OneMinusK * c.E3)
				c.TempValue += c.E3
			}
			c.E4 = c.TempValue / float64(c.Period)
			c.TempValue = c.E4
			c.E4Done = true
			c.Temp = c.Temp[0:0]
		case !c.E5Done:
			for _, v := range c.Temp {
				c.E1 = (c.K * v) + (c.OneMinusK * c.E1)
				c.E2 = (c.K * c.E1) + (c.OneMinusK * c.E2)
				c.E3 = (c.K * c.E2) + (c.OneMinusK * c.E3)
				c.E4 = (c.K * c.E3) + (c.OneMinusK * c.E4)
				c.TempValue += c.E4
			}
			c.E5 = c.TempValue / float64(c.Period)
			c.TempValue = c.E5
			c.E5Done = true
			c.Temp = c.Temp[0:0]
		case !c.E6Done:
			for _, v := range c.Temp {
				c.E1 = (c.K * v) + (c.OneMinusK * c.E1)
				c.E2 = (c.K * c.E1) + (c.OneMinusK * c.E2)
				c.E3 = (c.K * c.E2) + (c.OneMinusK * c.E3)
				c.E4 = (c.K * c.E3) + (c.OneMinusK * c.E4)
				c.E5 = (c.K * c.E4) + (c.OneMinusK * c.E5)
				c.TempValue += c.E5
			}
			c.E6 = c.TempValue / float64(c.Period)
			c.E6Done = true
			c.TempValue = c.Factor * c.Factor
			c.Temp = c.Temp[0:0]
			c.C1 = -(c.TempValue * c.Factor)
			c.C2 = 3.0 * (c.TempValue - c.C1)
			c.C3 = -6.0*c.TempValue - 3.0*(c.Factor-c.C1)
			c.C4 = 1.0 + 3.0*c.Factor - c.C1 + 3.0*c.TempValue
			c.Result = c.C1*c.E6 + c.C2*c.E5 + c.C3*c.E4 + c.C4*c.E3
		}
	}
	if c.E1Done && c.E2Done && c.E3Done && c.E4Done && c.E5Done && c.E6Done && len(c.Temp) > 0 {
		c.E1 = (c.K * c.Temp[0]) + (c.OneMinusK * c.E1)
		c.E2 = (c.K * c.E1) + (c.OneMinusK * c.E2)
		c.E3 = (c.K * c.E2) + (c.OneMinusK * c.E3)
		c.E4 = (c.K * c.E3) + (c.OneMinusK * c.E4)
		c.E5 = (c.K * c.E4) + (c.OneMinusK * c.E5)
		c.E6 = (c.K * c.E5) + (c.OneMinusK * c.E6)
		c.Result = c.C1*c.E6 + c.C2*c.E5 + c.C3*c.E4 + c.C4*c.E3
		c.Temp = c.Temp[0:0]
	}
	c.Count++
}

//Update Update the T3MA value of the current price
func (k *T3MA) Update(price float64, date time.Time) {
	k.Calculator.setPrice(price, date)
	k.Calculator.calcT3MA()
	k.Value = k.Calculator.Result
}

//Sum Returns the T3MA value of the current T3MA object
func (k *T3MA) Sum() float64 {
	return k.Value
}
