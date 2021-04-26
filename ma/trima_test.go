package ma

import (
	"reflect"
	"testing"
	"time"
)

func TestTRIMA_Sum(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *trimaCalculator
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestTRIMA_Sum", fields{
			Value: 0,
			Calculator: &trimaCalculator{
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
			d := &TRIMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			if got := d.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTRIMA_Update(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *trimaCalculator
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
		{"", fields{
			Value: 0,
			Calculator: &trimaCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
			},
		}, args{
			price: 0,
			date:  time.Now(),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &TRIMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			d.Update(tt.args.price, tt.args.date)
		})
	}
}

func TestTicker_NewTRIMA(t1 *testing.T) {
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
		want   *TRIMA
	}{
		{"TestTicker_NewTRIMA", fields{
			Price: 0,
			Date:  time.Now(),
		}, args{inTimePeriod: 0}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticker{
				Price: tt.fields.Price,
				Date:  tt.fields.Date,
			}
			if got := t.NewTRIMA(tt.args.inTimePeriod); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewTRIMA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trimaCalculator_calcTRIMA(t1 *testing.T) {
	type fields struct {
		Ticker       *Ticker
		Period       int32
		Count        int32
		Factor       float64
		Idx          float64
		MiddleIdx    int32
		TrailingIdx  int32
		Temp         []float64
		NumeratorSub float64
		NumeratorAdd float64
		Numerator    float64
		SubDone      bool
		AddDone      bool
		LastValue    float64
		Result       float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_trimaCalculator_calcTRIMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Period: 8,
		}},
		{"Test_trimaCalculator_calcTRIMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Period: 9,
		}},
		{"Test_trimaCalculator_calcTRIMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			TrailingIdx: 1,
			MiddleIdx:   4,
			Idx:         4,
			Temp:        make([]float64, 9, 10),
			Count:       10,
			Period:      9,
		}},
		{"Test_trimaCalculator_calcTRIMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Count:       8,
			AddDone:     true,
			SubDone:     true,
			TrailingIdx: 1,
			MiddleIdx:   4,
			Period:      9,
			Idx:         4,
			Temp:        make([]float64, 18),
		}},
		{"Test_trimaCalculator_calcTRIMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Count:       8,
			AddDone:     true,
			SubDone:     true,
			TrailingIdx: 1,
			MiddleIdx:   4,
			Period:      8,
			Idx:         4,
			Temp:        make([]float64, 18),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &trimaCalculator{
				Ticker:       tt.fields.Ticker,
				Period:       tt.fields.Period,
				Count:        tt.fields.Count,
				Factor:       tt.fields.Factor,
				Idx:          tt.fields.Idx,
				MiddleIdx:    tt.fields.MiddleIdx,
				TrailingIdx:  tt.fields.TrailingIdx,
				Temp:         tt.fields.Temp,
				NumeratorSub: tt.fields.NumeratorSub,
				NumeratorAdd: tt.fields.NumeratorAdd,
				Numerator:    tt.fields.Numerator,
				SubDone:      tt.fields.SubDone,
				AddDone:      tt.fields.AddDone,
				LastValue:    tt.fields.LastValue,
				Result:       tt.fields.Result,
			}
			t.calcTRIMA()
		})
	}
}

func BenchmarkTRIMA_Update(b *testing.B) {
	ticker := &Ticker{}
	trima := ticker.NewTRIMA(9)
	for n := 0; n < b.N; n++ {
		trima.Update(float64(n), time.Now())
	}
}
