package trade

import (
	"fmt"
	"tinkoff-invest-bot/internal/View"
	ema "tinkoff-invest-bot/internal/ema"
)

func (tink TinkBot) Test_S_With_Data_from_outerApi() {
	var etfsymbol, stocksymbol, interval, outputsize, end_date string
	var avarageCoifecent string //не float64 потомучто собираю url через url.Values а там параметры строки и не охото ковертить float64 в string

	//-----------------------------------------------

	View.ShowInfo(tink.id, "Выбираем Etf", "Spy", "QQQ")
	etfsymbol = fmt.Sprint(<-tink.ChUser)
	//fmt.Scanln(&symbol, &interval, &avarageCoifecent)

	View.ShowInfo(tink.id, "Выбираем инструмент", "aapl", "yndx")
	stocksymbol = fmt.Sprint(<-tink.ChUser)

	View.ShowInfo(tink.id, "Выбираем интервал свечи", "1min", "5min", "15min", "1h")
	interval = fmt.Sprint(<-tink.ChUser)

	View.ShowInfo(tink.id, "Выбираем ema", "10", "20", "50")

	avarageCoifecent = fmt.Sprint(<-tink.ChUser)

	outputsize = "1000"
	//------------------------------------
	etfmap := ema.ShowEma(etfsymbol, interval, avarageCoifecent, outputsize, end_date)
	//stockmap := ema.ShowEma(stocksymbol, interval, avarageCoifecent, outputsize)
	stockmap := ema.StockSlice(stocksymbol, interval, avarageCoifecent, outputsize, end_date)
	ema.Test_strategy(tink.id, stockmap, etfmap)
	//ema.History_testStrategyShort(stockmap, etfmap, nil, true)
	tink.Restart()
}
