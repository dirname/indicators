package mfi

import (
	"reflect"
	"testing"
	"time"
)

func TestMFI_Sum(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *mfiCalculator
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestMFI_Sum", fields{
			Value:      0,
			Calculator: nil,
		}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MFI{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			if got := r.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMFI_Update(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *mfiCalculator
	}
	type args struct {
		open  float64
		close float64
		high  float64
		low   float64
		vol   float64
		date  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestMFI_Update", fields{
			Value: 0,
			Calculator: &mfiCalculator{
				Ticker: &Ticker{
					Open:  0,
					Close: 0,
					High:  0,
					Low:   0,
					Vol:   0,
					Date:  time.Now(),
				},
				FlowIdx:    0,
				MaxFlowIdx: 10,
				POSSumMF:   0,
				NEGSumMF:   0,
				PrevValue:  12,
				Period:     0,
				Count:      0,
				TempValue:  0,
				CalcValue:  0,
				Result:     0,
				MoneyFlow:  make([]moneyFlow, 10),
			},
		}, args{
			open:  0,
			close: 0,
			high:  0,
			low:   0,
			vol:   0,
			date:  time.Now(),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MFI{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			r.Update(tt.args.open, tt.args.close, tt.args.high, tt.args.low, tt.args.vol, tt.args.date)
		})
	}
}

func TestTicker_NewMFI(t1 *testing.T) {
	type fields struct {
		Open  float64
		Close float64
		High  float64
		Low   float64
		Vol   float64
		Date  time.Time
	}
	type args struct {
		inTimePeriod int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MFI
	}{
		{"TestTicker_NewMFI", fields{
			Open:  0,
			Close: 0,
			High:  0,
			Low:   0,
			Vol:   0,
			Date:  time.Now(),
		}, args{14}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticker{
				Open:  tt.fields.Open,
				Close: tt.fields.Close,
				High:  tt.fields.High,
				Low:   tt.fields.Low,
				Vol:   tt.fields.Vol,
				Date:  tt.fields.Date,
			}
			if got := t.NewMFI(tt.args.inTimePeriod); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewMFI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTicker_setMarket(t1 *testing.T) {
	type fields struct {
		Open  float64
		Close float64
		High  float64
		Low   float64
		Vol   float64
		Date  time.Time
	}
	type args struct {
		open  float64
		close float64
		high  float64
		low   float64
		vol   float64
		date  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestTicker_setMarket", fields{
			Open:  0,
			Close: 0,
			High:  0,
			Low:   0,
			Vol:   0,
			Date:  time.Now(),
		}, args{
			open:  0,
			close: 0,
			high:  0,
			low:   0,
			vol:   0,
			date:  time.Now(),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticker{
				Open:  tt.fields.Open,
				Close: tt.fields.Close,
				High:  tt.fields.High,
				Low:   tt.fields.Low,
				Vol:   tt.fields.Vol,
				Date:  tt.fields.Date,
			}
			t.setMarket(tt.args.open, tt.args.close, tt.args.high, tt.args.low, tt.args.vol, tt.args.date)
		})
	}
}

func Test_mfiCalculator_calcMFI(t *testing.T) {
	type fields struct {
		Ticker     *Ticker
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
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_mfiCalculator_calcMFI", fields{
			Ticker: &Ticker{
				Open:  0,
				Close: 0,
				High:  0,
				Low:   0,
				Vol:   0,
				Date:  time.Now(),
			},
			FlowIdx:    0,
			MaxFlowIdx: 10,
			POSSumMF:   0,
			NEGSumMF:   0,
			PrevValue:  0,
			Period:     10,
			Count:      0,
			TempValue:  0,
			CalcValue:  0,
			Result:     0,
			MoneyFlow:  make([]moneyFlow, 10),
		}},
		{"Test_mfiCalculator_calcMFI", fields{
			Ticker: &Ticker{
				Open:  0,
				Close: 0,
				High:  0,
				Low:   0,
				Vol:   0,
				Date:  time.Now(),
			},
			FlowIdx:    0,
			MaxFlowIdx: 10,
			POSSumMF:   0,
			NEGSumMF:   0,
			PrevValue:  10,
			Period:     10,
			Count:      20,
			TempValue:  0,
			CalcValue:  0,
			Result:     0,
			MoneyFlow:  make([]moneyFlow, 10),
		}},
		{"Test_mfiCalculator_calcMFI", fields{
			Ticker: &Ticker{
				Open:  0,
				Close: 0,
				High:  0,
				Low:   0,
				Vol:   0,
				Date:  time.Now(),
			},
			FlowIdx:    9,
			MaxFlowIdx: 9,
			POSSumMF:   20,
			NEGSumMF:   0,
			PrevValue:  1,
			Period:     10,
			Count:      20,
			TempValue:  20,
			CalcValue:  0,
			Result:     0,
			MoneyFlow:  make([]moneyFlow, 10),
		}},
		{"Test_mfiCalculator_calcMFI", fields{
			Ticker: &Ticker{
				Open:  20,
				Close: 30,
				High:  40,
				Low:   50,
				Vol:   60,
				Date:  time.Now(),
			},
			FlowIdx:    0,
			MaxFlowIdx: 10,
			POSSumMF:   0,
			NEGSumMF:   0,
			PrevValue:  10,
			Period:     10,
			Count:      20,
			TempValue:  40,
			CalcValue:  0,
			Result:     0,
			MoneyFlow:  make([]moneyFlow, 10),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mfiCalculator{
				Ticker:     tt.fields.Ticker,
				FlowIdx:    tt.fields.FlowIdx,
				MaxFlowIdx: tt.fields.MaxFlowIdx,
				POSSumMF:   tt.fields.POSSumMF,
				NEGSumMF:   tt.fields.NEGSumMF,
				PrevValue:  tt.fields.PrevValue,
				Period:     tt.fields.Period,
				Count:      tt.fields.Count,
				TempValue:  tt.fields.TempValue,
				CalcValue:  tt.fields.CalcValue,
				Result:     tt.fields.Result,
				MoneyFlow:  tt.fields.MoneyFlow,
			}
			m.calcMFI()
		})
	}
}

func BenchmarkMFI_Update(b *testing.B) {
	ticker := &Ticker{}
	rsi := ticker.NewMFI(14)
	for n := 0; n < b.N; n++ {
		v := float64(n)
		rsi.Update(v, v, v, v, v, time.Now())
	}
}
