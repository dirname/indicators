package main

import (
	"fmt"
	"github.com/dirname/indicators/ma"
	"time"
)

func main() {
	macd := new(ma.Ticker).NewMACD(12, 26, 9)
	sma := new(ma.Ticker).NewSMA(9)
	ema := new(ma.Ticker).NewEMA(9)
	wma := new(ma.Ticker).NewWMA(9)
	dema := new(ma.Ticker).NewDEMA(9)
	data := GetCandlestick("1d", "20", "BTCUSDT")
	for _, v := range * data {
		date := time.Unix(0, v.Time*int64(time.Millisecond)).Format("2006-01-02 15:04:05")
		macd.Update(v.Close, time.Unix(0, v.Time*int64(time.Millisecond)))
		ema.Update(v.Close, time.Unix(0, v.Time*int64(time.Millisecond)))
		wma.Update(v.Close, time.Unix(0, v.Time*int64(time.Millisecond)))
		dema.Update(v.Close, time.Unix(0, v.Time*int64(time.Millisecond)))
		sma.Update(v.Close, time.Unix(0, v.Time*int64(time.Millisecond)))
		//fmt.Printf("%s MACD: %v EMA: %v WMA: %v DEMA: %v MA(9): %v\n", date, dema.Sum(), ema.Sum(), wma.Sum(), dema.Sum(), sma.Sum())
		fmt.Printf("%s %v\n", date, sma.Sum())
	}
	//fmt.Printf("%s %v", ticker.Date.Format("2006-01-02 15:04:05"), macd.OSC)
}
