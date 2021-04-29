package ma

import (
	"reflect"
	"testing"
	"time"
)

func TestT3MA_Sum(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *t3maCalculator
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestT3MA_Sum", fields{
			Value: 0,
			Calculator: &t3maCalculator{
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
			k := &T3MA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			if got := k.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestT3MA_Update(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *t3maCalculator
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
		{"TestT3MA_Update", fields{
			Value: 0,
			Calculator: &t3maCalculator{
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
			k := &T3MA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			k.Update(tt.args.price, tt.args.date)
		})
	}
}

func TestTicker_NewT3MA(t1 *testing.T) {
	type fields struct {
		Price float64
		Date  time.Time
	}
	type args struct {
		inTimePeriod int32
		factor       float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *T3MA
	}{
		{"TestTicker_NewT3MA", fields{
			Price: 0,
			Date:  time.Now(),
		}, args{inTimePeriod: 9, factor: 9}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticker{
				Price: tt.fields.Price,
				Date:  tt.fields.Date,
			}
			if got := t.NewT3MA(tt.args.inTimePeriod, tt.args.factor); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewT3MA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_t3maCalculator_calcT3MA(t *testing.T) {
	type fields struct {
		Ticker    *Ticker
		Period    int32
		Factor    float64
		Count     int32
		Result    float64
		K         float64
		OneMinusK float64
		E1        float64
		E2        float64
		E3        float64
		E4        float64
		E5        float64
		E6        float64
		C1        float64
		C2        float64
		C3        float64
		C4        float64
		E1Done    bool
		E2Done    bool
		E3Done    bool
		E4Done    bool
		E5Done    bool
		E6Done    bool
		Temp      []float64
		TempValue float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_t3maCalculator_calcT3MA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Period:    37,
			Factor:    9,
			Count:     100,
			E1Done:    true,
			E2Done:    true,
			E3Done:    false,
			E4Done:    false,
			E5Done:    false,
			E6Done:    false,
			Temp:      make([]float64, 35),
			TempValue: 0,
		}},
		{"Test_t3maCalculator_calcT3MA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Period:    37,
			Factor:    9,
			Count:     100,
			E1Done:    true,
			E2Done:    false,
			E3Done:    false,
			E4Done:    false,
			E5Done:    false,
			E6Done:    false,
			Temp:      make([]float64, 35),
			TempValue: 0,
		}},
		{"Test_t3maCalculator_calcT3MA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Period:    37,
			Factor:    9,
			Count:     100,
			E1Done:    true,
			E2Done:    true,
			E3Done:    true,
			E4Done:    false,
			E5Done:    false,
			E6Done:    false,
			Temp:      make([]float64, 35),
			TempValue: 0,
		}},
		{"Test_t3maCalculator_calcT3MA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Period:    37,
			Factor:    9,
			Count:     100,
			E1Done:    true,
			E2Done:    true,
			E3Done:    true,
			E4Done:    true,
			E5Done:    false,
			E6Done:    false,
			Temp:      make([]float64, 35),
			TempValue: 0,
		}},
		{"Test_t3maCalculator_calcT3MA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Period:    37,
			Factor:    9,
			Count:     100,
			E1Done:    true,
			E2Done:    true,
			E3Done:    true,
			E4Done:    true,
			E5Done:    true,
			E6Done:    false,
			Temp:      make([]float64, 35),
			TempValue: 0,
		}},
		{"Test_t3maCalculator_calcT3MA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Period:    37,
			Factor:    9,
			Count:     100,
			E1Done:    true,
			E2Done:    true,
			E3Done:    true,
			E4Done:    true,
			E5Done:    true,
			E6Done:    true,
			Temp:      make([]float64, 35),
			TempValue: 0,
		}},
		{"Test_t3maCalculator_calcT3MA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Period:    37,
			Factor:    9,
			Count:     100,
			E1Done:    false,
			E2Done:    true,
			E3Done:    true,
			E4Done:    true,
			E5Done:    true,
			E6Done:    true,
			Temp:      make([]float64, 35),
			TempValue: 0,
		}},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &t3maCalculator{
				Ticker:    tt.fields.Ticker,
				Period:    tt.fields.Period,
				Factor:    tt.fields.Factor,
				Count:     tt.fields.Count,
				Result:    tt.fields.Result,
				K:         tt.fields.K,
				OneMinusK: tt.fields.OneMinusK,
				E1:        tt.fields.E1,
				E2:        tt.fields.E2,
				E3:        tt.fields.E3,
				E4:        tt.fields.E4,
				E5:        tt.fields.E5,
				E6:        tt.fields.E6,
				C1:        tt.fields.C1,
				C2:        tt.fields.C2,
				C3:        tt.fields.C3,
				C4:        tt.fields.C4,
				E1Done:    tt.fields.E1Done,
				E2Done:    tt.fields.E2Done,
				E3Done:    tt.fields.E3Done,
				E4Done:    tt.fields.E4Done,
				E5Done:    tt.fields.E5Done,
				E6Done:    tt.fields.E6Done,
				Temp:      tt.fields.Temp,
				TempValue: tt.fields.TempValue,
			}
			c.calcT3MA()
		})
	}
}

func BenchmarkT3MA_Update(b *testing.B) {
	ticker := &Ticker{}
	t3ma := ticker.NewT3MA(9, 9)
	for n := 0; n < b.N; n++ {
		t3ma.Update(float64(n), time.Now())
	}
}
