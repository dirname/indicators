package ma

import (
	"math"
	"reflect"
	"testing"
	"time"
)

func TestMAMA_Sum(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *mamaCalculator
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
		want1  float64
	}{
		{"TestMAMA_Sum", fields{
			Value: 0,
			Calculator: &mamaCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				Deg:      180.0 / (4.0 * math.Atan(1)),
				Odd:      make([]float64, 3, 3),
				Even:     make([]float64, 3, 3),
				Q1Even:   make([]float64, 3, 3),
				Q1Odd:    make([]float64, 3, 3),
				JIEven:   make([]float64, 3, 3),
				JIOdd:    make([]float64, 3, 3),
				JQEven:   make([]float64, 3, 3),
				JQOdd:    make([]float64, 3, 3),
				WMAValue: make([]float64, 0, 5),
				ParamA:   0.0962,
				ParamB:   0.5769,
			},
		}, 0, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MAMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			got, got1 := d.Sum()
			if got != tt.want {
				t.Errorf("Sum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Sum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMAMA_Update(t *testing.T) {
	type fields struct {
		Value      float64
		Calculator *mamaCalculator
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
		{"TestMAMA_Update", fields{
			Value: 0,
			Calculator: &mamaCalculator{
				Ticker: &Ticker{
					Price: 0,
					Date:  time.Now(),
				},
				Deg:      180.0 / (4.0 * math.Atan(1)),
				Odd:      make([]float64, 3, 3),
				Even:     make([]float64, 3, 3),
				Q1Even:   make([]float64, 3, 3),
				Q1Odd:    make([]float64, 3, 3),
				JIEven:   make([]float64, 3, 3),
				JIOdd:    make([]float64, 3, 3),
				JQEven:   make([]float64, 3, 3),
				JQOdd:    make([]float64, 3, 3),
				WMAValue: make([]float64, 0, 5),
				ParamA:   0.0962,
				ParamB:   0.5769,
			},
		}, args{
			price: 0,
			date:  time.Now(),
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MAMA{
				Value:      tt.fields.Value,
				Calculator: tt.fields.Calculator,
			}
			d.Update(tt.args.price, tt.args.date)
		})
	}
}

func TestTicker_NewMAMA(t1 *testing.T) {
	type fields struct {
		Price float64
		Date  time.Time
	}
	type args struct {
		inFast float64
		inSlow float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MAMA
	}{
		{"TestTicker_NewMAMA", fields{
			Price: 0,
			Date:  time.Now(),
		}, args{
			inFast: 0,
			inSlow: 0,
		}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticker{
				Price: tt.fields.Price,
				Date:  tt.fields.Date,
			}
			if got := t.NewMAMA(tt.args.inFast, tt.args.inSlow); reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewMAMA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mamaCalculator_calcMAMA(t *testing.T) {
	type fields struct {
		Ticker          *Ticker
		Odd             []float64
		Even            []float64
		Q1Odd           []float64
		Q1Even          []float64
		JIOdd           []float64
		JIEven          []float64
		JQOdd           []float64
		JQEven          []float64
		WMAValue        []float64
		Deg             float64
		WMASub          float64
		WMASum          float64
		TrailingWMA     float64
		SmoothValue     float64
		Adjusted        float64
		Q1Value         float64
		Q2Value         float64
		I1Value         float64
		I2Value         float64
		Hilbert         float64
		ParamA          float64
		ParamB          float64
		detRender       float64
		PrevDetEven     float64
		PrevDetOdd      float64
		PrevInputEven   float64
		PrevQ1Odd       float64
		PrevQ1InputOdd  float64
		PrevInputOdd    float64
		PrevJIInput     float64
		PrevJQInput     float64
		PervQ1Even      float64
		EvenPrev3       float64
		EvenPrev2       float64
		OddPrev3        float64
		OddPrev2        float64
		PrevJIEven      float64
		PrevJQEven      float64
		PrevJIOdd       float64
		PrevJIInputOdd  float64
		PrevJQInputOdd  float64
		PrevJQOdd       float64
		PrevQ2          float64
		PrevI2          float64
		JI              float64
		JQ              float64
		Temp            float64
		Cache           float64
		PrevPhase       float64
		FastLimit       float64
		SlowLimit       float64
		MAMA            float64
		FAMA            float64
		Period          float64
		Re              float64
		Im              float64
		PrevQ1InputEven float64
		Count           int32
		HilbertIdx      int32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_mamaCalculator_calcMAMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
		}},
		{"Test_mamaCalculator_calcMAMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
			Count:    1,
		}},
		{"Test_mamaCalculator_calcMAMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
			Count:    2,
		}},
		{"Test_mamaCalculator_calcMAMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
			Count:    3,
		}},
		{"Test_mamaCalculator_calcMAMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
			Count:    20,
		}},
		{"Test_mamaCalculator_calcMAMA", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
			Count:    19,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mamaCalculator{
				Ticker:          tt.fields.Ticker,
				Odd:             tt.fields.Odd,
				Even:            tt.fields.Even,
				Q1Odd:           tt.fields.Q1Odd,
				Q1Even:          tt.fields.Q1Even,
				JIOdd:           tt.fields.JIOdd,
				JIEven:          tt.fields.JIEven,
				JQOdd:           tt.fields.JQOdd,
				JQEven:          tt.fields.JQEven,
				WMAValue:        tt.fields.WMAValue,
				Deg:             tt.fields.Deg,
				WMASub:          tt.fields.WMASub,
				WMASum:          tt.fields.WMASum,
				TrailingWMA:     tt.fields.TrailingWMA,
				SmoothValue:     tt.fields.SmoothValue,
				Adjusted:        tt.fields.Adjusted,
				Q1Value:         tt.fields.Q1Value,
				Q2Value:         tt.fields.Q2Value,
				I1Value:         tt.fields.I1Value,
				I2Value:         tt.fields.I2Value,
				Hilbert:         tt.fields.Hilbert,
				ParamA:          tt.fields.ParamA,
				ParamB:          tt.fields.ParamB,
				detRender:       tt.fields.detRender,
				PrevDetEven:     tt.fields.PrevDetEven,
				PrevDetOdd:      tt.fields.PrevDetOdd,
				PrevInputEven:   tt.fields.PrevInputEven,
				PrevQ1Odd:       tt.fields.PrevQ1Odd,
				PrevQ1InputOdd:  tt.fields.PrevQ1InputOdd,
				PrevInputOdd:    tt.fields.PrevInputOdd,
				PrevJIInput:     tt.fields.PrevJIInput,
				PrevJQInput:     tt.fields.PrevJQInput,
				PervQ1Even:      tt.fields.PervQ1Even,
				EvenPrev3:       tt.fields.EvenPrev3,
				EvenPrev2:       tt.fields.EvenPrev2,
				OddPrev3:        tt.fields.OddPrev3,
				OddPrev2:        tt.fields.OddPrev2,
				PrevJIEven:      tt.fields.PrevJIEven,
				PrevJQEven:      tt.fields.PrevJQEven,
				PrevJIOdd:       tt.fields.PrevJIOdd,
				PrevJIInputOdd:  tt.fields.PrevJIInputOdd,
				PrevJQInputOdd:  tt.fields.PrevJQInputOdd,
				PrevJQOdd:       tt.fields.PrevJQOdd,
				PrevQ2:          tt.fields.PrevQ2,
				PrevI2:          tt.fields.PrevI2,
				JI:              tt.fields.JI,
				JQ:              tt.fields.JQ,
				Temp:            tt.fields.Temp,
				Cache:           tt.fields.Cache,
				PrevPhase:       tt.fields.PrevPhase,
				FastLimit:       tt.fields.FastLimit,
				SlowLimit:       tt.fields.SlowLimit,
				MAMA:            tt.fields.MAMA,
				FAMA:            tt.fields.FAMA,
				Period:          tt.fields.Period,
				Re:              tt.fields.Re,
				Im:              tt.fields.Im,
				PrevQ1InputEven: tt.fields.PrevQ1InputEven,
				Count:           tt.fields.Count,
				HilbertIdx:      tt.fields.HilbertIdx,
			}
			s.calcMAMA()
		})
	}
}

