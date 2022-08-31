package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tinkoff-invest-bot/internal/bot"
	"tinkoff-invest-bot/internal/ema"
	"tinkoff-invest-bot/internal/work"
)

func main() {

	//getHistoryData.Run("figi.txt")
	ema.ReadWriteCSV("BBG000B9XRY4")
	ema.NewTimeFrame(30)
	ema.NewTimeFrame(5)
	ema.NewTimeFrame(15)
	//fmt.Println(ema.Candles)

	xrr := bot.Bot()
	bot.BOT = xrr.BOT
	work.Wg.Add(1)
	go func() {
		for update := range xrr.Updates {
			xrr.C.CheckNewUser(update)
			if xrr.C.Users[update.SentFrom().ID].TokenSanBox != "" {
				switch {
				case update.Message != nil:
					if update.Message.Command() == "start" && !xrr.C.Users[update.SentFrom().ID].WorkerOn {
						work.Worker(xrr, update.SentFrom().ID)
						//fmt.Println(xrr.C.Users[update.SentFrom().ID].WorkerOn)
					}
					//xrr.C.Users[update.SentFrom().ID].ChUser <- update.Message.Text
				case update.CallbackQuery != nil:
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
					if _, err := xrr.BOT.Request(callback); err != nil {
						log.Println(err)
					}
					msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, update.CallbackQuery.Data)
					msg.ParseMode = "HTML"
					if _, err := xrr.BOT.Request(msg); err != nil {
						log.Println(err)
					}
					xrr.C.Users[update.SentFrom().ID].ChUser <- update.CallbackQuery.Data
				}
			} else {
				switch update.Message.Command() {
				case "start":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
					msg.Text = "Введи токен"
					msg.ParseMode = "HTML"
					xrr.C.SetMessageId(update.SentFrom().ID, update.Message.MessageID)

					if _, err := xrr.BOT.Send(msg); err != nil {
						log.Println(err)
					}
				}
				if update.Message.MessageID == xrr.C.Users[update.SentFrom().ID].MessageId+2 {
					xrr.C.SetToken(update.SentFrom().ID, update.Message.Text)
					work.Worker(xrr, update.SentFrom().ID)
				}
			}
		}
	}()

	work.Wg.Wait()
}

//TODO 1. Сделать тест стратегии по данным со сторонего API
//        + а) Взять ema ETF и его цену определяя выше ценна ETF чем ema или нет (закинув в ema.History_testStrategyShort файл strategy)
//        + б) Взять ema ИНСТРУМНТА сравнить его с его ценой определяя выше ценна ИНСТРУМЕНТА чем ema или нет (закинув в ema.History_testStrategyShort файл strategy)
//         в) По итогам двух выше указанных пунктов покупать продавать инструмент в Тинькофф
//TODO 2. Подключить релизацию стратегии со стриминга данных
