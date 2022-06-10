package Controller

import (
	"fmt"
	"tinkoff-invest-bot/View"
	"tinkoff-invest-bot/ema"
	"tinkoff-invest-bot/sdk"
	"tinkoff-invest-bot/trade"
)

func Controller(command string) {
	switch command {
	//case "exit", "Exit":
	case "test":
		trade.Test_stratage()
	case "trade":
		trade.Run()
	//case "start real trade":
	case "etfema": //"help find instrument to trade"
		var symbol, interval string
		var avarageCoifecent float64
		View.ShowInfo("SPY 15min 20")
		fmt.Scanln(&symbol, &interval, &avarageCoifecent)
		ema.GetETFema(symbol, interval, avarageCoifecent, true)
	case "sandbox", "Sandbox", "SandBox": // аккаунт песочницы
		trade.AccSandBox()
	case "schedules": // полчаем все расписания торгов бирж
		Sk := sdk.NewServices() //todo убрать SK отсюда обьявлять его или в начале контролера и передавть его через аргумент функции или придумаать что то ещё, ТАК ОСТАВЛЯТЬ НЕ БЕЗОПАСНО!
		trade.GTHfaExchange(Sk)
		/*
			sheduleCboeSpx, _ := Sk.InstrumentsService.TradingSchedules("", timestamppb.Now(), timestamppb.New(time.Now().Add(24*time.Hour)))
			for _, value := range sheduleCboeSpx {
				fmt.Println(value.Exchange)
				fmt.Println(value.Days[0].Date.AsTime())
				fmt.Println(value.Days[0].EveningEndTime.AsTime())
				fmt.Println(value.Days[0].EndTime.AsTime())
				fmt.Println(value.Days[0].ClosingAuctionEndTime.AsTime())
				fmt.Println(value.Days[0].EveningOpeningAuctionStartTime.AsTime())
				fmt.Println(value.Days[0].PremarketEndTime.AsTime())
				fmt.Println(value.Days[0].EveningStartTime.AsTime())
				fmt.Println(value.Days[0].PremarketStartTime.AsTime())
				fmt.Println(value.Days[0].ClearingEndTime.AsTime())
				fmt.Println(value.Days[0].ClearingStartTime.AsTime())
				fmt.Println(value.Days[0].OpeningAuctionStartTime.AsTime())
				fmt.Println(value.Days[0].StartTime.AsTime())
				fmt.Println("")

			}
		*/
	}
}
