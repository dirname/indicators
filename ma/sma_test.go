package ma

import (
	"reflect"
	"testing"
	"time"
)

func TestSMA_Sum(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *smaCalculator
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestSMA_Sum", fields{
			Value: 0,
			Calculator: &smaCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				Temp:        make([]float64, 10),
				Period:      0,
				PeriodTotal: 0,
				MA:          0,
			},
		}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			if got := m.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSMA_Update(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *smaCalculator
	}
	type args struct {
		price float64
		date  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestSMA_Update", fields{
			Value: 0,
			Calculator: &smaCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				Temp:        make([]float64, 10),
				Period:      0,
				PeriodTotal: 0,
				MA:          0,
			},
		}, args{
			price: 0,
			date:  time.Now(),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			m.Update(tt.args.price, tt.args.date)
		})
	}
}

func TestTicker_NewSMA(t1 *testing.T) {
	type fields struct {
		Price float64
		Date  time.Time
	}
	type args struct {
		inTimePeriod int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SMA
	}{
		{"TestTicker_NewSMA", fields{
			Price: 0,
			Date:  time.Now(),
		}, args{inTimePeriod: 10}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticker{
				Price: tt.fields.Price,
				Date:  tt.fields.Date,
			}
			if got := t.NewSMA(tt.args.inTimePeriod); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewSMA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smaCalculator_calcSMA(t *testing.T) {
	type fields struct {
		Ticker      *Ticker
		Temp        []float64
		Period      int32
		PeriodTotal float64
		MA          float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_smaCalculator_calcSMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Temp:        make([]float64, 10),
			Period:      0,
			PeriodTotal: 0,
			MA:          0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &smaCalculator{
				Ticker:      tt.fields.Ticker,
				Temp:        tt.fields.Temp,
				Period:      tt.fields.Period,
				PeriodTotal: tt.fields.PeriodTotal,
				MA:          tt.fields.MA,
			}
			s.calcSMA()
		})
	}
}

func BenchmarkSMA_Update(b *testing.B) {
	ticker := &Ticker{}
	rsi := ticker.NewSMA(14)
	for n := 0; n < b.N; n++ {
		rsi.Update(float64(n), time.Now())
	}
}
