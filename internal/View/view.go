package View

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tinkoff-invest-bot/internal/bot"
)

func ShowInfo(Id int64, info ...interface{}) {
	msg := tgbotapi.NewMessage(Id, "")
	if len(info) != 1 {
		var keyboard [][]tgbotapi.InlineKeyboardButton
		for k, v := range info {
			var row []tgbotapi.InlineKeyboardButton
			if k == 0 {
				msg.Text = fmt.Sprint(v)
			} else {
				if value, ok := v.(string); ok {
					button := tgbotapi.NewInlineKeyboardButtonData(value, value)
					row = append(row, button)
				} else {
					fmt.Println("it's not a string")
				}
				keyboard = append(keyboard, row)
			}
		}
		board := tgbotapi.InlineKeyboardMarkup{
			InlineKeyboard: keyboard,
		}
		msg.ReplyMarkup = board
	} else {
		msg.Text = fmt.Sprint(info[0])
	}

	if _, err := bot.BOT.Send(msg); err != nil {
		log.Println(err)
	}

}
