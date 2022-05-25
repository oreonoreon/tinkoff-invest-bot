package main

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/ema"
	"tinkoff-invest-bot/sdk"
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
func main() {
	/*
		hc := hiscandle(7)
		for _, value := range hc {
			fmt.Println(value.Time.AsTime())
		}
	*/
	ema.EMA(hiscandle(7), 50)

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
