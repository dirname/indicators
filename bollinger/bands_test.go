package bollinger

import (
	"github.com/dirname/indicators/common"
	"github.com/dirname/indicators/ma"
	"reflect"
	"testing"
	"time"
)

func TestBands_Sum(t *testing.T) {
	type fields struct {
		Calculator *bollingerBandsCalculator
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
		{"TestBands_Sum", fields{Calculator: &bollingerBandsCalculator{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			MAObject: new(ma.Ticker).NewSMA(9),
			StdDev:   new(common.Ticker).NewStdDev(9, 1.0),
		}}, args{"TEST"}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bands{
				Calculator: tt.fields.Calculator,
			}
			if got := b.Sum(tt.args.flag); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBands_Update(t *testing.T) {
	type fields struct {
		Calculator *bollingerBandsCalculator
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
		{"TestBands_Update", fields{Calculator: &bollingerBandsCalculator{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			MAObject: new(ma.Ticker).NewSMA(9),
			StdDev:   new(common.Ticker).NewStdDev(9, 1.0),
		}}, args{
			price: 0,
			date:  time.Now(),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bands{
				Calculator: tt.fields.Calculator,
			}
			b.Update(tt.args.price, tt.args.date)
		})
	}
}

func TestTicker_NewBands(t1 *testing.T) {
	type fields struct {
		Price float64
		Date  time.Time
	}
	type args struct {
		inTimePeriod int32
		Upper        float64
		Lower        float64
		obj          ma.Object
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Bands
	}{
		{"TestTicker_NewBands", fields{
			Price: 0,
			Date:  time.Now(),
		}, args{
			inTimePeriod: 9,
			Upper:        0,
			Lower:        0,
			obj:          new(ma.Ticker).NewSMA(9),
		}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticker{
				Price: tt.fields.Price,
				Date:  tt.fields.Date,
			}
			if got := t.NewBands(tt.args.inTimePeriod, tt.args.Upper, tt.args.Lower, tt.args.obj); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewBands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bollingerBandsCalculator_calcBands(t *testing.T) {
	type fields struct {
		Ticker      *Ticker
		Upper       float64
		Lower       float64
		Period      int32
		UpperBands  float64
		LowerBands  float64
		MiddleBands float64
		MAObject    ma.Object
		StdDev      *common.StdDev
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_bollingerBandsCalculator_calcBands", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Upper:       0,
			Lower:       0,
			Period:      0,
			UpperBands:  0,
			LowerBands:  0,
			MiddleBands: 0,
			MAObject:    new(ma.Ticker).NewSMA(9),
			StdDev:      new(common.Ticker).NewStdDev(9, 1.0),
		}},
		{"Test_bollingerBandsCalculator_calcBands", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Upper:       1.0,
			Lower:       1.0,
			Period:      0,
			UpperBands:  0,
			LowerBands:  0,
			MiddleBands: 0,
			MAObject:    new(ma.Ticker).NewSMA(9),
			StdDev:      new(common.Ticker).NewStdDev(9, 1.0),
		}},
		{"Test_bollingerBandsCalculator_calcBands", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Upper:       1.0,
			Lower:       2.0,
			Period:      0,
			UpperBands:  0,
			LowerBands:  0,
			MiddleBands: 0,
			MAObject:    new(ma.Ticker).NewSMA(9),
			StdDev:      new(common.Ticker).NewStdDev(9, 1.0),
		}},
		{"Test_bollingerBandsCalculator_calcBands", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Upper:       2.0,
			Lower:       1.0,
			Period:      0,
			UpperBands:  0,
			LowerBands:  0,
			MiddleBands: 0,
			MAObject:    new(ma.Ticker).NewSMA(9),
			StdDev:      new(common.Ticker).NewStdDev(9, 1.0),
		}},
		{"Test_bollingerBandsCalculator_calcBands", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Upper:       4.0,
			Lower:       3.0,
			Period:      0,
			UpperBands:  0,
			LowerBands:  0,
			MiddleBands: 0,
			MAObject:    new(ma.Ticker).NewSMA(9),
			StdDev:      new(common.Ticker).NewStdDev(9, 1.0),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &bollingerBandsCalculator{
				Ticker:      tt.fields.Ticker,
				Upper:       tt.fields.Upper,
				Lower:       tt.fields.Lower,
				Period:      tt.fields.Period,
				UpperBands:  tt.fields.UpperBands,
				LowerBands:  tt.fields.LowerBands,
				MiddleBands: tt.fields.MiddleBands,
				MAObject:    tt.fields.MAObject,
				StdDev:      tt.fields.StdDev,
			}
			c.calcBands()
		})
	}
}

func BenchmarkBands_Update(b *testing.B) {
	ticker := &Ticker{}
	maTicker := &ma.Ticker{}
	bands := ticker.NewBands(20, 2, 2, maTicker.NewSMA(20))
	for n := 0; n < b.N; n++ {
		bands.Update(float64(n), time.Now())
	}
}
