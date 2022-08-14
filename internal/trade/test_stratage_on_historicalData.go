package trade

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
	"tinkoff-invest-bot/internal/View"
	"tinkoff-invest-bot/internal/ema"
)

func (tink TinkBot) Test_stratage() {
	var symbol, interval, outputsize, figi, end_date string
	var avarageCoifecent float64
	var source string
	stock := make(map[string]string)
	stock["aapl"] = "BBG000B9XRY4"
	stock["yndx"] = "BBG006L8G4H1"

	View.ShowInfo(tink.id, "Выбираем Etf", "Spy", "QQQ")
	symbol = fmt.Sprint(<-tink.ChUser)
	//fmt.Scanln(&symbol, &interval, &avarageCoifecent)

	View.ShowInfo(tink.id, "Выбираем инструмент", "aapl", "yndx")
	stocksymbol := fmt.Sprint(<-tink.ChUser)
	figi = stock[stocksymbol]

	View.ShowInfo(tink.id, "Выбираем интервал свечи", "1min", "5min", "15min", "1h")
	interval = fmt.Sprint(<-tink.ChUser)

	View.ShowInfo(tink.id, "Выбираем ema", "10", "20", "50")
	avarageCoifecent, err := strconv.ParseFloat(fmt.Sprint(<-tink.ChUser), 64)

	View.ShowInfo(tink.id, "Выбираем источник данных", "Tinkoff", "Outsource")
	source = fmt.Sprint(<-tink.ChUser)

	if err != nil {
		fmt.Println(err)
	}
	outputsize = "1000"

	/*
		View.ShowInfo("DВеди желаемый временой интервал или количество свечей. Пример '1000'")
		fmt.Scanln(&outputsize)

	*/
	etfsymbol := symbol

	sheduleCboeSpx, err := tink.Sk.InstrumentsService.TradingSchedules("cboe_spx", timestamppb.Now(), timestamppb.New(time.Now().Add(24*time.Hour)))
	if err != nil {
		panic(err) //todo ну а чё не паника то!)))
	}

	EMA := ema.NewEma()

	switch source {
	case "Tinkoff":
		// ---------------------ПОЛУЧАЕМ ЕМА ИНСТРУМЕНТА С ТИНЬКОФФ----------------------------------------------------
		EMA.EMAHistoric(tink.Sk.Hiscandle(60, figi, convertInterval(interval)), avarageCoifecent)
		// ---------------------ПОЛУЧАЕМ EMA ETFа СО СТОРОНЕГО СЕРВИСА----------------------------------------------------
		Mapapa := ema.ShowEma(symbol, interval, fmt.Sprint(avarageCoifecent), outputsize, end_date)
		// ---------------------!НУ И ВИШЕНКА НА ТОРТЕ - ТЕСТИМ СТРАТЕГИЮ!----------------------------------------------------
		ema.History_testStrategyShort(tink.id, ema.SliceEma, Mapapa, sheduleCboeSpx[0], true)
	case "Outsource":
		etfmap := ema.ShowEma(etfsymbol, interval, fmt.Sprint(avarageCoifecent), outputsize, end_date)
		stockmap := ema.StockSlice(stocksymbol, interval, fmt.Sprint(avarageCoifecent), outputsize, end_date)
		ema.Test_strategy(tink.id, stockmap, etfmap)
	}
	/*
		// ---------------------ПОЛУЧАЕМ ЕМА ИНСТРУМЕНТА С ТИНЬКОФФ----------------------------------------------------
		EMA.EMAHistoric(Sk.Hiscandle(60, figi, convertInterval(interval)), avarageCoifecent)
		// ---------------------ПОЛУЧАЕМ EMA ETFа СО СТОРОНЕГО СЕРВИСА----------------------------------------------------
		//Mapapa := ema.GetETFema(symbol, interval, outputsize, avarageCoifecent)
		Mapapa := ema.ShowEma(symbol, interval, fmt.Sprint(avarageCoifecent), outputsize, end_date)
		// ---------------------!НУ И ВИШЕНКА НА ТОРТЕ - ТЕСТИМ СТРАТЕГИЮ!----------------------------------------------------
		ema.History_testStrategyShort(ema.SliceEma, Mapapa, sheduleCboeSpx[0], true)
		//-----------------------------------------------------------------------------------------------------------------
	*/
	tink.Restart()

}
