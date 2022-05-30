package ema

import (
	"fmt"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/convertQuotation"
)

//var ema float64

var readytoshort bool
var accountBalance float64 = 10000
var sold bool

var price float64

type Ema struct {
	ema float64
}

func NewEma() *Ema {
	return new(Ema)
}
func (e *Ema) EMAHistoric(data []*t.HistoricCandle, n float64) {
	e.ema = convertQuotation.Convert(data[0].Open.Units, data[0].Open.Nano)
	a := 2 / (n + 1)
	for _, value := range data {
		if !value.IsComplete {
			break
		}
		if value.Time.AsTime().Weekday() == time.Saturday || value.Time.AsTime().Weekday() == time.Sunday {
			continue
		}
		price = convertQuotation.Convert(value.Close.Units, value.Close.Nano) //цена закрытия
		e.ema = a*price + (1-a)*e.ema
		strategyShort(value.Time.AsTime(), e.ema)
		fmt.Println(e.ema)
		fmt.Println(value.Time.AsTime(), price)
	}
	fmt.Println("Ema ", e.ema)
	fmt.Println("Account balance ", accountBalance)
	fmt.Println("The amount of deals ", i)
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

var i int

func strategyShort(time time.Time, ema float64) {
	if price > ema && readytoshort == false {
		readytoshort = true
	}
	if price < ema && readytoshort {
		sellinstr()
		readytoshort = false
		fmt.Println("sell")
		fmt.Println("время сделки ", time) //время операции
		fmt.Println("цена сделки", price)
		i++
	}
	if price > ema && sold {
		buyinstr()
		fmt.Println("buyback")
		fmt.Println("время сделки ", time) //время операции
		fmt.Println("цена сделки", price)  //цена закрытия она же сделки
		readytoshort = true                //false
		i++
	}
}

func sellinstr() {

	accountBalance += price
	sold = true
}
func buyinstr() {

	accountBalance -= price
	sold = false
}
