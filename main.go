package main

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/sdk"
)

type TimeToStamp struct {
	from *timestamppb.Timestamp
	to   *timestamppb.Timestamp
}

func (s TimeToStamp) ConvertTime() *TimeToStamp {
	oneTrainingPeriod := time.Now().Add(-24 * time.Hour)
	s.from = timestamppb.New(oneTrainingPeriod)
	s.to = timestamppb.Now()
	return &s
}
func NewTimetoStamp() *TimeToStamp {
	return new(TimeToStamp)
}
func main() {
	s := NewTimetoStamp()
	historicalCandles, err := sdk.NewMarketDataService().GetCandles("BBG006L8G4H1", s.ConvertTime().from, s.ConvertTime().to, t.CandleInterval_CANDLE_INTERVAL_5_MIN)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(historicalCandles)

	mds := sdk.NewMarketDataStream()
	mds.Send(&t.MarketDataRequest{
		Payload: &t.MarketDataRequest_SubscribeCandlesRequest{
			SubscribeCandlesRequest: &t.SubscribeCandlesRequest{
				SubscriptionAction: t.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments: []*t.CandleInstrument{
					{
						Figi:     "BBG006L8G4H1",
						Interval: t.SubscriptionInterval_SUBSCRIPTION_INTERVAL_ONE_MINUTE,
					},
				},
			},
		},
	})

	for {
		recv, err := mds.Recv()
		if err != nil {
			panic(err)
		}

		fmt.Println(recv)

	}
}
