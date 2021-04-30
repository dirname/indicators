package bollinger

import (
	"github.com/dirname/indicators/common"
	"github.com/dirname/indicators/ma"
	"time"
)

//bollingerBandsCalculator bollingerBandsCalculator
type bollingerBandsCalculator struct {
	*Ticker
	Upper       float64
	Lower       float64
	Period      int32
	UpperBands  float64
	LowerBands  float64
	MiddleBands float64
	MAObject    ma.Object
	StdDev      *common.StdDev
}

//Bands Bands
type Bands struct {
	Calculator *bollingerBandsCalculator
}

//NewBands new NewBands
func (t *Ticker) NewBands(inTimePeriod int32, Upper float64, Lower float64, obj ma.Object) *Bands {
	calculator := &bollingerBandsCalculator{
		Ticker:   t,
		Upper:    Upper,
		Lower:    Lower,
		Period:   inTimePeriod,
		MAObject: obj,
		StdDev:   new(common.Ticker).NewStdDev(inTimePeriod, 1.0),
	}
	return &Bands{
		Calculator: calculator,
	}
}

//calcBands calculate Bands
func (c *bollingerBandsCalculator) calcBands() {
	c.MiddleBands = c.MAObject.Sum()
	switch {
	case c.Upper == c.Lower:
		if c.Upper == 1.0 {
			c.UpperBands = c.MiddleBands + c.StdDev.Sum()
			c.LowerBands = c.MiddleBands - c.StdDev.Sum()
		} else {
			c.UpperBands = c.MiddleBands + c.StdDev.Sum()*c.Upper
			c.LowerBands = c.MiddleBands - c.StdDev.Sum()*c.Upper
		}
	case c.Upper == 1.0:
		c.UpperBands = c.MiddleBands + c.StdDev.Sum()
		c.LowerBands = c.MiddleBands - (c.StdDev.Sum() * c.Lower)
	case c.Lower == 1.0:
		c.LowerBands = c.MiddleBands - c.StdDev.Sum()
		c.UpperBands = c.MiddleBands + (c.StdDev.Sum() * c.Upper)
	default:
		c.UpperBands = c.MiddleBands + (c.StdDev.Sum() * c.Upper)
		c.LowerBands = c.MiddleBands - (c.StdDev.Sum() * c.Lower)
	}
}

//Update Update the T3MA value of the current price
func (b *Bands) Update(price float64, date time.Time) {
	b.Calculator.setPrice(price, date)
	b.Calculator.MAObject.Update(price, date)
	b.Calculator.StdDev.Update(price, date)
	b.Calculator.calcBands()
}

//Sum Returns the Bands value of the current Bollinger object
func (b *Bands) Sum(flag string) float64 {
	switch flag {
	case "UPPER":
		return b.Calculator.UpperBands
	case "LOWER":
		return b.Calculator.LowerBands
	default:
		return b.Calculator.MiddleBands
	}
}
