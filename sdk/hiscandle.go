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

func (sk *Services) Hiscandle(n time.Duration) []*t.HistoricCandle {
	s := NewTimetoStamp()
	historicalCandles, err := sk.MarketDataService.GetCandles("BBG006L8G4H1", s.ConvertTime(n).from, s.ConvertTime(n).to, t.CandleInterval_CANDLE_INTERVAL_5_MIN)
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
		return sk.Hiscandle(n - 1)
	}
}
