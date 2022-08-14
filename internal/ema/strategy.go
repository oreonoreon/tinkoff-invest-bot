package ema

import (
	"fmt"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/internal/View"
)

//-------------------TESTING STRATEGY PART--------------------------------------

type gonabeaMillionaire struct {
	amountOfDeals  int     // число сделок
	tax            float64 // комиссия
	accountBalance float64 //фейковый баланс для исторических данных

	readytoshort bool
	readytoLong  bool
	sold         bool
	bought       bool

	etfunderprice   bool
	marketgonaclose bool
}

var debug bool

func History_testStrategyShort(id int64, SliceEma []Ema, Mapapa map[time.Time]Price_Ema, schedule *t.TradingSchedule, test ...bool) {
	imMillionaire := new(gonabeaMillionaire)
	for key, value := range SliceEma {
		if key < 300 {
			continue
		}
		//----------------проверка открытия рынка через формат проверки time записаного в Ema и установки marketgonaclose в true или false--------------
		checkTheMarketOpen(value, schedule, imMillionaire)
		//---------------------ОПРЕДЕЛЯЕМ ВЫШЕ ИЛИ НИЖЕ ЦЕНА ETF-а(ПОЛЧУЧЕНОГО С СТОРОНЕГО API) ЕГО EMA -----------------------------------------------
		if Mapapa[value.time].ema > Mapapa[value.time].price {
			imMillionaire.etfunderprice = true
		} else {
			imMillionaire.etfunderprice = false
		}
		//-----------------------------------------------------------------------------------------------------------------
		shortornot(id, value, imMillionaire)
		longornot(id, value, imMillionaire)

	}
	if len(test) != 0 {
		View.ShowInfo(id, fmt.Sprintf("Account balance %v", imMillionaire.accountBalance))
		View.ShowInfo(id, fmt.Sprintf("TAX  %v", imMillionaire.tax))
		View.ShowInfo(id, fmt.Sprintf("The amount of deals  %v", imMillionaire.amountOfDeals))
		//fmt.Println("Account balance ", imMillionaire.accountBalance)
		//fmt.Println("The amount of deals ", imMillionaire.amountOfDeals)
	}
}

//checkTheMarketOpen возьмёт время из структуры Ema сравнит с установлеными нами временными интервалами работы биржы или с //schedule (расписанием открытий и закрытий рынка)// и выдаст результат
func checkTheMarketOpen(value Ema, schedule *t.TradingSchedule, imMillionaire *gonabeaMillionaire) {

	year, mon, day := value.time.Date()
	loc := value.time.Location()
	if value.time.Before(time.Date(year, mon, day, 19, 45, 0, 0, loc)) && !value.time.Before(time.Date(year, mon, day, 13, 30, 0, 0, loc)) {
		imMillionaire.marketgonaclose = false
	} else {
		imMillionaire.marketgonaclose = true
	}

}

//--------------------------------ЛОГИКА ШОРТ-СТРАТЕГИИ-----------------------------------------------
func shortornot(id int64, value Ema, imMillionaire *gonabeaMillionaire) {
	if !imMillionaire.marketgonaclose {
		if value.price > value.ema && imMillionaire.readytoshort == false {
			imMillionaire.readytoshort = true
		}
		if value.price < value.ema && imMillionaire.readytoshort && imMillionaire.etfunderprice {
			sellinstr(value.price, imMillionaire)
			//---------------------------------------------------
			imMillionaire.sold = true
			//---------------------------------------------------
			imMillionaire.readytoshort = false

			View.ShowInfo(id, fmt.Sprintf("short SELL \nвремя сделки %v\nцена сделки %v\nema %v\n", value.time, value.price, value.ema))
			if debug {
				fmt.Println("short SELL")
				fmt.Println("время сделки ", value.time) //время операции
				fmt.Println("цена сделки", value.price)
				fmt.Println("ema", value.ema) //ema на момент открытие сделки
			}

			imMillionaire.amountOfDeals++
			imMillionaire.tax = imMillionaire.tax + value.price*0.0004
		}

		if value.price > value.ema && imMillionaire.sold { //|| (imMillionaire.sold && value.time.Equal(time.Date(year, mon, day, 19, 45, 0, 0, loc)))
			buyinstr(value.price, imMillionaire)
			//---------------------------------------------------
			imMillionaire.sold = false
			//---------------------------------------------------
			View.ShowInfo(id, fmt.Sprintf("short BUYBACK \nвремя сделки %v\nцена сделки %v\nema %v\n", value.time, value.price, value.ema))

			if debug {
				fmt.Println("short BUYBACK")
				fmt.Println("время сделки ", value.time) //время операции
				fmt.Println("цена сделки", value.price)  //цена закрытия она же сделки
				fmt.Println("ema", value.ema)            //ema на момент закрытия сделки
			}

			imMillionaire.readytoshort = true //false
			//imMillionaire.amountOfDeals++
			imMillionaire.tax = imMillionaire.tax + value.price*0.0004
		}

	} else if imMillionaire.marketgonaclose && imMillionaire.sold {
		buyinstr(value.price, imMillionaire)
		//---------------------------------------------------
		imMillionaire.sold = false
		//---------------------------------------------------
		View.ShowInfo(id, fmt.Sprintf("END DAY CLOSE \nвремя сделки %v\nцена сделки %v\nema %v\n", value.time, value.price, value.ema))

		if debug {
			fmt.Println("Close short position AT THE END OF THE DAY")
			fmt.Println("время сделки ", value.time) //время операции
			fmt.Println("цена сделки", value.price)  //цена закрытия она же сделки
			fmt.Println("ema", value.ema)            //ema на момент закрытия сделки
		}

		imMillionaire.readytoshort = true //false
		//imMillionaire.amountOfDeals++
		imMillionaire.tax = imMillionaire.tax + value.price*0.0004
	}
	//fmt.Println("Ну и времячко ", value.time)
}