func Test_mamaCalculator_isPowerOfTwo(t *testing.T) {
	type fields struct {
		Ticker          *Ticker
		Odd             []float64
		Even            []float64
		Q1Odd           []float64
		Q1Even          []float64
		JIOdd           []float64
		JIEven          []float64
		JQOdd           []float64
		JQEven          []float64
		WMAValue        []float64
		Deg             float64
		WMASub          float64
		WMASum          float64
		TrailingWMA     float64
		SmoothValue     float64
		Adjusted        float64
		Q1Value         float64
		Q2Value         float64
		I1Value         float64
		I2Value         float64
		Hilbert         float64
		ParamA          float64
		ParamB          float64
		detRender       float64
		PrevDetEven     float64
		PrevDetOdd      float64
		PrevInputEven   float64
		PrevQ1Odd       float64
		PrevQ1InputOdd  float64
		PrevInputOdd    float64
		PrevJIInput     float64
		PrevJQInput     float64
		PervQ1Even      float64
		EvenPrev3       float64
		EvenPrev2       float64
		OddPrev3        float64
		OddPrev2        float64
		PrevJIEven      float64
		PrevJQEven      float64
		PrevJIOdd       float64
		PrevJIInputOdd  float64
		PrevJQInputOdd  float64
		PrevJQOdd       float64
		PrevQ2          float64
		PrevI2          float64
		JI              float64
		JQ              float64
		Temp            float64
		Cache           float64
		PrevPhase       float64
		FastLimit       float64
		SlowLimit       float64
		MAMA            float64
		FAMA            float64
		Period          float64
		Re              float64
		Im              float64
		PrevQ1InputEven float64
		Count           int32
		HilbertIdx      int32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_mamaCalculator_isPowerOfTwo", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
		}},
		{"Test_mamaCalculator_isPowerOfTwo", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:        180.0 / (4.0 * math.Atan(1)),
			Odd:        make([]float64, 3, 3),
			Even:       make([]float64, 3, 3),
			Q1Even:     make([]float64, 3, 3),
			Q1Odd:      make([]float64, 3, 3),
			JIEven:     make([]float64, 3, 3),
			JIOdd:      make([]float64, 3, 3),
			JQEven:     make([]float64, 3, 3),
			JQOdd:      make([]float64, 3, 3),
			WMAValue:   make([]float64, 0, 5),
			ParamA:     0.0962,
			ParamB:     0.5769,
			HilbertIdx: 2,
			EvenPrev3:  3.0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mamaCalculator{
				Ticker:          tt.fields.Ticker,
				Odd:             tt.fields.Odd,
				Even:            tt.fields.Even,
				Q1Odd:           tt.fields.Q1Odd,
				Q1Even:          tt.fields.Q1Even,
				JIOdd:           tt.fields.JIOdd,
				JIEven:          tt.fields.JIEven,
				JQOdd:           tt.fields.JQOdd,
				JQEven:          tt.fields.JQEven,
				WMAValue:        tt.fields.WMAValue,
				Deg:             tt.fields.Deg,
				WMASub:          tt.fields.WMASub,
				WMASum:          tt.fields.WMASum,
				TrailingWMA:     tt.fields.TrailingWMA,
				SmoothValue:     tt.fields.SmoothValue,
				Adjusted:        tt.fields.Adjusted,
				Q1Value:         tt.fields.Q1Value,
				Q2Value:         tt.fields.Q2Value,
				I1Value:         tt.fields.I1Value,
				I2Value:         tt.fields.I2Value,
				Hilbert:         tt.fields.Hilbert,
				ParamA:          tt.fields.ParamA,
				ParamB:          tt.fields.ParamB,
				detRender:       tt.fields.detRender,
				PrevDetEven:     tt.fields.PrevDetEven,
				PrevDetOdd:      tt.fields.PrevDetOdd,
				PrevInputEven:   tt.fields.PrevInputEven,
				PrevQ1Odd:       tt.fields.PrevQ1Odd,
				PrevQ1InputOdd:  tt.fields.PrevQ1InputOdd,
				PrevInputOdd:    tt.fields.PrevInputOdd,
				PrevJIInput:     tt.fields.PrevJIInput,
				PrevJQInput:     tt.fields.PrevJQInput,
				PervQ1Even:      tt.fields.PervQ1Even,
				EvenPrev3:       tt.fields.EvenPrev3,
				EvenPrev2:       tt.fields.EvenPrev2,
				OddPrev3:        tt.fields.OddPrev3,
				OddPrev2:        tt.fields.OddPrev2,
				PrevJIEven:      tt.fields.PrevJIEven,
				PrevJQEven:      tt.fields.PrevJQEven,
				PrevJIOdd:       tt.fields.PrevJIOdd,
				PrevJIInputOdd:  tt.fields.PrevJIInputOdd,
				PrevJQInputOdd:  tt.fields.PrevJQInputOdd,
				PrevJQOdd:       tt.fields.PrevJQOdd,
				PrevQ2:          tt.fields.PrevQ2,
				PrevI2:          tt.fields.PrevI2,
				JI:              tt.fields.JI,
				JQ:              tt.fields.JQ,
				Temp:            tt.fields.Temp,
				Cache:           tt.fields.Cache,
				PrevPhase:       tt.fields.PrevPhase,
				FastLimit:       tt.fields.FastLimit,
				SlowLimit:       tt.fields.SlowLimit,
				MAMA:            tt.fields.MAMA,
				FAMA:            tt.fields.FAMA,
				Period:          tt.fields.Period,
				Re:              tt.fields.Re,
				Im:              tt.fields.Im,
				PrevQ1InputEven: tt.fields.PrevQ1InputEven,
				Count:           tt.fields.Count,
				HilbertIdx:      tt.fields.HilbertIdx,
			}
			s.isPowerOfTwo()
		})
	}
}

