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
