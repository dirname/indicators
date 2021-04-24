package common

import (
	"reflect"
	"testing"
	"time"
)

func TestStdDev_Sum(t *testing.T) {
	type fields struct {
		Value    float64
		Dev      float64
		Variance *Variance
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestStdDev_Sum", fields{
			Value: 0,
			Dev:   0,
			Variance: &Variance{
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
			},
		}, 0},
		{"TestStdDev_Sum", fields{
			Value: 0,
			Dev:   0,
			Variance: &Variance{
				Value: 20,
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
			},
		}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StdDev{
				Value:    tt.fields.Value,
				Dev:      tt.fields.Dev,
				Variance: tt.fields.Variance,
			}
			if got := s.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTicker_NewStdDev(t1 *testing.T) {
	type fields struct {
		Price float64
		Date  time.Time
	}
	type args struct {
		inTimePeriod int32
		inDev        float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *StdDev
	}{
		{"TestTicker_NewStdDev", fields{
			Price: 0,
			Date:  time.Now(),
		}, args{
			inTimePeriod: 0,
			inDev:        0,
		}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticker{
				Price: tt.fields.Price,
				Date:  tt.fields.Date,
			}
			if got := t.NewStdDev(tt.args.inTimePeriod, tt.args.inDev); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewStdDev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkVariance_Sum(b *testing.B) {
	ticker := &Ticker{}
	StdDev := ticker.NewStdDev(14, 1.0)
	for n := 0; n < b.N; n++ {
		StdDev.Update(float64(n), time.Now())
	}
}
