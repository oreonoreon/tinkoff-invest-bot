package trade

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"tinkoff-invest-bot/View"
	"tinkoff-invest-bot/ema"
	"tinkoff-invest-bot/sdk"
)

func Test_stratage() {
	var symbol, interval string
	var avarageCoifecent float64
	View.ShowInfo("SPY 15min 20")
	fmt.Scanln(&symbol, &interval, &avarageCoifecent)
	Sk = sdk.NewServices()
	sheduleCboeSpx, err := Sk.InstrumentsService.TradingSchedules("cboe_spx", timestamppb.Now(), timestamppb.New(time.Now().Add(24*time.Hour)))
	if err != nil {
		panic(err) //todo ну а чё не паника то!)))
	}

	var figi = "BBG000B9XRY4" // фиги AAPL
	EMA := ema.NewEma()       //View.ShowInfo(*EMA)показал что после restart() всё вроде коректно обнуляется
	// ---------------------ПОЛУЧАЕМ ЕМА ИНСТРУМЕНТА С ТИНЬКОФФ----------------------------------------------------
	EMA.EMAHistoric(Sk.Hiscandle(30, figi, convertInterval(interval)), avarageCoifecent)
	// ---------------------ПОЛУЧАЕМ EMA ETFа СО СТОРОНЕГО СЕРВИСА----------------------------------------------------
	ema.GetETFema(symbol, interval, avarageCoifecent)
	// ---------------------!НУ И ВИШЕНКА НА ТОРТЕ - ТЕСТИМ СТРАТЕГИЮ!----------------------------------------------------
	ema.History_testStrategyShort(ema.SliceEma, ema.Mapapa, sheduleCboeSpx[0], true)
	//-----------------------------------------------------------------------------------------------------------------
	restart()
	//restart(Test_stratage)
}
