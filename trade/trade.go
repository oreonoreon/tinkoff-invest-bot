package trade

import (
	"fmt"
	"log"
	"os"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/View"
	"tinkoff-invest-bot/ema"
	"tinkoff-invest-bot/sdk"
)

var Sk *sdk.Services // TODO РЕШИТЬ ЧТО ТО С ОБЬЯВЛЕНИЕМ SK, ТАК ОСТАВЛЯТЬ ОПАСНО И НЕПРАВЕЛЬНО!

func Run() {
	//log := loggy.GetLogger().Sugar()
	Sk = sdk.NewServices()
	var figi = "BBG000B9XRY4"
	EMA := ema.NewEma()
	EMA.EMAHistoric(Sk.Hiscandle(30, figi, t.CandleInterval_CANDLE_INTERVAL_5_MIN), 20)
	sandBox()

	mds := Sk.MarketDataServiceStream
	err := mds.Send(marketDataStream_SubscribeCandles(figi, t.SubscriptionInterval_SUBSCRIPTION_INTERVAL_FIVE_MINUTES))

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
		}

		View.ShowInfo(recv) //fmt.Println(recv)

		EMA.EMAstream(recv.GetCandle(), 20)
		StrategyShort(time.Now(), EMA.Getema(), EMA.Getprice())
	}
	//restart(Run)
}

/*
func restart(xfunc func()) {
	var answer string
	fmt.Println("Do you want to Restart program or Exit?\nType Restart or Exit")
	fmt.Scanln(&answer)
	switch answer {
	case "Exit", "exit":
		os.Exit(0)
	case "Restart", "restart":
		xfunc() //Run()
	}

}*/
func restart() {
	var answer string
	fmt.Println("Do you want to Restart program or Exit?\nType Restart or Exit")
	fmt.Scanln(&answer)
	switch answer {
	case "Exit", "exit":
		os.Exit(0)
	default:
		View.ShowInfo("RESTART")
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
	View.ShowInfo(sandboxPortfolio) //fmt.Println(sandboxPortfolio)
	//fmt.Println(Sk.SandboxService.GetSandboxAccounts())
}
func AccSandBox() {
	if Sk == nil {
		Sk = sdk.NewServices()
	}
	account, err := Sk.SandboxService.GetSandboxAccounts()
	if err != nil {
		View.ShowInfo(err)
	}
	portfolio, err := Sk.SandboxService.GetSandboxPortfolio(account[0].Id)
	if err != nil {
		View.ShowInfo(err)
	}
	View.ShowInfo(portfolio)
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
						Figi:     figi,                 // "BBG006L8G4H1", // "BBG000B9XRY4", "BBG000BDTBL9"
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
