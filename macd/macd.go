package macd

import "time"

//Ticker ticker price information
type Ticker struct {
	Price float64
	Date  time.Time
}

//EMA EMA calculation object
type EMA struct {
	*Ticker
	Temp   float64
	Period int32
	Count  int32
	MA     float64
	K      float64
}

//MACD MACD Object
type MACD struct {
	slow   *EMA
	fast   *EMA
	signal *EMA
	DIF    float64
	DEM    float64
	OSC    float64
}

//Update Update the MACD value of the current price
func (m *MACD) Update(price float64, date time.Time) {
	lookBackSignal := m.signal.Period - 1
	lookBackTotal := lookBackSignal
	lookBackTotal += m.slow.Period - 1
	m.fast.setPrice(price, date)
	m.fast.calcEMA()
	m.slow.calcEMA()
	if m.fast.Count > lookBackTotal-1 {
		m.DIF = m.fast.MA - m.slow.MA
	}
	m.signal.setPrice(m.DIF, date)
	m.signal.calcEMA()
	m.DEM = m.signal.MA
	if m.fast.Count > lookBackTotal {
		m.OSC = m.DIF - m.signal.MA
	}
}

//Sum Returns the MACD value of the current MACD object
func (m *MACD) Sum(flag string) float64 {
	switch flag {
	case "DIF":
		return m.DIF
	case "DEM":
		return m.DEM
	default:
		return m.OSC
	}
}

//calcEMA calculate EMA
func (m *EMA) calcEMA() {
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

//setPrice set up price
func (t *Ticker) setPrice(price float64, date time.Time) {
	t.Price = price
	t.Date = date
}

//NewMACD factory function return a new object of MACD
func (t *Ticker) NewMACD(inFastPeriod, inSlowPeriod, inSignalPeriod int32) *MACD {
	if inSlowPeriod < inFastPeriod {
		inSlowPeriod, inFastPeriod = inFastPeriod, inSlowPeriod
	}
	slowEMA := &EMA{Ticker: t}
	fastEMA := &EMA{Ticker: t}
	signEMA := &EMA{Ticker: &Ticker{}}
	if inSlowPeriod != 0 {
		slowEMA.K = 2.0 / float64(inSlowPeriod+1)
	} else {
		inSlowPeriod = 26
		slowEMA.K = 0.075
	}
	if inFastPeriod != 0 {
		fastEMA.K = 2.0 / float64(inFastPeriod+1)
	} else {
		inFastPeriod = 12
		fastEMA.K = 0.15
	}
	signEMA.K = 2.0 / float64(inSignalPeriod+1)
	fastEMA.Period = inFastPeriod
	slowEMA.Period = inSlowPeriod
	signEMA.Period = inSignalPeriod
	return &MACD{
		slow:   slowEMA,
		fast:   fastEMA,
		signal: signEMA,
	}
}
