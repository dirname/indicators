package ma

import (
	"math"
	"time"
)

//mamaCalculator MAMA
type mamaCalculator struct {
	*Ticker
	Odd, Even, Q1Odd, Q1Even, JIOdd, JIEven, JQOdd, JQEven, WMAValue                                                    []float64
	Deg, WMASub, WMASum, TrailingWMA, SmoothValue, Adjusted, Q1Value, Q2Value, I1Value, I2Value                         float64
	Hilbert, ParamA, ParamB, detRender, PrevDetEven, PrevDetOdd, PrevInputEven, PrevQ1Odd, PrevQ1InputOdd, PrevInputOdd float64
	PrevJIInput, PrevJQInput, PervQ1Even, EvenPrev3, EvenPrev2, OddPrev3, OddPrev2, PrevJIEven, PrevJQEven              float64
	PrevJIOdd, PrevJIInputOdd, PrevJQInputOdd, PrevJQOdd, PrevQ2, PrevI2, JI, JQ, Temp, Cache                           float64
	PrevPhase, FastLimit, SlowLimit, MAMA, FAMA, Period, Re, Im, PrevQ1InputEven                                        float64
	Count, HilbertIdx                                                                                                   int32
}

//MAMA MAMA
type MAMA struct {
	Value      float64
	Calculator *mamaCalculator
}

//NewMAMA new MAMA
func (t *Ticker) NewMAMA(inFast, inSlow float64) *MAMA {
	calculator := &mamaCalculator{
		Ticker:    t,
		FastLimit: inFast,
		SlowLimit: inSlow,
		Deg:       180.0 / (4.0 * math.Atan(1)),
		Odd:       make([]float64, 3, 3),
		Even:      make([]float64, 3, 3),
		Q1Even:    make([]float64, 3, 3),
		Q1Odd:     make([]float64, 3, 3),
		JIEven:    make([]float64, 3, 3),
		JIOdd:     make([]float64, 3, 3),
		JQEven:    make([]float64, 3, 3),
		JQOdd:     make([]float64, 3, 3),
		WMAValue:  make([]float64, 0, 5),
		ParamA:    0.0962,
		ParamB:    0.5769,
	}
	return &MAMA{
		Calculator: calculator,
	}
}

//calcMAMA calculate MAMA
func (s *mamaCalculator) calcMAMA() {
	s.WMAValue = append(s.WMAValue, s.Price)
	switch {
	case s.Count == 0:
		s.WMASub = s.Price
		s.WMASum = s.Price
	case s.Count == 1:
		s.WMASum += s.Price * 2.0
		s.WMASub += s.Price
	case s.Count == 2:
		s.WMASub += s.Price
		s.WMASum += s.Price * 3.0
	case s.Count >= 3:
		s.WMASub += s.Price
		s.WMASub -= s.TrailingWMA
		s.WMASum += s.Price * 4.0
		s.TrailingWMA = s.WMAValue[0]
		s.WMAValue = s.WMAValue[1:]
		s.SmoothValue = s.WMASum * 0.1
		s.WMASum -= s.WMASub
	}
	if s.Count >= 12 {
		s.Adjusted = (0.075 * s.Period) + 0.54
		s.Q2Value, s.I2Value, s.Cache = 0.0, 0.0, 0.0
		if s.Count&1 == 0 {
			s.isPowerOfTwo()
		} else {
			s.notPowerOfTwo()
		}
		s.sum()
	}
	s.Count++
}

//sum sum MAMA and FAMA
func (s *mamaCalculator) sum() {
	s.sumOne()
	s.sumTwo()
	s.sumThree()
}

//calcMAMA day is a power of 2
func (s *mamaCalculator) isPowerOfTwo() {
	s.Hilbert = s.ParamA * s.SmoothValue
	s.detRender = -s.Even[s.HilbertIdx]
	s.Even[s.HilbertIdx] = s.Hilbert
	s.detRender += s.Hilbert
	s.detRender -= s.PrevDetEven
	s.PrevDetEven = s.ParamB * s.PrevInputEven
	s.detRender += s.PrevDetEven
	s.PrevInputEven = s.SmoothValue
	s.detRender *= s.Adjusted

	s.Hilbert = s.ParamA * s.detRender
	s.Q1Value = -s.Q1Even[s.HilbertIdx]
	s.Q1Even[s.HilbertIdx] = s.Hilbert
	s.Q1Value += s.Hilbert
	s.Q1Value -= s.PervQ1Even
	s.PervQ1Even = s.ParamB * s.PrevQ1InputEven
	s.Q1Value += s.PervQ1Even
	s.PrevQ1InputEven = s.detRender
	s.Q1Value *= s.Adjusted

	s.Hilbert = s.ParamA * s.EvenPrev3
	s.JI = -s.JIEven[s.HilbertIdx]
	s.JIEven[s.HilbertIdx] = s.Hilbert
	s.JI += s.Hilbert
	s.JI -= s.PrevJIEven
	s.PrevJIEven = s.ParamB * s.PrevJIInput
	s.JI += s.PrevJIEven
	s.PrevJIInput = s.EvenPrev3
	s.JI *= s.Adjusted

	s.Hilbert = s.ParamA * s.Q1Value
	s.JQ = -s.JQEven[s.HilbertIdx]
	s.JQEven[s.HilbertIdx] = s.Hilbert
	s.JQ += s.Hilbert
	s.JQ -= s.PrevJQEven
	s.PrevJQEven = s.ParamB * s.PrevJQInput
	s.JQ += s.PrevJQEven
	s.PrevJQInput = s.Q1Value
	s.JQ *= s.Adjusted
	s.HilbertIdx++
	if s.HilbertIdx == 3 {
		s.HilbertIdx = 0
	}
	s.Q2Value = (0.2 * (s.Q1Value + s.JI)) + (0.8 * s.PrevQ2)
	s.I2Value = (0.2 * (s.EvenPrev3 - s.JQ)) + (0.8 * s.PrevI2)
	s.OddPrev3 = s.OddPrev2
	s.OddPrev2 = s.detRender
	if s.EvenPrev3 != 0.0 {
		s.Cache = math.Atan(s.Q1Value/s.EvenPrev3) * s.Deg
	} else {
		s.Cache = 0
	}
}