func Test_mamaCalculator_notPowerOfTwo(t *testing.T) {
	type fields struct {
		Ticker          *Ticker
		Odd             []float64
		Even            []float64
		Q1Odd           []float64
		Q1Even          []float64
		JIOdd           []float64
		JIEven          []float64
		JQOdd           []float64
		JQEven          []float64
		WMAValue        []float64
		Deg             float64
		WMASub          float64
		WMASum          float64
		TrailingWMA     float64
		SmoothValue     float64
		Adjusted        float64
		Q1Value         float64
		Q2Value         float64
		I1Value         float64
		I2Value         float64
		Hilbert         float64
		ParamA          float64
		ParamB          float64
		detRender       float64
		PrevDetEven     float64
		PrevDetOdd      float64
		PrevInputEven   float64
		PrevQ1Odd       float64
		PrevQ1InputOdd  float64
		PrevInputOdd    float64
		PrevJIInput     float64
		PrevJQInput     float64
		PervQ1Even      float64
		EvenPrev3       float64
		EvenPrev2       float64
		OddPrev3        float64
		OddPrev2        float64
		PrevJIEven      float64
		PrevJQEven      float64
		PrevJIOdd       float64
		PrevJIInputOdd  float64
		PrevJQInputOdd  float64
		PrevJQOdd       float64
		PrevQ2          float64
		PrevI2          float64
		JI              float64
		JQ              float64
		Temp            float64
		Cache           float64
		PrevPhase       float64
		FastLimit       float64
		SlowLimit       float64
		MAMA            float64
		FAMA            float64
		Period          float64
		Re              float64
		Im              float64
		PrevQ1InputEven float64
		Count           int32
		HilbertIdx      int32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_mamaCalculator_notPowerOfTwo", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
		}},
		{"Test_mamaCalculator_notPowerOfTwo", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
			OddPrev3: 4.0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mamaCalculator{
				Ticker:          tt.fields.Ticker,
				Odd:             tt.fields.Odd,
				Even:            tt.fields.Even,
				Q1Odd:           tt.fields.Q1Odd,
				Q1Even:          tt.fields.Q1Even,
				JIOdd:           tt.fields.JIOdd,
				JIEven:          tt.fields.JIEven,
				JQOdd:           tt.fields.JQOdd,
				JQEven:          tt.fields.JQEven,
				WMAValue:        tt.fields.WMAValue,
				Deg:             tt.fields.Deg,
				WMASub:          tt.fields.WMASub,
				WMASum:          tt.fields.WMASum,
				TrailingWMA:     tt.fields.TrailingWMA,
				SmoothValue:     tt.fields.SmoothValue,
				Adjusted:        tt.fields.Adjusted,
				Q1Value:         tt.fields.Q1Value,
				Q2Value:         tt.fields.Q2Value,
				I1Value:         tt.fields.I1Value,
				I2Value:         tt.fields.I2Value,
				Hilbert:         tt.fields.Hilbert,
				ParamA:          tt.fields.ParamA,
				ParamB:          tt.fields.ParamB,
				detRender:       tt.fields.detRender,
				PrevDetEven:     tt.fields.PrevDetEven,
				PrevDetOdd:      tt.fields.PrevDetOdd,
				PrevInputEven:   tt.fields.PrevInputEven,
				PrevQ1Odd:       tt.fields.PrevQ1Odd,
				PrevQ1InputOdd:  tt.fields.PrevQ1InputOdd,
				PrevInputOdd:    tt.fields.PrevInputOdd,
				PrevJIInput:     tt.fields.PrevJIInput,
				PrevJQInput:     tt.fields.PrevJQInput,
				PervQ1Even:      tt.fields.PervQ1Even,
				EvenPrev3:       tt.fields.EvenPrev3,
				EvenPrev2:       tt.fields.EvenPrev2,
				OddPrev3:        tt.fields.OddPrev3,
				OddPrev2:        tt.fields.OddPrev2,
				PrevJIEven:      tt.fields.PrevJIEven,
				PrevJQEven:      tt.fields.PrevJQEven,
				PrevJIOdd:       tt.fields.PrevJIOdd,
				PrevJIInputOdd:  tt.fields.PrevJIInputOdd,
				PrevJQInputOdd:  tt.fields.PrevJQInputOdd,
				PrevJQOdd:       tt.fields.PrevJQOdd,
				PrevQ2:          tt.fields.PrevQ2,
				PrevI2:          tt.fields.PrevI2,
				JI:              tt.fields.JI,
				JQ:              tt.fields.JQ,
				Temp:            tt.fields.Temp,
				Cache:           tt.fields.Cache,
				PrevPhase:       tt.fields.PrevPhase,
				FastLimit:       tt.fields.FastLimit,
				SlowLimit:       tt.fields.SlowLimit,
				MAMA:            tt.fields.MAMA,
				FAMA:            tt.fields.FAMA,
				Period:          tt.fields.Period,
				Re:              tt.fields.Re,
				Im:              tt.fields.Im,
				PrevQ1InputEven: tt.fields.PrevQ1InputEven,
				Count:           tt.fields.Count,
				HilbertIdx:      tt.fields.HilbertIdx,
			}
			s.notPowerOfTwo()
		})
	}
}

