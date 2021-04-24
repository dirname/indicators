package ma

import (
	"reflect"
	"testing"
	"time"
)

func TestTEMA_Sum(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *temaCalculator
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"TestTEMA_Sum", fields{
			Value: 0,
			Calculator: &temaCalculator{
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
				ThirdEMA: &EMA{
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
				ThirdCount:  0,
				Period:      0,
				Result:      0,
			},
		}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &TEMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			if got := d.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTEMA_Update(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *temaCalculator
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
		{"TestTEMA_Update", fields{
			Value: 0,
			Calculator: &temaCalculator{
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
				ThirdEMA: &EMA{
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
				ThirdCount:  0,
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
			d := &TEMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			d.Update(tt.args.price, tt.args.date)
		})
	}
}

func TestTicker_NewTEMA(t1 *testing.T) {
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
		want   *TEMA
	}{
		{"TestTicker_NewTEMA", fields{
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
			if got := t.NewTEMA(tt.args.inTimePeriod); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewTEMA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_temaCalculator_calcTEMA(t *testing.T) {
	type fields struct {
		Ticker      *Ticker
		EMA         *EMA
		SecondEMA   *EMA
		ThirdEMA    *EMA
		FirstCount  int32
		SecondCount int32
		ThirdCount  int32
		Period      int32
		Result      float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_temaCalculator_calcTEMA", fields{
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
			ThirdEMA: &EMA{
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
			ThirdCount:  0,
			Period:      0,
			Result:      0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &temaCalculator{
				Ticker:      tt.fields.Ticker,
				EMA:         tt.fields.EMA,
				SecondEMA:   tt.fields.SecondEMA,
				ThirdEMA:    tt.fields.ThirdEMA,
				FirstCount:  tt.fields.FirstCount,
				SecondCount: tt.fields.SecondCount,
				ThirdCount:  tt.fields.ThirdCount,
				Period:      tt.fields.Period,
				Result:      tt.fields.Result,
			}
			s.calcTEMA()
		})
	}
}

func BenchmarkTEMA_Update(b *testing.B) {
	ticker := &Ticker{}
	tema := ticker.NewTEMA(9)
	for n := 0; n < b.N; n++ {
		tema.Update(float64(n), time.Now())
	}
}
