package ema

import (
	"fmt"
	"time"
	"tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/convertQuotation"
)

var ema float64 = 1703 // 1695.8
var price float64
var readytoshort bool
var accountBalance float64 = 10000
var sold bool

func EMA(data *investapi.GetCandlesResponse, n float64) {

	a := 2 / (n + 1)
	for _, value := range data.Candles {
		price = convertQuotation.Convert(value.Close.Units, value.Close.Nano) //цена закрытия
		//fmt.Println(value.Time.AsTime())
		ema = a*price + (1-a)*ema
		strategyShort(value.Time.AsTime())
		//fmt.Println(price)
	}
	fmt.Println(ema)
	fmt.Println(accountBalance)
	fmt.Println(i)
}

var i int

func strategyShort(time time.Time) {
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
	accountBalance += price
	sold = true
}
func buyinstr() {
	accountBalance -= price
	sold = false
}