func Test_mamaCalculator_sum(t *testing.T) {
	type fields struct {
		Ticker          *Ticker
		Odd             []float64
		Even            []float64
		Q1Odd           []float64
		Q1Even          []float64
		JIOdd           []float64
		JIEven          []float64
		JQOdd           []float64
		JQEven          []float64
		WMAValue        []float64
		Deg             float64
		WMASub          float64
		WMASum          float64
		TrailingWMA     float64
		SmoothValue     float64
		Adjusted        float64
		Q1Value         float64
		Q2Value         float64
		I1Value         float64
		I2Value         float64
		Hilbert         float64
		ParamA          float64
		ParamB          float64
		detRender       float64
		PrevDetEven     float64
		PrevDetOdd      float64
		PrevInputEven   float64
		PrevQ1Odd       float64
		PrevQ1InputOdd  float64
		PrevInputOdd    float64
		PrevJIInput     float64
		PrevJQInput     float64
		PervQ1Even      float64
		EvenPrev3       float64
		EvenPrev2       float64
		OddPrev3        float64
		OddPrev2        float64
		PrevJIEven      float64
		PrevJQEven      float64
		PrevJIOdd       float64
		PrevJIInputOdd  float64
		PrevJQInputOdd  float64
		PrevJQOdd       float64
		PrevQ2          float64
		PrevI2          float64
		JI              float64
		JQ              float64
		Temp            float64
		Cache           float64
		PrevPhase       float64
		FastLimit       float64
		SlowLimit       float64
		MAMA            float64
		FAMA            float64
		Period          float64
		Re              float64
		Im              float64
		PrevQ1InputEven float64
		Count           int32
		HilbertIdx      int32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_mamaCalculator_sum", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mamaCalculator{
				Ticker:          tt.fields.Ticker,
				Odd:             tt.fields.Odd,
				Even:            tt.fields.Even,
				Q1Odd:           tt.fields.Q1Odd,
				Q1Even:          tt.fields.Q1Even,
				JIOdd:           tt.fields.JIOdd,
				JIEven:          tt.fields.JIEven,
				JQOdd:           tt.fields.JQOdd,
				JQEven:          tt.fields.JQEven,
				WMAValue:        tt.fields.WMAValue,
				Deg:             tt.fields.Deg,
				WMASub:          tt.fields.WMASub,
				WMASum:          tt.fields.WMASum,
				TrailingWMA:     tt.fields.TrailingWMA,
				SmoothValue:     tt.fields.SmoothValue,
				Adjusted:        tt.fields.Adjusted,
				Q1Value:         tt.fields.Q1Value,
				Q2Value:         tt.fields.Q2Value,
				I1Value:         tt.fields.I1Value,
				I2Value:         tt.fields.I2Value,
				Hilbert:         tt.fields.Hilbert,
				ParamA:          tt.fields.ParamA,
				ParamB:          tt.fields.ParamB,
				detRender:       tt.fields.detRender,
				PrevDetEven:     tt.fields.PrevDetEven,
				PrevDetOdd:      tt.fields.PrevDetOdd,
				PrevInputEven:   tt.fields.PrevInputEven,
				PrevQ1Odd:       tt.fields.PrevQ1Odd,
				PrevQ1InputOdd:  tt.fields.PrevQ1InputOdd,
				PrevInputOdd:    tt.fields.PrevInputOdd,
				PrevJIInput:     tt.fields.PrevJIInput,
				PrevJQInput:     tt.fields.PrevJQInput,
				PervQ1Even:      tt.fields.PervQ1Even,
				EvenPrev3:       tt.fields.EvenPrev3,
				EvenPrev2:       tt.fields.EvenPrev2,
				OddPrev3:        tt.fields.OddPrev3,
				OddPrev2:        tt.fields.OddPrev2,
				PrevJIEven:      tt.fields.PrevJIEven,
				PrevJQEven:      tt.fields.PrevJQEven,
				PrevJIOdd:       tt.fields.PrevJIOdd,
				PrevJIInputOdd:  tt.fields.PrevJIInputOdd,
				PrevJQInputOdd:  tt.fields.PrevJQInputOdd,
				PrevJQOdd:       tt.fields.PrevJQOdd,
				PrevQ2:          tt.fields.PrevQ2,
				PrevI2:          tt.fields.PrevI2,
				JI:              tt.fields.JI,
				JQ:              tt.fields.JQ,
				Temp:            tt.fields.Temp,
				Cache:           tt.fields.Cache,
				PrevPhase:       tt.fields.PrevPhase,
				FastLimit:       tt.fields.FastLimit,
				SlowLimit:       tt.fields.SlowLimit,
				MAMA:            tt.fields.MAMA,
				FAMA:            tt.fields.FAMA,
				Period:          tt.fields.Period,
				Re:              tt.fields.Re,
				Im:              tt.fields.Im,
				PrevQ1InputEven: tt.fields.PrevQ1InputEven,
				Count:           tt.fields.Count,
				HilbertIdx:      tt.fields.HilbertIdx,
			}
			s.sum()
		})
	}
}

