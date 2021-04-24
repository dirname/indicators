package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//Market market data
type Market []struct {
	Time  int64
	Close float64
	Open  float64
	Low   float64
	High  float64
	Vol   float64
}

//GetCandlestick get candlestick
func GetCandlestick(period, size, symbol string) *Market {
	url := "https://api.binance.com/api/v3/klines?symbol=" + symbol + "&interval=" + period + "&limit=" + size
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var i []interface{}
	var m Market
	json.Unmarshal(body, &i)
	for _, v := range i {
		i := v.([]interface{})
		close, _ := strconv.ParseFloat(i[4].(string), 64)
		open, _ := strconv.ParseFloat(i[1].(string), 64)
		high, _ := strconv.ParseFloat(i[2].(string), 64)
		low, _ := strconv.ParseFloat(i[3].(string), 64)
		vol, _ := strconv.ParseFloat(i[5].(string), 64)
		m = append(m, struct {
			Time  int64
			Close float64
			Open  float64
			Low   float64
			High  float64
			Vol   float64
		}{Time: int64(i[0].(float64)), Close: close, Open: open, High: high, Low: low, Vol: vol})
	}
	m = m[0 : len(m)-1] // remove now data
	return &m
}
