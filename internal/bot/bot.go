package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"time"
)

var duration = time.Minute
var BOT *tgbotapi.BotAPI

func newBot() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	os.Setenv("token", halyava007BotToken)
	bot, err := tgbotapi.NewBotAPI(os.Getenv("token"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	return bot, updates
}
func Bot() *Xrr {
	//конектимся к телеграм и получаем апдейты
	bot, updates := newBot()
	//создаём кеш
	c := New()
	xrr := newXrr(bot, c, updates)
	//загружаем в кеш Юзеров из Excel
	loadFromExcel(c)

	//запись кэша в excel файл
	callAtStartOfProgram(c.writeCacheInFile, 5*duration)

	// Loop through each update.
	//getUpdates(updates, bot, c)

	//xrr.GetUp()

	//View.ChXrr <- *xrr
	//View.Wg.Wait()
	return xrr
}

//func (x Xrr) GetUp() {
//	go func() {
//		for {
//			text := <-View.Ch
//			if msg, test := text.(tgbotapi.MessageConfig); test {
//				if _, err := x.BOT.Send(msg); err != nil {
//					log.Println(err)
//				}
//			} else {
//				fmt.Println("it's not tgbotapi.MessageConfig")
//			}
//		}
//	}()
//
//	for update := range x.Updates {
//		caseIncomingMessage(update, x.BOT, x.C)
//		callbackQueryCase(update, x.BOT, x.C)
//	}
//}
