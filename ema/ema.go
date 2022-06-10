package ema

import (
	"fmt"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/View"
	"tinkoff-invest-bot/convertQuotation"
)

type Ema struct {
	ema   float64
	price float64
	time  time.Time
}

func NewEma() *Ema {
	return new(Ema)
}

var SliceEma []Ema

func (e *Ema) EMAHistoric(data []*t.HistoricCandle, avarageCoifecent float64) {
	e.ema = convertQuotation.Convert(data[0].Open.Units, data[0].Open.Nano)
	a := 2 / (avarageCoifecent + 1)

	SliceEma = make([]Ema, 0)

	//---------------------------------ПРОБЕЖИМСЯ ПО СЛАЙСУ ИСТОРИЧЕСКИХ СВЕЧЕЙ ВЫЧИСЛЯЯ EMA-------------------------------------------------------------
	for _, value := range data {
		if !value.IsComplete {
			break
		}
		if value.Time.AsTime().Weekday() == time.Saturday || value.Time.AsTime().Weekday() == time.Sunday {
			continue
		}
		e.price = convertQuotation.Convert(value.Close.Units, value.Close.Nano) //цена закрытия
		e.ema = a*e.price + (1-a)*e.ema
		e.time = value.Time.AsTime()
		/*
			if len(test) != 0 {
				e.strategyShort(value.Time.AsTime(), e.ema)
				View.ShowInfo(e.ema)                                                                                   //fmt.Println(e.ema)
				View.ShowInfo(fmt.Sprintf("Время закрытия свечи %v и цена закрытия %v", value.Time.AsTime(), e.price)) //fmt.Println(value.Time.AsTime(), price)
			}
		*/

		SliceEma = append(SliceEma, *e)
	}
	//-------------------------------------------------------------------------------------------------------------
	//View.ShowInfo(SliceEma)
	View.ShowInfo(fmt.Sprintf("Ema %v", e.ema))

}
func (e *Ema) EMAstream(candle *t.Candle, avarageCoifecent float64) {
	a := 2 / (avarageCoifecent + 1)
	if candle == nil {
		return
	}
	e.price = convertQuotation.Convert(candle.GetClose().GetUnits(), candle.GetClose().GetNano()) //цена закрытия
	//fmt.Println(value.Time.AsTime())
	e.ema = a*e.price + (1-a)*e.ema
	//stratagy.StrategyShort(candle.Time.AsTime(), e.ema, price)
	//fmt.Println(price)
	fmt.Println(e.ema)

}
func (e Ema) Getema() float64 {
	return e.ema
}
func (e Ema) Getprice() float64 {
	return e.price
}

/*
var i int // число сделок
var accountBalance float64 = 10000 //фейковй баланс для исторических данных

var readytoshort bool
var sold bool

func (e *Ema) strategyShort(time time.Time, ema float64) {
	if e.price > ema && e.readytoshort == false {
		e.readytoshort = true
	}
	if e.price < ema && e.readytoshort {
		e.sellinstr()
		e.readytoshort = false
		fmt.Println("sell")
		fmt.Println("время сделки ", time) //время операции
		fmt.Println("цена сделки", e.price)
		e.i++
	}
	if e.price > ema && e.sold {
		e.buyinstr()
		fmt.Println("buyback")
		fmt.Println("время сделки ", time)  //время операции
		fmt.Println("цена сделки", e.price) //цена закрытия она же сделки
		e.readytoshort = true               //false
		e.i++
	}
}

func (e *Ema) sellinstr() {

	accountBalance += e.price
	e.sold = true
}
func (e *Ema) buyinstr() {

	accountBalance -= e.price
	e.sold = false
}
*/
