package ma

import (
	"reflect"
	"testing"
	"time"
)

func TestEMA_Sum(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *emaCalculator
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestEMA_Sum", fields{
			Value: 0,
			Calculator: &emaCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				Temp:   0,
				Period: 0,
				Count:  0,
				MA:     0,
				K:      0,
			},
		}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &EMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			if got := m.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEMA_Update(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *emaCalculator
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
		{"TestEMA_Update", fields{
			Value: 0,
			Calculator: &emaCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				Temp:   0,
				Period: 0,
				Count:  0,
				MA:     0,
				K:      0,
			},
		}, args{
			price: 0,
			date:  time.Now(),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &EMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			m.Update(tt.args.price, tt.args.date)
		})
	}
}

func TestTicker_NewEMA(t1 *testing.T) {
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
		want   *EMA
	}{
		{"TestTicker_NewEMA", fields{
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
			if got := t.NewEMA(tt.args.inTimePeriod); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewEMA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_emaCalculator_calcEMA(t *testing.T) {
	type fields struct {
		Ticker *Ticker
		Temp   float64
		Period int32
		Count  int32
		MA     float64
		K      float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_emaCalculator_calcEMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Temp:   0,
			Period: 0,
			Count:  0,
			MA:     0,
			K:      0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &emaCalculator{
				Ticker: tt.fields.Ticker,
				Temp:   tt.fields.Temp,
				Period: tt.fields.Period,
				Count:  tt.fields.Count,
				MA:     tt.fields.MA,
				K:      tt.fields.K,
			}
			m.calcEMA()
		})
	}
}

func BenchmarkEMA_Update(b *testing.B) {
	ticker := &Ticker{}
	ema := ticker.NewEMA(9)
	for n := 0; n < b.N; n++ {
		ema.Update(float64(n), time.Now())
	}
}
