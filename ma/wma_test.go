package ma

import (
	"reflect"
	"testing"
	"time"
)

func TestTicker_NewWMA(t1 *testing.T) {
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
		want   *WMA
	}{
		{"TestTicker_NewWMA", fields{
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
			if got := t.NewWMA(tt.args.inTimePeriod); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewWMA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWMA_Sum(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *wmaCalculator
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestWMA_Sum", fields{
			Value: 0,
			Calculator: &wmaCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				Divider:       0,
				PeriodSum:     0,
				PeriodSub:     0,
				TrailingValue: 0,
				Temp:          make([]float64, 0, 10),
				Period:        0,
				Result:        0,
			},
		}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			if got := w.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWMA_Update(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *wmaCalculator
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
		{"TestWMA_Update", fields{
			Value: 0,
			Calculator: &wmaCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				Divider:       0,
				PeriodSum:     0,
				PeriodSub:     0,
				TrailingValue: 0,
				Temp:          make([]float64, 0, 10),
				Period:        0,
				Result:        0,
			},
		}, args{
			price: 0,
			date:  time.Now(),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			w.Update(tt.args.price, tt.args.date)
		})
	}
}

func Test_wmaCalculator_calcWMA(t *testing.T) {
	type fields struct {
		Ticker        *Ticker
		Divider       int32
		PeriodSum     float64
		PeriodSub     float64
		TrailingValue float64
		Count         int32
		Temp          []float64
		TempValue     float64
		Period        int32
		PeriodTotal   float64
		Result        float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_wmaCalculator_calcWMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Divider:       0,
			PeriodSum:     0,
			PeriodSub:     0,
			TrailingValue: 0,
			Count:         0,
			Temp:          make([]float64, 0, 10),
			TempValue:     0,
			Period:        0,
			PeriodTotal:   0,
			Result:        0,
		}},
		{"Test_wmaCalculator_calcWMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Divider:       0,
			PeriodSum:     0,
			PeriodSub:     0,
			TrailingValue: 0,
			Count:         0,
			Temp:          make([]float64, 0, 10),
			TempValue:     0,
			Period:        10,
			PeriodTotal:   0,
			Result:        0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &wmaCalculator{
				Ticker:        tt.fields.Ticker,
				Divider:       tt.fields.Divider,
				PeriodSum:     tt.fields.PeriodSum,
				PeriodSub:     tt.fields.PeriodSub,
				TrailingValue: tt.fields.TrailingValue,
				Temp:          tt.fields.Temp,
				Period:        tt.fields.Period,
				Result:        tt.fields.Result,
			}
			s.calcWMA()
		})
	}
}

func BenchmarkWMA_Update(b *testing.B) {
	ticker := &Ticker{}
	wma := ticker.NewWMA(9)
	for n := 0; n < b.N; n++ {
		wma.Update(float64(n), time.Now())
	}
}
