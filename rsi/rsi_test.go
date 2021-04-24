package rsi

import (
	"reflect"
	"testing"
	"time"
)

func TestRSI_Sum(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *rsiCalculator
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestRSI_Sum", fields{
			Value:      0,
			Calculator: nil,
		}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RSI{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			if got := r.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRSI_Update(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *rsiCalculator
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
		{"TestRSI_Update", fields{
			Value: 0,
			Calculator: &rsiCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				TempValue: 0,
				calcValue: 0,
				PervValue: 0,
				Period:    0,
				Count:     0,
				PervGain:  0,
				PervLoss:  0,
				Result:    0,
			},
		}, args{
			price: 0,
			date:  time.Now(),
		}},
		{"TestRSI_Update", fields{
			Value: 0,
			Calculator: &rsiCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				TempValue: 0.0,
				calcValue: 0,
				PervValue: 0,
				Period:    1,
				Count:     20,
				PervGain:  1,
				PervLoss:  -1,
				Result:    0,
			},
		}, args{
			price: 0,
			date:  time.Now(),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RSI{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			r.Update(tt.args.price, tt.args.date)
		})
	}
}

func TestTicker_NewRSI(t1 *testing.T) {
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
		want   *RSI
	}{
		{"TestTicker_NewRSI", fields{
			Price: 0,
			Date:  time.Now(),
		}, args{14}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticker{
				Price: tt.fields.Price,
				Date:  tt.fields.Date,
			}
			if got := t.NewRSI(tt.args.inTimePeriod); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewRSI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTicker_setPrice(t1 *testing.T) {
	type fields struct {
		Price float64
		Date  time.Time
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
		{"TestTicker_setPrice", fields{
			Price: 0,
			Date:  time.Now(),
		}, args{
			price: 0,
			date:  time.Now(),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticker{
				Price: tt.fields.Price,
				Date:  tt.fields.Date,
			}
			t.setPrice(tt.args.price, tt.args.date)
		})
	}
}

func Test_rsiCalculator_calcRSI(t *testing.T) {
	type fields struct {
		Ticker    *Ticker
		TempValue float64
		calcValue float64
		PervValue float64
		Period    int32
		Count     int32
		PervGain  float64
		PervLoss  float64
		Result    float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_rsiCalculator_calcRSI", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			TempValue: 0,
			calcValue: 0,
			PervValue: 0,
			Period:    0,
			Count:     0,
			PervGain:  0,
			PervLoss:  0,
			Result:    0,
		}},
		{"Test_rsiCalculator_calcRSI", fields{
			Ticker: &Ticker{
				Price: 100,
				Date:  time.Now(),
			},
			TempValue: 0,
			calcValue: -1,
			PervValue: 200,
			Period:    5,
			Count:     10,
			PervGain:  0,
			PervLoss:  0,
			Result:    0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rsiCalculator{
				Ticker:    tt.fields.Ticker,
				TempValue: tt.fields.TempValue,
				calcValue: tt.fields.calcValue,
				PervValue: tt.fields.PervValue,
				Period:    tt.fields.Period,
				Count:     tt.fields.Count,
				PervGain:  tt.fields.PervGain,
				PervLoss:  tt.fields.PervLoss,
				Result:    tt.fields.Result,
			}
			r.calcRSI()
		})
	}
}

func BenchmarkRSI_Update(b *testing.B) {
	ticker := &Ticker{}
	rsi := ticker.NewRSI(14)
	for n := 0; n < b.N; n++ {
		rsi.Update(float64(n), time.Now())
	}
}