func Test_mamaCalculator_sumOne(t *testing.T) {
	type fields struct {
		Ticker          *Ticker
		Odd             []float64
		Even            []float64
		Q1Odd           []float64
		Q1Even          []float64
		JIOdd           []float64
		JIEven          []float64
		JQOdd           []float64
		JQEven          []float64
		WMAValue        []float64
		Deg             float64
		WMASub          float64
		WMASum          float64
		TrailingWMA     float64
		SmoothValue     float64
		Adjusted        float64
		Q1Value         float64
		Q2Value         float64
		I1Value         float64
		I2Value         float64
		Hilbert         float64
		ParamA          float64
		ParamB          float64
		detRender       float64
		PrevDetEven     float64
		PrevDetOdd      float64
		PrevInputEven   float64
		PrevQ1Odd       float64
		PrevQ1InputOdd  float64
		PrevInputOdd    float64
		PrevJIInput     float64
		PrevJQInput     float64
		PervQ1Even      float64
		EvenPrev3       float64
		EvenPrev2       float64
		OddPrev3        float64
		OddPrev2        float64
		PrevJIEven      float64
		PrevJQEven      float64
		PrevJIOdd       float64
		PrevJIInputOdd  float64
		PrevJQInputOdd  float64
		PrevJQOdd       float64
		PrevQ2          float64
		PrevI2          float64
		JI              float64
		JQ              float64
		Temp            float64
		Cache           float64
		PrevPhase       float64
		FastLimit       float64
		SlowLimit       float64
		MAMA            float64
		FAMA            float64
		Period          float64
		Re              float64
		Im              float64
		PrevQ1InputEven float64
		Count           int32
		HilbertIdx      int32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_mamaCalculator_sumOne", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
		}},
		{"Test_mamaCalculator_sumOne", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:       180.0 / (4.0 * math.Atan(1)),
			Odd:       make([]float64, 3, 3),
			Even:      make([]float64, 3, 3),
			Q1Even:    make([]float64, 3, 3),
			Q1Odd:     make([]float64, 3, 3),
			JIEven:    make([]float64, 3, 3),
			JIOdd:     make([]float64, 3, 3),
			JQEven:    make([]float64, 3, 3),
			JQOdd:     make([]float64, 3, 3),
			WMAValue:  make([]float64, 0, 5),
			ParamA:    0.0962,
			ParamB:    0.5769,
			PrevPhase: 3,
			Cache:     1,
			Temp:      2.0,
			FastLimit: 200,
			SlowLimit: 500,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mamaCalculator{
				Ticker:          tt.fields.Ticker,
				Odd:             tt.fields.Odd,
				Even:            tt.fields.Even,
				Q1Odd:           tt.fields.Q1Odd,
				Q1Even:          tt.fields.Q1Even,
				JIOdd:           tt.fields.JIOdd,
				JIEven:          tt.fields.JIEven,
				JQOdd:           tt.fields.JQOdd,
				JQEven:          tt.fields.JQEven,
				WMAValue:        tt.fields.WMAValue,
				Deg:             tt.fields.Deg,
				WMASub:          tt.fields.WMASub,
				WMASum:          tt.fields.WMASum,
				TrailingWMA:     tt.fields.TrailingWMA,
				SmoothValue:     tt.fields.SmoothValue,
				Adjusted:        tt.fields.Adjusted,
				Q1Value:         tt.fields.Q1Value,
				Q2Value:         tt.fields.Q2Value,
				I1Value:         tt.fields.I1Value,
				I2Value:         tt.fields.I2Value,
				Hilbert:         tt.fields.Hilbert,
				ParamA:          tt.fields.ParamA,
				ParamB:          tt.fields.ParamB,
				detRender:       tt.fields.detRender,
				PrevDetEven:     tt.fields.PrevDetEven,
				PrevDetOdd:      tt.fields.PrevDetOdd,
				PrevInputEven:   tt.fields.PrevInputEven,
				PrevQ1Odd:       tt.fields.PrevQ1Odd,
				PrevQ1InputOdd:  tt.fields.PrevQ1InputOdd,
				PrevInputOdd:    tt.fields.PrevInputOdd,
				PrevJIInput:     tt.fields.PrevJIInput,
				PrevJQInput:     tt.fields.PrevJQInput,
				PervQ1Even:      tt.fields.PervQ1Even,
				EvenPrev3:       tt.fields.EvenPrev3,
				EvenPrev2:       tt.fields.EvenPrev2,
				OddPrev3:        tt.fields.OddPrev3,
				OddPrev2:        tt.fields.OddPrev2,
				PrevJIEven:      tt.fields.PrevJIEven,
				PrevJQEven:      tt.fields.PrevJQEven,
				PrevJIOdd:       tt.fields.PrevJIOdd,
				PrevJIInputOdd:  tt.fields.PrevJIInputOdd,
				PrevJQInputOdd:  tt.fields.PrevJQInputOdd,
				PrevJQOdd:       tt.fields.PrevJQOdd,
				PrevQ2:          tt.fields.PrevQ2,
				PrevI2:          tt.fields.PrevI2,
				JI:              tt.fields.JI,
				JQ:              tt.fields.JQ,
				Temp:            tt.fields.Temp,
				Cache:           tt.fields.Cache,
				PrevPhase:       tt.fields.PrevPhase,
				FastLimit:       tt.fields.FastLimit,
				SlowLimit:       tt.fields.SlowLimit,
				MAMA:            tt.fields.MAMA,
				FAMA:            tt.fields.FAMA,
				Period:          tt.fields.Period,
				Re:              tt.fields.Re,
				Im:              tt.fields.Im,
				PrevQ1InputEven: tt.fields.PrevQ1InputEven,
				Count:           tt.fields.Count,
				HilbertIdx:      tt.fields.HilbertIdx,
			}
			s.sumOne()
		})
	}
}

