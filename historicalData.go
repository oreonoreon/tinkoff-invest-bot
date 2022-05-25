package main

import (
	"log"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/sdk"
)

var hc []*t.HistoricCandle

func hiscandle(n time.Duration) []*t.HistoricCandle {
	s := NewTimetoStamp()
	historicalCandles, err := sdk.NewMarketDataService().GetCandles("BBG006L8G4H1", s.ConvertTime(n).from, s.ConvertTime(n).to, t.CandleInterval_CANDLE_INTERVAL_15_MIN)
	if err != nil {
		log.Println(err)
	}
	/*
		for _, value := range historicalCandles {
			fmt.Println(value.Time.AsTime())
		}
	*/

	hc = append(hc, historicalCandles...)
	if n == 0 {
		return hc
	} else {
		return hiscandle(n - 1)
	}
}
