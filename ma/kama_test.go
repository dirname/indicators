package ma

import (
	"reflect"
	"testing"
	"time"
)

func TestKAMA_Sum(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *kamaCalculator
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestKAMA_Sum", fields{
			Value: 0,
			Calculator: &kamaCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
			},
		}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KAMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			if got := k.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKAMA_Update(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *kamaCalculator
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
		{"TestKAMA_Update", fields{
			Value: 0,
			Calculator: &kamaCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				Temp:  make([]float64, 3),
				Count: 1,
			},
		}, args{
			price: 0,
			date:  time.Now(),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KAMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			k.Update(tt.args.price, tt.args.date)
		})
	}
}

func TestKamaCalculator_calcKAMA(t *testing.T) {
	type fields struct {
		Ticker        *Ticker
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
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestKamaCalculator_calcKAMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Count:  2,
			Period: 8,
			Temp:   make([]float64, 8),
		}},
		{"TestKamaCalculator_calcKAMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Count:  3,
			Period: 3,
			Temp:   make([]float64, 8),
		}},
		{"TestKamaCalculator_calcKAMA", fields{
			Ticker: &Ticker{
				Price: 1000,
				Date:  time.Now(),
			},
			Count:  4,
			Period: 3,
			Temp:   make([]float64, 8),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &kamaCalculator{
				Ticker:        tt.fields.Ticker,
				Period:        tt.fields.Period,
				Count:         tt.fields.Count,
				Result:        tt.fields.Result,
				Max:           tt.fields.Max,
				Diff:          tt.fields.Diff,
				SumROC:        tt.fields.SumROC,
				Temp:          tt.fields.Temp,
				PrevKama:      tt.fields.PrevKama,
				PeriodROC:     tt.fields.PeriodROC,
				TrailingValue: tt.fields.TrailingValue,
				TempValue:     tt.fields.TempValue,
			}
			c.calcKAMA()
		})
	}
}

func TestTicker_NewKAMA(t1 *testing.T) {
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
		want   *KAMA
	}{
		{"TestTicker_NewKAMA", fields{
			Price: 0,
			Date:  time.Now(),
		}, args{0}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticker{
				Price: tt.fields.Price,
				Date:  tt.fields.Date,
			}
			if got := t.NewKAMA(tt.args.inTimePeriod); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewKAMA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkKAMA_Update(b *testing.B) {
	ticker := &Ticker{}
	kama := ticker.NewKAMA(9)
	for n := 0; n < b.N; n++ {
		kama.Update(float64(n), time.Now())
	}
}