func Test_mamaCalculator_sumThree(t *testing.T) {
	type fields struct {
		Ticker          *Ticker
		Odd             []float64
		Even            []float64
		Q1Odd           []float64
		Q1Even          []float64
		JIOdd           []float64
		JIEven          []float64
		JQOdd           []float64
		JQEven          []float64
		WMAValue        []float64
		Deg             float64
		WMASub          float64
		WMASum          float64
		TrailingWMA     float64
		SmoothValue     float64
		Adjusted        float64
		Q1Value         float64
		Q2Value         float64
		I1Value         float64
		I2Value         float64
		Hilbert         float64
		ParamA          float64
		ParamB          float64
		detRender       float64
		PrevDetEven     float64
		PrevDetOdd      float64
		PrevInputEven   float64
		PrevQ1Odd       float64
		PrevQ1InputOdd  float64
		PrevInputOdd    float64
		PrevJIInput     float64
		PrevJQInput     float64
		PervQ1Even      float64
		EvenPrev3       float64
		EvenPrev2       float64
		OddPrev3        float64
		OddPrev2        float64
		PrevJIEven      float64
		PrevJQEven      float64
		PrevJIOdd       float64
		PrevJIInputOdd  float64
		PrevJQInputOdd  float64
		PrevJQOdd       float64
		PrevQ2          float64
		PrevI2          float64
		JI              float64
		JQ              float64
		Temp            float64
		Cache           float64
		PrevPhase       float64
		FastLimit       float64
		SlowLimit       float64
		MAMA            float64
		FAMA            float64
		Period          float64
		Re              float64
		Im              float64
		PrevQ1InputEven float64
		Count           int32
		HilbertIdx      int32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_mamaCalculator_sumThree", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
		}},
		{"Test_mamaCalculator_sumThree", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
			Period:   60,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mamaCalculator{
				Ticker:          tt.fields.Ticker,
				Odd:             tt.fields.Odd,
				Even:            tt.fields.Even,
				Q1Odd:           tt.fields.Q1Odd,
				Q1Even:          tt.fields.Q1Even,
				JIOdd:           tt.fields.JIOdd,
				JIEven:          tt.fields.JIEven,
				JQOdd:           tt.fields.JQOdd,
				JQEven:          tt.fields.JQEven,
				WMAValue:        tt.fields.WMAValue,
				Deg:             tt.fields.Deg,
				WMASub:          tt.fields.WMASub,
				WMASum:          tt.fields.WMASum,
				TrailingWMA:     tt.fields.TrailingWMA,
				SmoothValue:     tt.fields.SmoothValue,
				Adjusted:        tt.fields.Adjusted,
				Q1Value:         tt.fields.Q1Value,
				Q2Value:         tt.fields.Q2Value,
				I1Value:         tt.fields.I1Value,
				I2Value:         tt.fields.I2Value,
				Hilbert:         tt.fields.Hilbert,
				ParamA:          tt.fields.ParamA,
				ParamB:          tt.fields.ParamB,
				detRender:       tt.fields.detRender,
				PrevDetEven:     tt.fields.PrevDetEven,
				PrevDetOdd:      tt.fields.PrevDetOdd,
				PrevInputEven:   tt.fields.PrevInputEven,
				PrevQ1Odd:       tt.fields.PrevQ1Odd,
				PrevQ1InputOdd:  tt.fields.PrevQ1InputOdd,
				PrevInputOdd:    tt.fields.PrevInputOdd,
				PrevJIInput:     tt.fields.PrevJIInput,
				PrevJQInput:     tt.fields.PrevJQInput,
				PervQ1Even:      tt.fields.PervQ1Even,
				EvenPrev3:       tt.fields.EvenPrev3,
				EvenPrev2:       tt.fields.EvenPrev2,
				OddPrev3:        tt.fields.OddPrev3,
				OddPrev2:        tt.fields.OddPrev2,
				PrevJIEven:      tt.fields.PrevJIEven,
				PrevJQEven:      tt.fields.PrevJQEven,
				PrevJIOdd:       tt.fields.PrevJIOdd,
				PrevJIInputOdd:  tt.fields.PrevJIInputOdd,
				PrevJQInputOdd:  tt.fields.PrevJQInputOdd,
				PrevJQOdd:       tt.fields.PrevJQOdd,
				PrevQ2:          tt.fields.PrevQ2,
				PrevI2:          tt.fields.PrevI2,
				JI:              tt.fields.JI,
				JQ:              tt.fields.JQ,
				Temp:            tt.fields.Temp,
				Cache:           tt.fields.Cache,
				PrevPhase:       tt.fields.PrevPhase,
				FastLimit:       tt.fields.FastLimit,
				SlowLimit:       tt.fields.SlowLimit,
				MAMA:            tt.fields.MAMA,
				FAMA:            tt.fields.FAMA,
				Period:          tt.fields.Period,
				Re:              tt.fields.Re,
				Im:              tt.fields.Im,
				PrevQ1InputEven: tt.fields.PrevQ1InputEven,
				Count:           tt.fields.Count,
				HilbertIdx:      tt.fields.HilbertIdx,
			}
			s.sumThree()
		})
	}
}

