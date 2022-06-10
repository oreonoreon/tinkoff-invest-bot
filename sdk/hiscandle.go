package sdk

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
)

type TimeToStamp struct {
	from *timestamppb.Timestamp
	to   *timestamppb.Timestamp
}

func (s TimeToStamp) ConvertTime(days time.Duration) *TimeToStamp {
	x := -24 * days
	oneTrainingPeriod := time.Now().Add(x * time.Hour)
	s.from = timestamppb.New(oneTrainingPeriod.Add(-24 * time.Hour))
	s.to = timestamppb.New(oneTrainingPeriod)
	return &s
}
func NewTimetoStamp() *TimeToStamp {
	return new(TimeToStamp)
}

var hc []*t.HistoricCandle
var lt []*t.Trade

func (sk *Services) Hiscandle(numberOfperiud time.Duration, figi string, candleInterval t.CandleInterval) []*t.HistoricCandle {
	s := NewTimetoStamp()
	historicalCandles, err := sk.MarketDataService.GetCandles(figi, s.ConvertTime(numberOfperiud).from, s.ConvertTime(numberOfperiud).to, candleInterval)
	if err != nil {
		log.Println(err)
	}
	hc = append(hc, historicalCandles...)
	if numberOfperiud == 0 {
		return hc
	} else {
		return sk.Hiscandle(numberOfperiud-1, figi, candleInterval)
	}
}

func (sk *Services) HisLastTrade(numberOfperiud time.Duration, figi string) []*t.Trade {
	s := NewTimetoStamp()
	lastTrades, err := sk.MarketDataService.GetLastTrades(figi, s.ConvertTime(numberOfperiud).from, s.ConvertTime(numberOfperiud).to)
	if err != nil {
		log.Println(err)
	}
	lt = append(lt, lastTrades...)
	if numberOfperiud == 0 {
		return lt
	} else {
		return sk.HisLastTrade(numberOfperiud-1, figi)
	}
	return nil
}
