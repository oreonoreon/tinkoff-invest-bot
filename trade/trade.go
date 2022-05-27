package trade

import (
	"fmt"
	"log"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/ema"
	"tinkoff-invest-bot/sdk"
)

var Sk *sdk.Services

func sandBox() {
	sadboxaccount, err := Sk.SandboxService.GetSandboxAccounts()
	if err != nil {
		log.Fatal("sadboxaccount", err)
	}
	accountId := sadboxaccount[0].Id
	sandboxPortfolio, err := Sk.SandboxService.GetSandboxPortfolio(accountId)
	if err != nil {
		log.Fatal("sandboxPortfolio", err)
	}
	if sandboxPortfolio.TotalAmountCurrencies.Units == 0 {
		_, err := Sk.SandboxService.SandboxPayIn(accountId, &t.MoneyValue{
			Currency: "rub",
			Units:    10000,
			Nano:     0,
		})
		if err != nil {
			log.Fatalln("SandboxPayIn ", err)
		}
	}
	fmt.Println(sandboxPortfolio)
	//fmt.Println(Sk.SandboxService.GetSandboxAccounts())
}

func Run() {
	//log := loggy.GetLogger().Sugar()
	Sk = sdk.NewServices()
	EMA := ema.NewEma()
	EMA.EMAHistoric(Sk.Hiscandle(7), 20)
	sandBox()

	mds := Sk.MarketDataServiceStream
	err := mds.Send(&t.MarketDataRequest{
		Payload: &t.MarketDataRequest_SubscribeCandlesRequest{
			SubscribeCandlesRequest: &t.SubscribeCandlesRequest{
				SubscriptionAction: t.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments: []*t.CandleInstrument{
					{
						Figi:     "BBG006L8G4H1", // "BBG000B9XRY4", //"BBG006L8G4H1",
						Interval: t.SubscriptionInterval_SUBSCRIPTION_INTERVAL_FIVE_MINUTES,
					},
				},
			},
		},
	})

	/*
		mds.Send(&t.MarketDataRequest{
			Payload: &t.MarketDataRequest_SubscribeLastPriceRequest{
				SubscribeLastPriceRequest: &t.SubscribeLastPriceRequest{
					SubscriptionAction: t.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
					Instruments: []*t.LastPriceInstrument{
						{
							Figi: "BBG000B9XRY4", //"BBG006L8G4H1",
						},
					},
				},
			},
		})
	*/
	/*
		mds.Send(&t.MarketDataRequest{
			Payload: &t.MarketDataRequest_SubscribeInfoRequest{
				SubscribeInfoRequest: &t.SubscribeInfoRequest{
					SubscriptionAction: t.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
					Instruments: []*t.InfoInstrument{
						{
							Figi: "BBG000B9XRY4",
						},
					},
				},
			},
		})
	*/

	if err != nil {
		log.Fatal(err)
	}

	for {
		recv, err := mds.Recv()
		if err != nil {
			panic(err)
		}

		//fmt.Println(recv.GetCandle().GetClose().GetUnits())
		EMA.EMAstream(recv.GetCandle(), 20)
		StrategyShort(time.Now(), EMA.Getema(), ema.Getprice())
	}
}