func Test_mamaCalculator_sumTwo(t *testing.T) {
	type fields struct {
		Ticker          *Ticker
		Odd             []float64
		Even            []float64
		Q1Odd           []float64
		Q1Even          []float64
		JIOdd           []float64
		JIEven          []float64
		JQOdd           []float64
		JQEven          []float64
		WMAValue        []float64
		Deg             float64
		WMASub          float64
		WMASum          float64
		TrailingWMA     float64
		SmoothValue     float64
		Adjusted        float64
		Q1Value         float64
		Q2Value         float64
		I1Value         float64
		I2Value         float64
		Hilbert         float64
		ParamA          float64
		ParamB          float64
		detRender       float64
		PrevDetEven     float64
		PrevDetOdd      float64
		PrevInputEven   float64
		PrevQ1Odd       float64
		PrevQ1InputOdd  float64
		PrevInputOdd    float64
		PrevJIInput     float64
		PrevJQInput     float64
		PervQ1Even      float64
		EvenPrev3       float64
		EvenPrev2       float64
		OddPrev3        float64
		OddPrev2        float64
		PrevJIEven      float64
		PrevJQEven      float64
		PrevJIOdd       float64
		PrevJIInputOdd  float64
		PrevJQInputOdd  float64
		PrevJQOdd       float64
		PrevQ2          float64
		PrevI2          float64
		JI              float64
		JQ              float64
		Temp            float64
		Cache           float64
		PrevPhase       float64
		FastLimit       float64
		SlowLimit       float64
		MAMA            float64
		FAMA            float64
		Period          float64
		Re              float64
		Im              float64
		PrevQ1InputEven float64
		Count           int32
		HilbertIdx      int32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Test_mamaCalculator_sumTwo", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
		}},
		{"Test_mamaCalculator_sumTwo", fields{
			Ticker: &Ticker{
				Price: 0,
				Date:  time.Now(),
			},
			Deg:      180.0 / (4.0 * math.Atan(1)),
			Odd:      make([]float64, 3, 3),
			Even:     make([]float64, 3, 3),
			Q1Even:   make([]float64, 3, 3),
			Q1Odd:    make([]float64, 3, 3),
			JIEven:   make([]float64, 3, 3),
			JIOdd:    make([]float64, 3, 3),
			JQEven:   make([]float64, 3, 3),
			JQOdd:    make([]float64, 3, 3),
			WMAValue: make([]float64, 0, 5),
			ParamA:   0.0962,
			ParamB:   0.5769,
			I2Value:  1,
			PrevI2:   1,
			PrevQ2:   1,
			Q2Value:  1,
			Re:       1,
			Im:       1,
			Cache:    1,
			Temp:     2,
			Period:   3,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mamaCalculator{
				Ticker:          tt.fields.Ticker,
				Odd:             tt.fields.Odd,
				Even:            tt.fields.Even,
				Q1Odd:           tt.fields.Q1Odd,
				Q1Even:          tt.fields.Q1Even,
				JIOdd:           tt.fields.JIOdd,
				JIEven:          tt.fields.JIEven,
				JQOdd:           tt.fields.JQOdd,
				JQEven:          tt.fields.JQEven,
				WMAValue:        tt.fields.WMAValue,
				Deg:             tt.fields.Deg,
				WMASub:          tt.fields.WMASub,
				WMASum:          tt.fields.WMASum,
				TrailingWMA:     tt.fields.TrailingWMA,
				SmoothValue:     tt.fields.SmoothValue,
				Adjusted:        tt.fields.Adjusted,
				Q1Value:         tt.fields.Q1Value,
				Q2Value:         tt.fields.Q2Value,
				I1Value:         tt.fields.I1Value,
				I2Value:         tt.fields.I2Value,
				Hilbert:         tt.fields.Hilbert,
				ParamA:          tt.fields.ParamA,
				ParamB:          tt.fields.ParamB,
				detRender:       tt.fields.detRender,
				PrevDetEven:     tt.fields.PrevDetEven,
				PrevDetOdd:      tt.fields.PrevDetOdd,
				PrevInputEven:   tt.fields.PrevInputEven,
				PrevQ1Odd:       tt.fields.PrevQ1Odd,
				PrevQ1InputOdd:  tt.fields.PrevQ1InputOdd,
				PrevInputOdd:    tt.fields.PrevInputOdd,
				PrevJIInput:     tt.fields.PrevJIInput,
				PrevJQInput:     tt.fields.PrevJQInput,
				PervQ1Even:      tt.fields.PervQ1Even,
				EvenPrev3:       tt.fields.EvenPrev3,
				EvenPrev2:       tt.fields.EvenPrev2,
				OddPrev3:        tt.fields.OddPrev3,
				OddPrev2:        tt.fields.OddPrev2,
				PrevJIEven:      tt.fields.PrevJIEven,
				PrevJQEven:      tt.fields.PrevJQEven,
				PrevJIOdd:       tt.fields.PrevJIOdd,
				PrevJIInputOdd:  tt.fields.PrevJIInputOdd,
				PrevJQInputOdd:  tt.fields.PrevJQInputOdd,
				PrevJQOdd:       tt.fields.PrevJQOdd,
				PrevQ2:          tt.fields.PrevQ2,
				PrevI2:          tt.fields.PrevI2,
				JI:              tt.fields.JI,
				JQ:              tt.fields.JQ,
				Temp:            tt.fields.Temp,
				Cache:           tt.fields.Cache,
				PrevPhase:       tt.fields.PrevPhase,
				FastLimit:       tt.fields.FastLimit,
				SlowLimit:       tt.fields.SlowLimit,
				MAMA:            tt.fields.MAMA,
				FAMA:            tt.fields.FAMA,
				Period:          tt.fields.Period,
				Re:              tt.fields.Re,
				Im:              tt.fields.Im,
				PrevQ1InputEven: tt.fields.PrevQ1InputEven,
				Count:           tt.fields.Count,
				HilbertIdx:      tt.fields.HilbertIdx,
			}
			s.sumThree()
		})
	}
}
