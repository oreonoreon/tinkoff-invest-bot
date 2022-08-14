package ema

import (
	"fmt"
	"time"
	"tinkoff-invest-bot/internal/View"
)

//-----------------------------------------------------------------------------------------------------------------------
func Test_strategy(id int64, SliceEma []Ema, etfmap map[time.Time]Price_Ema) {
	imMillionaire := new(gonabeaMillionaire)
	for _, value := range SliceEma {

		//----------------проверка открытия рынка через формат проверки time записаного в Ema и установки marketgonaclose в true или false--------------
		MarketOpen(value.time, imMillionaire)
		//---------------------ОПРЕДЕЛЯЕМ ВЫШЕ ИЛИ НИЖЕ ЦЕНА ETF-а(ПОЛЧУЧЕНОГО С СТОРОНЕГО API) ЕГО EMA -----------------------------------------------
		if etfmap[value.time].ema > etfmap[value.time].price {
			imMillionaire.etfunderprice = true
		} else {
			imMillionaire.etfunderprice = false
		}
		//-----------------------------------------------------------------------------------------------------------------
		shortornot(id, value, imMillionaire)
		longornot(id, value, imMillionaire)

	}
	View.ShowInfo(id, fmt.Sprintf("Account balance %v", imMillionaire.accountBalance))
	View.ShowInfo(id, fmt.Sprintf("TAX  %v", imMillionaire.tax))
	View.ShowInfo(id, fmt.Sprintf("The amount of deals  %v", imMillionaire.amountOfDeals))
	/*
		fmt.Println("Account balance ", imMillionaire.accountBalance)
		fmt.Println("The amount of deals ", imMillionaire.amountOfDeals)
	*/
}
func StreamTrade_strategy(id int64, value Ema, etfmap map[time.Time]Price_Ema) {
	imMillionaire := new(gonabeaMillionaire)

	//----------------проверка открытия рынка через формат проверки time записаного в Ema и установки marketgonaclose в true или false--------------
	MarketOpen(value.time, imMillionaire)
	//---------------------ОПРЕДЕЛЯЕМ ВЫШЕ ИЛИ НИЖЕ ЦЕНА ETF-а(ПОЛЧУЧЕНОГО С СТОРОНЕГО API) ЕГО EMA -----------------------------------------------
	if etfmap[value.time].ema > etfmap[value.time].price {
		imMillionaire.etfunderprice = true
	} else {
		imMillionaire.etfunderprice = false
	}
	//-----------------------------------------------------------------------------------------------------------------
	shortornot(id, value, imMillionaire)
	longornot(id, value, imMillionaire)

}
func MarketOpen(key time.Time, imMillionaire *gonabeaMillionaire) {

	year, mon, day := key.Date()
	loc := key.Location()
	if key.Before(time.Date(year, mon, day, 19, 45, 0, 0, loc)) && !key.Before(time.Date(year, mon, day, 13, 30, 0, 0, loc)) {
		imMillionaire.marketgonaclose = false
	} else {
		imMillionaire.marketgonaclose = true
	}

}

//-----------------------------------------------------------------------------------------------------------------------
//--------------------------------ЛОГИКА ШОРТ-СТРАТЕГИИ-----------------------------------------------
func shortornot2(value Ema, imMillionaire *gonabeaMillionaire) {
	if !imMillionaire.marketgonaclose {
		if value.price > value.ema && imMillionaire.readytoshort == false {
			imMillionaire.readytoshort = true
		}
		if value.price < value.ema && imMillionaire.readytoshort && imMillionaire.etfunderprice {
			sellinstr2(value.price, imMillionaire)
			//---------------------------------------------------
			imMillionaire.sold = true
			//---------------------------------------------------
			imMillionaire.readytoshort = false
			fmt.Println("short SELL")
			fmt.Println("время сделки ", value.time) //время операции
			fmt.Println("цена сделки", value.price)
			fmt.Println("ema", value.ema) //ema на момент открытие сделки
			imMillionaire.amountOfDeals++
		}

		if value.price > value.ema && imMillionaire.sold { //|| (imMillionaire.sold && value.time.Equal(time.Date(year, mon, day, 19, 45, 0, 0, loc)))
			buyinstr2(value.price, imMillionaire)
			//---------------------------------------------------
			imMillionaire.sold = false
			//---------------------------------------------------
			fmt.Println("short BUYBACK")
			fmt.Println("время сделки ", value.time) //время операции
			fmt.Println("цена сделки", value.price)  //цена закрытия она же сделки
			fmt.Println("ema", value.ema)            //ema на момент закрытия сделки
			imMillionaire.readytoshort = true        //false
			imMillionaire.amountOfDeals++
		}

	} else if imMillionaire.marketgonaclose && imMillionaire.sold {
		buyinstr2(value.price, imMillionaire)
		//---------------------------------------------------
		imMillionaire.sold = false
		//---------------------------------------------------
		fmt.Println("Close short position AT THE END OF THE DAY")
		fmt.Println("время сделки ", value.time) //время операции
		fmt.Println("цена сделки", value.price)  //цена закрытия она же сделки
		fmt.Println("ema", value.ema)            //ema на момент закрытия сделки
		imMillionaire.readytoshort = true        //false
		imMillionaire.amountOfDeals++
	}

}

//--------------------------------ЛОГИКА ЛОНГ-СТРАТЕГИИ--------------------------------------------------------
func longornot2(value Ema, imMillionaire *gonabeaMillionaire) {
	if !imMillionaire.marketgonaclose {
		if value.price < value.ema && imMillionaire.readytoLong == false {
			imMillionaire.readytoLong = true
		}
		if value.price > value.ema && imMillionaire.readytoLong && !imMillionaire.etfunderprice {
			buyinstr2(value.price, imMillionaire)
			//---------------------------------------------------
			imMillionaire.bought = true
			//---------------------------------------------------
			imMillionaire.readytoLong = false
			fmt.Println("long BUY")
			fmt.Println("время сделки ", value.time) //время операции
			fmt.Println("цена сделки", value.price)
			fmt.Println("ema", value.ema) //ema на момент открытие сделки
			imMillionaire.amountOfDeals++
		}

		if value.price < value.ema && imMillionaire.bought { //todo подумать подходит ли sold маркер для лонг стратегии
			sellinstr2(value.price, imMillionaire)
			//---------------------------------------------------
			imMillionaire.bought = false
			//---------------------------------------------------
			fmt.Println("long Close position")
			fmt.Println("время сделки ", value.time) //время операции
			fmt.Println("цена сделки", value.price)  //цена закрытия она же сделки
			fmt.Println("ema", value.ema)            //ema на момент закрытия сделки
			imMillionaire.readytoLong = true         //false
			imMillionaire.amountOfDeals++
		}

	} else if imMillionaire.marketgonaclose && imMillionaire.bought {
		sellinstr2(value.price, imMillionaire)
		//---------------------------------------------------
		imMillionaire.bought = false
		//---------------------------------------------------
		fmt.Println("Close position AT THE END OF THE DAY")
		fmt.Println("время сделки ", value.time) //время операции
		fmt.Println("цена сделки", value.price)  //цена закрытия она же сделки
		fmt.Println("ema", value.ema)            //ema на момент закрытия сделки
		imMillionaire.readytoLong = true         //false
		imMillionaire.amountOfDeals++
	}
}

func sellinstr2(price float64, imMillionaire *gonabeaMillionaire) {

	imMillionaire.accountBalance += price
	//imMillionaire.sold = true
}
func buyinstr2(price float64, imMillionaire *gonabeaMillionaire) {

	imMillionaire.accountBalance -= price
	//imMillionaire.sold = false
}
