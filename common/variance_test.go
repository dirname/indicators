package common

import (
	"reflect"
	"testing"
	"time"
)

func TestTicker_NewVariance(t1 *testing.T) {
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
		want   *Variance
	}{
		{"TestTicker_NewVariance", fields{
			Price: 0,
			Date:  time.Now(),
		}, args{10}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticker{
				Price: tt.fields.Price,
				Date:  tt.fields.Date,
			}
			if got := t.NewVariance(tt.args.inTimePeriod); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewVariance() = %v, want %v", got, tt.want)
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

func TestVariance_Sum(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *varianceCalculator
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestVariance_Sum", fields{
			Value: 0,
			Calculator: &varianceCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				Temp:        make([]float64, 0, 10),
				TempValue:   0,
				Period:      0,
				PeriodTotal: 0,
				Result:      0,
			},
		}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Variance{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			if got := v.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariance_Update(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *varianceCalculator
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
		{"TestVariance_Update", fields{
			Value: 0,
			Calculator: &varianceCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				Temp:        make([]float64, 0, 10),
				TempValue:   0,
				Period:      0,
				PeriodTotal: 0,
				Result:      0,
			},
		}, args{
			price: 0,
			date:  time.Now(),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Variance{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			v.Update(tt.args.price, tt.args.date)
		})
	}
}

func Test_varianceCalculator_calcVariance(t *testing.T) {
	type fields struct {
		Ticker      *Ticker
		Temp        []float64
		TempValue   float64
		Period      int32
		PeriodTotal float64
		Result      float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_varianceCalculator_calcVariance", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Temp:        make([]float64, 0, 10),
			TempValue:   0,
			Period:      0,
			PeriodTotal: 0,
			Result:      0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &varianceCalculator{
				Ticker:      tt.fields.Ticker,
				Temp:        tt.fields.Temp,
				TempValue:   tt.fields.TempValue,
				Period:      tt.fields.Period,
				PeriodTotal: tt.fields.PeriodTotal,
				Result:      tt.fields.Result,
			}
			s.calcVariance()
		})
	}
}

func BenchmarkVariance_Update(b *testing.B) {
	ticker := &Ticker{}
	Var := ticker.NewVariance(14)
	for n := 0; n < b.N; n++ {
		Var.Update(float64(n), time.Now())
	}
}
