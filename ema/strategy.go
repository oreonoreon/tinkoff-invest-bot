package ema

import (
	"fmt"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
)

//-------------------TESTING STRATEGY PART--------------------------------------

type gonabeaMillionaire struct {
	amountOfDeals  int     // число сделок
	accountBalance float64 //фейковый баланс для исторических данных

	readytoshort bool
	sold         bool

	etfAbovePrice   bool
	marketgonaclose bool
}

func History_testStrategyShort(SliceEma []Ema, Mapapa map[time.Time]EtfEma, schedule *t.TradingSchedule, test ...bool) {
	imMillionaire := new(gonabeaMillionaire)
	//fmt.Println("!LOOK HERE!", *imMillionaire, "!LOOK HERE!")
	for key, value := range SliceEma {
		if key < 300 {
			continue
		}
		//!value.time.Before(schedule.Days[0].StartTime.AsTime()) && !value.time.After(schedule.Days[0].EndTime.AsTime().Add(-20*time.Minute))
		//всё клёво да вот в schedule.Days[0] забита дата, тобижь день месяц и всёэто начиная с сейчас и в будущее)))
		h, _, _ := value.time.Clock()

		if (h >= 13) && (h <= 19) {
			if Mapapa[value.time].ema > Mapapa[value.time].price {
				imMillionaire.etfAbovePrice = true
			}
			shortornot(value, imMillionaire)
		}
	}
	if len(test) != 0 {
		fmt.Println("Account balance ", imMillionaire.accountBalance)
		fmt.Println("The amount of deals ", imMillionaire.amountOfDeals)
	}
}

func shortornot(value Ema, imMillionaire *gonabeaMillionaire) {
	if value.price > value.ema && imMillionaire.readytoshort == false {
		imMillionaire.readytoshort = true
	}
	if value.price < value.ema && imMillionaire.readytoshort && imMillionaire.etfAbovePrice {
		sellinstr(value.price, imMillionaire)
		imMillionaire.readytoshort = false
		fmt.Println("sell")
		fmt.Println("время сделки ", value.time) //время операции
		fmt.Println("цена сделки", value.price)
		fmt.Println("ema", value.ema) //ema на момент открытие сделки
		imMillionaire.amountOfDeals++
	}
	h, m, _ := value.time.Clock()
	if (value.price > value.ema && imMillionaire.sold) || (imMillionaire.sold && h == 19 && m == 45) {
		buyinstr(value.price, imMillionaire)
		fmt.Println("buyback")
		fmt.Println("время сделки ", value.time) //время операции
		fmt.Println("цена сделки", value.price)  //цена закрытия она же сделки
		fmt.Println("ema", value.ema)            //ema на момент закрытия сделки
		imMillionaire.readytoshort = true        //false
		imMillionaire.amountOfDeals++
	}

}

func sellinstr(price float64, imMillionaire *gonabeaMillionaire) {

	imMillionaire.accountBalance += price
	imMillionaire.sold = true
}
func buyinstr(price float64, imMillionaire *gonabeaMillionaire) {

	imMillionaire.accountBalance -= price
	imMillionaire.sold = false
}
