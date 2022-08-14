package ema

import (
	"fmt"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/internal/convertQuotation"
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
	e.ema = convertQuotation.Convert(data[0].Open.Units, data[0].Open.Nano) //самую первую ячейку слайса []*t.HistoricCandle полужим равной ема инструмента, как начальную точку расчёта
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

		SliceEma = append(SliceEma, *e)
	}
	//-------------------------------------------------------------------------------------------------------------
	//View.ShowInfo(SliceEma)
	//View.ShowInfo(id, bot.BOT, fmt.Sprintf("Ema %v", e.ema))

}
func (e *Ema) EMAstream(candle *t.Candle, averageCoefficient float64) {
	if candle == nil {
		return
	}
	a := 2 / (averageCoefficient + 1)
	e.price = convertQuotation.Convert(candle.GetClose().GetUnits(), candle.GetClose().GetNano()) //цена закрытия
	e.time = candle.GetTime().AsTime()
	e.ema = a*e.price + (1-a)*e.ema
	fmt.Println(e.ema)

}
func (e *Ema) EMAstreamETF(lastprice *t.LastPrice, avarageCoifecent float64) {
	if lastprice == nil {
		return
	}
	a := 2 / (avarageCoifecent + 1)
	e.price = convertQuotation.Convert(lastprice.GetPrice().GetUnits(), lastprice.GetPrice().GetNano()) //цена закрытия
	e.time = lastprice.Time.AsTime()
	e.ema = a*e.price + (1-a)*e.ema

	//fmt.Println(price)
	fmt.Println(e.ema)
	lastprice.ProtoReflect().Descriptor()
}

func (e Ema) Getema() float64 {
	return e.ema
}
func (e Ema) Getprice() float64 {
	return e.price
}
func (e *Ema) Setema(newema float64) float64 {
	e.ema = newema
	return e.ema
}
func (e *Ema) Setprice(newprice float64) float64 {
	e.price = newprice
	return e.price
}