//--------------------------------ЛОГИКА ЛОНГ-СТРАТЕГИИ--------------------------------------------------------
func longornot(id int64, value Ema, imMillionaire *gonabeaMillionaire) {
	if !imMillionaire.marketgonaclose {
		if value.price < value.ema && imMillionaire.readytoLong == false {
			imMillionaire.readytoLong = true
		}
		if value.price > value.ema && imMillionaire.readytoLong && !imMillionaire.etfunderprice {
			buyinstr(value.price, imMillionaire)
			//---------------------------------------------------
			imMillionaire.bought = true
			//---------------------------------------------------
			imMillionaire.readytoLong = false

			View.ShowInfo(id, fmt.Sprintf("long BUY \nвремя сделки %v\nцена сделки %v\nema %v\n", value.time, value.price, value.ema))
			if debug {
				fmt.Println("long BUY")
				fmt.Println("время сделки ", value.time) //время операции
				fmt.Println("цена сделки", value.price)
				fmt.Println("ema", value.ema) //ema на момент открытие сделки
			}
			imMillionaire.amountOfDeals++
			imMillionaire.tax = imMillionaire.tax + value.price*0.0004
		}

		if value.price < value.ema && imMillionaire.bought { //todo подумать подходит ли sold маркер для лонг стратегии
			sellinstr(value.price, imMillionaire)
			//---------------------------------------------------
			imMillionaire.bought = false
			//---------------------------------------------------
			View.ShowInfo(id, fmt.Sprintf("long Close position \nвремя сделки %v\nцена сделки %v\nema %v\n", value.time, value.price, value.ema))

			if debug {
				fmt.Println("long Close position")
				fmt.Println("время сделки ", value.time) //время операции
				fmt.Println("цена сделки", value.price)  //цена закрытия она же сделки
				fmt.Println("ema", value.ema)            //ema на момент закрытия сделки
			}

			imMillionaire.readytoLong = true //false
			//imMillionaire.amountOfDeals++
			imMillionaire.tax = imMillionaire.tax + value.price*0.0004
		}

	} else if imMillionaire.marketgonaclose && imMillionaire.bought {
		sellinstr(value.price, imMillionaire)
		//---------------------------------------------------
		imMillionaire.bought = false
		//---------------------------------------------------
		View.ShowInfo(id, fmt.Sprintf("END DAY CLOSE \nвремя сделки %v\nцена сделки %v\nema %v\n", value.time, value.price, value.ema))

		if debug {
			fmt.Println("Close position AT THE END OF THE DAY")
			fmt.Println("время сделки ", value.time) //время операции
			fmt.Println("цена сделки", value.price)  //цена закрытия она же сделки
			fmt.Println("ema", value.ema)            //ema на момент закрытия сделки
		}

		imMillionaire.readytoLong = true //false
		//imMillionaire.amountOfDeals++
		imMillionaire.tax = imMillionaire.tax + value.price*0.0004
	}
}

func sellinstr(price float64, imMillionaire *gonabeaMillionaire) {

	imMillionaire.accountBalance += price
	//imMillionaire.sold = true
}
func buyinstr(price float64, imMillionaire *gonabeaMillionaire) {

	imMillionaire.accountBalance -= price
	//imMillionaire.sold = false
}
