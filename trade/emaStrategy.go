package trade

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/loggy"
)

var i int
var readytoshort bool
var sold bool

func StrategyShort(time time.Time, ema float64, price float64) {
	if price > ema && readytoshort == false {
		readytoshort = true
	}
	if price < ema && readytoshort {
		sellinstr()
		readytoshort = false
		fmt.Println("sell")
		fmt.Println(time) //время операции
		fmt.Println(price)
		i++
	}
	if price > ema && sold {
		buyinstr()
		fmt.Println("buyback")
		fmt.Println(time)   //время операции
		fmt.Println(price)  //цена закрытия она же сделки
		readytoshort = true //false
		i++
	}
}

func sellinstr() {

	saccount, err := Sk.SandboxService.GetSandboxAccounts()
	if err != nil {
		log.Fatalln("GetSandboxAccounts", err)
	}
	postorderResp, err := Sk.SandboxService.PostSandboxOrder(&t.PostOrderRequest{
		Figi:      "BBG006L8G4H1", //TODO всавить уже функцию выбора инструмента ато задрало вставляти фиги Яндекса
		Quantity:  1,
		Direction: t.OrderDirection_ORDER_DIRECTION_SELL,
		AccountId: saccount[0].Id,
		OrderType: t.OrderType_ORDER_TYPE_MARKET,
		OrderId:   uuid.New().String(),
	})
	if err != nil {
		loggy.GetLogger().Sugar().Warn("postorderResp ", err)
	}
	fmt.Println(postorderResp)
	//accountBalance += price
	sold = true
}
func buyinstr() {

	saccount, err := Sk.SandboxService.GetSandboxAccounts()
	if err != nil {
		log.Fatalf("GetSandboxAccounts", err)
	}
	postorderResp, err := Sk.SandboxService.PostSandboxOrder(&t.PostOrderRequest{
		Figi:      "BBG006L8G4H1", //TODO всавить уже функцию выбора инструмента ато задрало вставляти фиги Яндекса
		Quantity:  1,
		Direction: t.OrderDirection_ORDER_DIRECTION_BUY,
		AccountId: saccount[0].Id,
		OrderType: t.OrderType_ORDER_TYPE_MARKET,
		OrderId:   uuid.New().String(),
	})
	if err != nil {
		loggy.GetLogger().Sugar().Warn("postorderResp ", err)
	}
	fmt.Println(postorderResp)
	fmt.Println(Sk.SandboxService.GetSandboxOrders(saccount[0].Id))
	//accountBalance -= price
	sold = false
}
