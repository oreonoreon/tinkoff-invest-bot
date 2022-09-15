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

	return xrr
}