//calcMAMA day not a power of 2
func (s *mamaCalculator) notPowerOfTwo() {
	s.Hilbert = s.ParamA * s.SmoothValue
	s.detRender = -s.Odd[s.HilbertIdx]
	s.Odd[s.HilbertIdx] = s.Hilbert
	s.detRender += s.Hilbert
	s.detRender -= s.PrevDetOdd
	s.PrevDetOdd = s.ParamB * s.PrevInputOdd
	s.detRender += s.PrevDetOdd
	s.PrevInputOdd = s.SmoothValue
	s.detRender *= s.Adjusted

	s.Hilbert = s.ParamA * s.detRender
	s.Q1Value = -s.Q1Odd[s.HilbertIdx]
	s.Q1Odd[s.HilbertIdx] = s.Hilbert
	s.Q1Value += s.Hilbert
	s.Q1Value -= s.PrevQ1Odd
	s.PrevQ1Odd = s.ParamB * s.PrevQ1InputOdd
	s.Q1Value += s.PrevQ1Odd
	s.PrevQ1InputOdd = s.detRender
	s.Q1Value *= s.Adjusted

	s.Hilbert = s.ParamA * s.OddPrev3
	s.JI = -s.JIOdd[s.HilbertIdx]
	s.JIOdd[s.HilbertIdx] = s.Hilbert
	s.JI += s.Hilbert
	s.JI -= s.PrevJIOdd
	s.PrevJIOdd = s.ParamB * s.PrevJIInputOdd
	s.JI += s.PrevJIOdd
	s.PrevJIInputOdd = s.OddPrev3
	s.JI *= s.Adjusted

	s.Hilbert = s.ParamA * s.Q1Value
	s.JQ = -s.JQOdd[s.HilbertIdx]
	s.JQOdd[s.HilbertIdx] = s.Hilbert
	s.JQ += s.Hilbert
	s.JQ -= s.PrevJQOdd
	s.PrevJQOdd = s.ParamB * s.PrevJQInputOdd
	s.JQ += s.PrevJQOdd
	s.PrevJQInputOdd = s.Q1Value
	s.JQ *= s.Adjusted

	s.Q2Value = (0.2 * (s.Q1Value + s.JI)) + (0.8 * s.PrevQ2)
	s.I2Value = (0.2 * (s.OddPrev3 - s.JQ)) + (0.8 * s.PrevI2)
	s.EvenPrev3 = s.EvenPrev2
	s.EvenPrev2 = s.detRender
	if s.OddPrev3 != 0.0 {
		s.Cache = math.Atan(s.Q1Value/s.OddPrev3) * s.Deg
	} else {
		s.Cache = 0.0
	}
}

//sumOne set up one of sum
func (s *mamaCalculator) sumOne() {
	s.Temp = s.PrevPhase - s.Cache
	s.PrevPhase = s.Cache
	if s.Temp < 1.0 {
		s.Temp = 1.0
	}
	if s.Temp > 1.0 {
		s.Temp = s.FastLimit / s.Temp
		if s.Temp < s.SlowLimit {
			s.Temp = s.SlowLimit
		}
	} else {
		s.Temp = s.FastLimit
	}
	s.MAMA = (s.Temp * s.Price) + ((1 - s.Temp) * s.MAMA)
	s.Temp *= 0.5
	s.FAMA = (s.Temp * s.MAMA) + ((1 - s.Temp) * s.FAMA)
}

//sumTwo set up two of sum
func (s *mamaCalculator) sumTwo() {
	s.Re = (0.2 * ((s.I2Value * s.PrevI2) + (s.Q2Value * s.PrevQ2))) + (0.8 * s.Re)
	s.Im = (0.2 * ((s.I2Value * s.PrevQ2) - (s.Q2Value * s.PrevI2))) + (0.8 * s.Im)
	s.PrevQ2 = s.Q2Value
	s.PrevI2 = s.I2Value
	s.Temp = s.Period
	if (s.Im != 0.0) && (s.Re != 0.0) {
		s.Period = 360.0 / (math.Atan(s.Im/s.Re) * s.Deg)
	}
	s.Cache = 1.5 * s.Temp
	if s.Period > s.Cache {
		s.Period = s.Cache
	}
	s.Cache = 0.67 * s.Temp
	if s.Period < s.Cache {
		s.Period = s.Cache
	}
}

//sumThree set up three of sum
func (s *mamaCalculator) sumThree() {
	if s.Period < 6 {
		s.Period = 6
	} else if s.Period > 50 {
		s.Period = 50
	}
	s.Period = (0.2 * s.Period) + (0.8 * s.Temp)
}

//Update Update the MAMA value of the current price
func (d *MAMA) Update(price float64, date time.Time) {
	d.Calculator.setPrice(price, date)
	d.Calculator.calcMAMA()
}

//Sum Returns the MAMA value of the current MAMA object
func (d *MAMA) Sum() (float64, float64) {
	return d.Calculator.MAMA, d.Calculator.FAMA
}
