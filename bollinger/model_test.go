package bollinger

import (
	"testing"
	"time"
)

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
