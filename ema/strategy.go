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
		//----------------проверка открытия рынка через формат проверки time записаного в Ema и установки marketgonaclose в true или false--------------
		checkTheMarketOpen(value, schedule, imMillionaire)

		if Mapapa[value.time].ema > Mapapa[value.time].price {
			imMillionaire.etfAbovePrice = true
		}
		shortornot(value, imMillionaire)

		//-----------------------------------------------------------------------------------------------------------------
		/*
			h, _, _ := value.time.Clock()
			if (h >= 13) && (h <= 19) {
				if Mapapa[value.time].ema > Mapapa[value.time].price {
					imMillionaire.etfAbovePrice = true
				}
				shortornot(value, imMillionaire)
			}
		*/
	}
	if len(test) != 0 {
		fmt.Println("Account balance ", imMillionaire.accountBalance)
		fmt.Println("The amount of deals ", imMillionaire.amountOfDeals)
	}
}

//checkTheMarketOpen возьмёт время из структуры Ema сравнит с schedule (расписанием открытий и закрытий рынка) и выдаст результат
func checkTheMarketOpen(value Ema, schedule *t.TradingSchedule, imMillionaire *gonabeaMillionaire) {

	year, mon, day := value.time.Date()
	loc := value.time.Location()
	if value.time.Before(time.Date(year, mon, day, 19, 45, 0, 0, loc)) && !value.time.Before(time.Date(year, mon, day, 13, 30, 0, 0, loc)) {
		imMillionaire.marketgonaclose = false
	} else {
		imMillionaire.marketgonaclose = true
	}

	/*
		if value.time.Hour() >= schedule.Days[0].StartTime.AsTime().Hour() && value.time.Hour() <= schedule.Days[0].EndTime.AsTime().Add(-20*time.Minute).Hour() {
			imMillionaire.marketgonaclose = true
		} else {
			imMillionaire.marketgonaclose = false
		}
	*/
	//!value.time.Before(schedule.Days[0].StartTime.AsTime()) && !value.time.After(schedule.Days[0].EndTime.AsTime().Add(-20*time.Minute))
}

//--------------------------------ЛОГИКА ШОРТ-СТРАТЕГИИ-----------------------------------------------
func shortornot(value Ema, imMillionaire *gonabeaMillionaire) {
	if !imMillionaire.marketgonaclose {
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
		/*
			year, mon, day := value.time.Date()
			loc := value.time.Location()
		*/
		if value.price > value.ema && imMillionaire.sold { //|| (imMillionaire.sold && value.time.Equal(time.Date(year, mon, day, 19, 45, 0, 0, loc)))
			buyinstr(value.price, imMillionaire)
			fmt.Println("buyback")
			fmt.Println("время сделки ", value.time) //время операции
			fmt.Println("цена сделки", value.price)  //цена закрытия она же сделки
			fmt.Println("ema", value.ema)            //ema на момент закрытия сделки
			imMillionaire.readytoshort = true        //false
			imMillionaire.amountOfDeals++
		}
		/*
			h, m, _ := value.time.Clock()
			if (value.price > value.ema && imMillionaire.sold) || (imMillionaire.sold && h == 19 && m == 45) { //todo исправить h==19 && m==45 будет пропускать время  болеше него если свечи меньши 15 минут
				buyinstr(value.price, imMillionaire)
				fmt.Println("buyback")
				fmt.Println("время сделки ", value.time) //время операции
				fmt.Println("цена сделки", value.price)  //цена закрытия она же сделки
				fmt.Println("ema", value.ema)            //ema на момент закрытия сделки
				imMillionaire.readytoshort = true        //false
				imMillionaire.amountOfDeals++
			}
		*/
	} else if imMillionaire.marketgonaclose && imMillionaire.sold {
		buyinstr(value.price, imMillionaire)
		fmt.Println("buyback")
		fmt.Println("время сделки ", value.time) //время операции
		fmt.Println("цена сделки", value.price)  //цена закрытия она же сделки
		fmt.Println("ema", value.ema)            //ema на момент закрытия сделки
		imMillionaire.readytoshort = true        //false
		imMillionaire.amountOfDeals++
	}

}

//----------------------------------------------------------------------------------------------------
func longornot(value Ema, imMillionaire *gonabeaMillionaire) {

}

func sellinstr(price float64, imMillionaire *gonabeaMillionaire) {

	imMillionaire.accountBalance += price
	imMillionaire.sold = true
}
func buyinstr(price float64, imMillionaire *gonabeaMillionaire) {

	imMillionaire.accountBalance -= price
	imMillionaire.sold = false
}
