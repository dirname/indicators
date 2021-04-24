package ma

import (
	"reflect"
	"testing"
	"time"
)

func TestDEMA_Sum(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *demaCalculator
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestDEMA_Sum", fields{
			Value: 0,
			Calculator: &demaCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				EMA: &EMA{
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
				},
				SecondEMA: &EMA{
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
				},
				FirstCount:  0,
				SecondCount: 0,
				Period:      0,
				Result:      0,
			},
		}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DEMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			if got := d.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDEMA_Update(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *demaCalculator
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
		{"TestDEMA_Update", fields{
			Value: 0,
			Calculator: &demaCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				EMA: &EMA{
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
				},
				SecondEMA: &EMA{
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
				},
				FirstCount:  0,
				SecondCount: 0,
				Period:      0,
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
			d := &DEMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			d.Update(tt.args.price, tt.args.date)
		})
	}
}

func TestTicker_NewDEMA(t1 *testing.T) {
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
		want   *DEMA
	}{
		{"TestTicker_NewDEMA", fields{
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
			if got := t.NewDEMA(tt.args.inTimePeriod); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewDEMA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_demaCalculator_calcDEMA(t *testing.T) {
	type fields struct {
		Ticker      *Ticker
		EMA         *EMA
		SecondEMA   *EMA
		FirstCount  int32
		SecondCount int32
		Period      int32
		Result      float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_demaCalculator_calcDEMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			EMA: &EMA{
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
			},
			SecondEMA: &EMA{
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
			},
			FirstCount:  0,
			SecondCount: 0,
			Period:      0,
			Result:      0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &demaCalculator{
				Ticker:      tt.fields.Ticker,
				EMA:         tt.fields.EMA,
				SecondEMA:   tt.fields.SecondEMA,
				FirstCount:  tt.fields.FirstCount,
				SecondCount: tt.fields.SecondCount,
				Period:      tt.fields.Period,
				Result:      tt.fields.Result,
			}
			s.calcDEMA()
		})
	}
}

func BenchmarkDEMA_Update(b *testing.B) {
	ticker := &Ticker{}
	dema := ticker.NewDEMA(9)
	for n := 0; n < b.N; n++ {
		dema.Update(float64(n), time.Now())
	}
}
