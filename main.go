package main

import (
	"fmt"
	"tinkoff-invest-bot/Controller"
	"tinkoff-invest-bot/View"
)

func main() {
	//test_stratage_on_historicalData()
	//sandbox_trade()

	//view
	//controler
	//model

	//trade.Run()
	for {
		View.ShowInfo("Введи test или trade или etfema или sandbox или schedules, и может быть где нибудь ты заработаешь;)")
		var command string
		fmt.Scanln(&command)
		Controller.Controller(command)
	}
}
