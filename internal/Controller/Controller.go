package Controller

import (
	"tinkoff-invest-bot/internal/trade"
)

func Controller(tinkBot trade.TinkBot, command string) {

	switch command {
	case "exit", "Exit":
		tinkBot.Restart()
	case "test":
		tinkBot.Test_stratage()

	case "trade":
		tinkBot.StartTrade("BBG000B9XRY4", "BBG000BDTBL9", "15min", 50)
	case "sandbox", "Sandbox", "SandBox": // аккаунт песочницы
		tinkBot.AccSandBox()
	case "test2":
		tinkBot.Test_S_With_Data_from_outerApi()
	case "stream":
		tinkBot.StreamDataLastPrice()
	case "schedules": // полчаем все расписания торгов бирж

		tinkBot.GTHfaExchange()
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
