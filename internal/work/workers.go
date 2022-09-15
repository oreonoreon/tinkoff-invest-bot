package work

import (
	"fmt"
	"sync"
	"tinkoff-invest-bot/internal/Controller"
	"tinkoff-invest-bot/internal/View"
	"tinkoff-invest-bot/internal/bot"
	"tinkoff-invest-bot/internal/config"
	"tinkoff-invest-bot/internal/trade"
)

var Wg sync.WaitGroup
var i int

func Workers(xrr *bot.Xrr) {
	for id, user := range xrr.C.Users {
		if user.TokenSanBox != "" {
			Wg.Add(1)
			go RunTinkBot(id, user)
		}
	}
}
func Worker(xrr *bot.Xrr, id int64) {
	user := xrr.C.Users[id]
	if user.TokenSanBox != "" {
		Wg.Add(1)
		go RunTinkBot(id, user)
		xrr.C.SetWorkerOn(id, true)
	}

}

func RunTinkBot(id int64, user bot.User) {
	i++
	fmt.Println(i)
	config.Token = user.TokenSanBox
	tinkBot := trade.NewTinkBot(id, user.ChUser)
	defer Wg.Done()
	for {
		text := "Введи test или trade или sandbox или schedules, и может быть где-нибудь ты заработаешь;)"
		View.ShowInfo(id, text, "test", "trade", "sandbox", "stream", "schedules", "exit")
		command := <-user.ChUser
		Controller.Controller(tinkBot, fmt.Sprint(command))
	}
	//Wg.Done()
}
