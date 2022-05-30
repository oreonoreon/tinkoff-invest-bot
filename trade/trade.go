package trade

import (
	"fmt"
	"log"
	"os"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/ema"
	"tinkoff-invest-bot/sdk"
)

var Sk *sdk.Services

func Run() {
	//log := loggy.GetLogger().Sugar()
	Sk = sdk.NewServices()
	EMA := ema.NewEma()
	EMA.EMAHistoric(Sk.Hiscandle(30, "BBG006L8G4H1", t.CandleInterval_CANDLE_INTERVAL_5_MIN), 20)
	sandBox()

	mds := Sk.MarketDataServiceStream
	err := mds.Send(marketDataStream_SubscribeCandles("BBG006L8G4H1", t.SubscriptionInterval_SUBSCRIPTION_INTERVAL_FIVE_MINUTES))

	if err != nil {
		log.Fatal(err)
	}
	//buyinstr()
	for {
		recv, err := mds.Recv()
		if err != nil {
			//panic(err)
			log.Println("ошибка во время стрима", err)
			time.Sleep(20 * time.Second)
			//<-context.Context()
		}

		fmt.Println(recv)

		EMA.EMAstream(recv.GetCandle(), 20)
		StrategyShort(time.Now(), EMA.Getema(), ema.Getprice())
	}
	//restart()
}
func restart() {
	var answer string
	fmt.Println("Do you want to Restart program or Exit?\nType Restart or Exit")
	fmt.Scanln(&answer)
	switch answer {
	case "Exit", "exit":
		os.Exit(0)
	case "Restart", "restart":
		Run()
	}

}

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

func marketDataStream_SubscribeLastPrice(figi string) *t.MarketDataRequest {
	return &t.MarketDataRequest{
		Payload: &t.MarketDataRequest_SubscribeLastPriceRequest{
			SubscribeLastPriceRequest: &t.SubscribeLastPriceRequest{
				SubscriptionAction: t.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments: []*t.LastPriceInstrument{
					{
						Figi: figi, //"BBG000B9XRY4", //"BBG006L8G4H1",
					},
				},
			},
		},
	}
}

func marketDataStream_SubscribeCandles(figi string, subscriptionInterval t.SubscriptionInterval) *t.MarketDataRequest {
	return &t.MarketDataRequest{
		Payload: &t.MarketDataRequest_SubscribeCandlesRequest{
			SubscribeCandlesRequest: &t.SubscribeCandlesRequest{
				SubscriptionAction: t.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments: []*t.CandleInstrument{
					{
						Figi:     figi,                 // "BBG006L8G4H1", // "BBG000B9XRY4", //"BBG006L8G4H1", "BBG000BDTBL9"
						Interval: subscriptionInterval, //t.SubscriptionInterval_SUBSCRIPTION_INTERVAL_FIVE_MINUTES,
					},
				},
			},
		},
	}
}

func marketDataStream_SubscribeInfo(figi string) *t.MarketDataRequest {
	return &t.MarketDataRequest{
		Payload: &t.MarketDataRequest_SubscribeInfoRequest{
			SubscribeInfoRequest: &t.SubscribeInfoRequest{
				SubscriptionAction: t.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments: []*t.InfoInstrument{
					{
						Figi: figi, //"BBG000B9XRY4",
					},
				},
			},
		},
	}
}
