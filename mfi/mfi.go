package mfi

import (
	"time"
)

//Ticker ticker price information
type Ticker struct {
	Open  float64
	Close float64
	High  float64
	Low   float64
	Vol   float64
	Date  time.Time
}

//moneyFlow money flow
type moneyFlow struct {
	positive float64
	negative float64
}

//mfiCalculator MFI calculator
type mfiCalculator struct {
	*Ticker
	FlowIdx    int32
	MaxFlowIdx int32
	POSSumMF   float64
	NEGSumMF   float64
	PrevValue  float64
	Period     int32
	Count      int32
	TempValue  float64
	CalcValue  float64
	Result     float64
	MoneyFlow  []moneyFlow
}

//MFI MFI object
type MFI struct {
	Value      float64
	Calculator *mfiCalculator
}

//NewMFI factory function return a new object of MFI
func (t *Ticker) NewMFI(inTimePeriod int32) *MFI {
	calculator := &mfiCalculator{
		Ticker:     t,
		FlowIdx:    0,
		MaxFlowIdx: inTimePeriod - 1,
		Period:     inTimePeriod,
		MoneyFlow:  make([]moneyFlow, inTimePeriod),
	}
	return &MFI{
		Calculator: calculator,
	}
}

//setMarket set up market data
func (t *Ticker) setMarket(open, close, high, low, vol float64, date time.Time) {
	t.Open = open
	t.Close = close
	t.High = high
	t.Low = low
	t.Vol = vol
	t.Date = date
}

//calcMFI calculate MFI
func (m *mfiCalculator) calcMFI() {
	if m.Count == 0 {
		m.PrevValue = (m.High + m.Low + m.Close) / 3.0
	}
	if m.Count > m.Period {
		m.POSSumMF -= m.MoneyFlow[m.FlowIdx].positive
		m.NEGSumMF -= m.MoneyFlow[m.FlowIdx].negative
	}
	m.TempValue = (m.High + m.Low + m.Close) / 3.0
	m.CalcValue = m.TempValue - m.PrevValue
	m.PrevValue = m.TempValue
	m.TempValue *= m.Vol
	switch {
	case m.CalcValue < 0:
		m.MoneyFlow[m.FlowIdx].negative = m.TempValue
		m.NEGSumMF += m.TempValue
		m.MoneyFlow[m.FlowIdx].positive = 0.0
	case m.CalcValue > 0:
		m.MoneyFlow[m.FlowIdx].positive = m.TempValue
		m.POSSumMF += m.TempValue
		m.MoneyFlow[m.FlowIdx].negative = 0.0
	default:
		m.MoneyFlow[m.FlowIdx].positive = 0.0
		m.MoneyFlow[m.FlowIdx].negative = 0.0
	}
	m.FlowIdx++
	if m.FlowIdx > m.MaxFlowIdx {
		m.FlowIdx = 0
	}
	if m.Count >= m.Period {
		m.TempValue = m.POSSumMF + m.NEGSumMF
		if m.TempValue < 1.0 {
			m.Result = 0.0
		} else {
			m.Result = 100 * (m.POSSumMF / m.TempValue)
		}
	}
	m.Count++
}

//Update Update the MFI value of the current price
func (r *MFI) Update(open, close, high, low, vol float64, date time.Time) {
	r.Calculator.setMarket(open, close, high, low, vol, date)
	r.Calculator.calcMFI()
	r.Value = r.Calculator.Result
}

//Sum Returns the MFI value of the current MFI object
func (r *MFI) Sum() float64 {
	return r.Value
}
