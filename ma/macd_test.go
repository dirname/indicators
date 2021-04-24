package ma

import (
	"reflect"
	"testing"
	"time"
)

func TestEMA_calcEMA(t *testing.T) {
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
		{"TestEMA_calcEMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Temp:   0,
			Period: 1,
			Count:  0,
			MA:     0,
			K:      0,
		}},
		{"TestEMA_calcEMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Temp:   0,
			Period: 1,
			Count:  2,
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

func TestMACD_Sum(t *testing.T) {
	type fields struct {
		slow   *emaCalculator
		fast   *emaCalculator
		signal *emaCalculator
		DIF    float64
		DEM    float64
		OSC    float64
	}
	type args struct {
		flag string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"TestMACD_Sum", fields{
			slow:   nil,
			fast:   nil,
			signal: nil,
			DIF:    0,
			DEM:    0,
			OSC:    0,
		}, args{flag: "DIF"}, 0},
		{"TestMACD_Sum", fields{
			slow:   nil,
			fast:   nil,
			signal: nil,
			DIF:    0,
			DEM:    0,
			OSC:    0,
		}, args{flag: "DEM"}, 0},
		{"TestMACD_Sum", fields{
			slow:   nil,
			fast:   nil,
			signal: nil,
			DIF:    0,
			DEM:    0,
			OSC:    0,
		}, args{flag: "OSC"}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MACD{
				slow:   tt.fields.slow,
				fast:   tt.fields.fast,
				signal: tt.fields.signal,
				DIF:    tt.fields.DIF,
				DEM:    tt.fields.DEM,
				OSC:    tt.fields.OSC,
			}
			if got := m.Sum(tt.args.flag); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMACD_Update(t *testing.T) {
	type fields struct {
		slow   *emaCalculator
		fast   *emaCalculator
		signal *emaCalculator
		DIF    float64
		DEM    float64
		OSC    float64
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
		{"TestMACD_Update", fields{
			slow: &emaCalculator{
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
			fast: &emaCalculator{
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
			signal: &emaCalculator{
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
			DIF: 0,
			DEM: 0,
			OSC: 0,
		}, args{
			price: 0,
			date:  time.Now(),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MACD{
				slow:   tt.fields.slow,
				fast:   tt.fields.fast,
				signal: tt.fields.signal,
				DIF:    tt.fields.DIF,
				DEM:    tt.fields.DEM,
				OSC:    tt.fields.OSC,
			}
			m.Update(tt.args.price, tt.args.date)
		})
	}
}

func TestTicker_NewMACD(t1 *testing.T) {
	type fields struct {
		Price float64
		Date  time.Time
	}
	type args struct {
		inFastPeriod   int32
		inSlowPeriod   int32
		inSignalPeriod int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MACD
	}{
		{"TestTicker_NewMACD", fields{
			Price: 0,
			Date:  time.Time{},
		}, args{
			inFastPeriod:   0,
			inSlowPeriod:   0,
			inSignalPeriod: 0,
		}, nil},
		{"TestTicker_NewMACD", fields{
			Price: 0,
			Date:  time.Time{},
		}, args{
			inFastPeriod:   20,
			inSlowPeriod:   10,
			inSignalPeriod: 0,
		}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticker{
				Price: tt.fields.Price,
				Date:  tt.fields.Date,
			}
			if got := t.NewMACD(tt.args.inFastPeriod, tt.args.inSlowPeriod, tt.args.inSignalPeriod); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewMACD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMACD_Update(b *testing.B) {
	ticker := &Ticker{}
	macd := ticker.NewMACD(12, 26, 9)
	for n := 0; n < b.N; n++ {
		macd.Update(float64(n), time.Now())
	}
}
