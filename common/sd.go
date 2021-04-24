package common

import (
	"math"
)

//StdDev standard deviation
type StdDev struct {
	Value float64
	Dev   float64
	*Variance
}

//NewStdDev new standard deviation
func (t *Ticker) NewStdDev(inTimePeriod int32, inDev float64) *StdDev {
	return &StdDev{Variance: t.NewVariance(inTimePeriod), Dev: inDev}
}

//Sum Returns the standard deviation value of the current  standard deviation object
func (s *StdDev) Sum() float64 {
	s.Value = 0.0
	if !(s.Variance.Value < 1e-14) {
		s.Value = math.Sqrt(s.Variance.Value) * s.Dev
	}
	return s.Value
}
