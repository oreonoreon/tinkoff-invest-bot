package ema

import (
	"fmt"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/convertQuotation"
)

//var ema float64 //= 1642.3 // 1695.8
/*
var readytoshort bool
var accountBalance float64 = 10000
var sold bool
*/
var price float64

type Ema struct {
	ema float64
}

func NewEma() *Ema {
	return new(Ema)
}
func (e *Ema) EMAHistoric(data []*t.HistoricCandle, n float64) {
	e.ema = convertQuotation.Convert(data[0].Close.Units, data[0].Close.Nano)
	a := 2 / (n + 1)
	for _, value := range data {
		if !value.IsComplete {
			break
		}
		price = convertQuotation.Convert(value.Close.Units, value.Close.Nano) //цена закрытия
		//fmt.Println(value.Time.AsTime())
		e.ema = a*price + (1-a)*e.ema
		//strategyShort(value.Time.AsTime())
		//fmt.Println(price)
	}
	fmt.Println(e.ema)
	//fmt.Println(accountBalance)
	//fmt.Println(i)
}
func (e *Ema) EMAstream(candle *t.Candle, n float64) {
	a := 2 / (n + 1)
	if candle == nil {
		return
	}
	price = convertQuotation.Convert(candle.GetClose().GetUnits(), candle.GetClose().GetNano()) //цена закрытия
	//fmt.Println(value.Time.AsTime())
	e.ema = a*price + (1-a)*e.ema
	//stratagy.StrategyShort(candle.Time.AsTime(), e.ema, price)
	//fmt.Println(price)
	fmt.Println(e.ema)

}
func (e Ema) Getema() float64 {
	return e.ema
}
func Getprice() float64 {
	return price
}

/*
var i int

func strategyShort(time time.Time, ema float64) {
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
	saccount, err := trade.Sk.SandboxService.GetSandboxAccounts()
	if err != nil {
		log.Fatalln("GetSandboxAccounts", err)
	}
	trade.Sk.SandboxService.PostSandboxOrder(&t.PostOrderRequest{
		Figi:      "BBG006L8G4H1", //TODO всавить уже функцию выбора инструмента ато задрало вставляти фиги Яндекса
		Quantity:  1,
		Direction: t.OrderDirection_ORDER_DIRECTION_SELL,
		AccountId: saccount[0].Id,
		OrderType: t.OrderType_ORDER_TYPE_MARKET,
		OrderId:   uuid.New().String(),
	})
	fmt.Println(trade.Sk.SandboxService.GetSandboxOrders(saccount[0].Id))
	//accountBalance += price
	sold = true
}
func buyinstr() {
	saccount, err := trade.Sk.SandboxService.GetSandboxAccounts()
	if err != nil {
		log.Fatalf("GetSandboxAccounts", err)
	}
	trade.Sk.SandboxService.PostSandboxOrder(&t.PostOrderRequest{
		Figi:      "BBG006L8G4H1", //TODO всавить уже функцию выбора инструмента ато задрало вставляти фиги Яндекса
		Quantity:  1,
		Direction: t.OrderDirection_ORDER_DIRECTION_BUY,
		AccountId: saccount[0].Id,
		OrderType: t.OrderType_ORDER_TYPE_MARKET,
		OrderId:   uuid.New().String(),
	})
	fmt.Println(trade.Sk.SandboxService.GetSandboxOrders(saccount[0].Id))
	//accountBalance -= price
	sold = false
}
*/
