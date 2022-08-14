package trade

import (
	"fmt"
	"log"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/internal/View"
	"tinkoff-invest-bot/internal/ema"
	"tinkoff-invest-bot/internal/sdk"
)

type TinkBot struct {
	id     int64
	Sk     *sdk.Services
	ChUser chan any
}

func NewTinkBot(id int64, ChUser chan any) TinkBot {
	return TinkBot{id: id, Sk: sdk.NewServices(), ChUser: ChUser}
}

func (tink TinkBot) StartTrade(figi, symbol, interval string, avarageCoifecent float64) {
	figiEtf := Сonvert_Symbol_To_Figi(symbol) //todo написать ковертор символа в фиги

	EMA := ema.NewEma()
	// ---------------------ПОЛУЧАЕМ ЕМА ИНСТРУМЕНТА С ТИНЬКОФФ----------------------------------------------------
	EMA.EMAHistoric(tink.Sk.Hiscandle(60, figi, convertInterval(interval)), avarageCoifecent)
	// ---------------------ПОЛУЧАЕМ EMA ETFа СО СТОРОНЕГО СЕРВИСА----------------------------------------------------
	symbol = "spy" //todo написать ковертор символа в фиги
	Mapapa := ema.ShowEma(symbol, interval, fmt.Sprint(avarageCoifecent), "1", "")
	emaEtf, priceEtf := ema.Mapapa_Ema_Price(Mapapa)
	EMAETF := ema.NewEma()
	EMAETF.Setema(emaEtf)
	EMAETF.Setprice(priceEtf)
	tink.sandBox()
	View.ShowInfo(tink.id, EMAETF)
	View.ShowInfo(tink.id, EMA)
	//View.ShowInfo(id, bot.BOT,Mapapa)
	View.ShowInfo(tink.id, emaEtf)

	//--------------------------------ПОДПИСЫВАЕМСЯ НА ПОТОК РЫНОЧНЫХ ДАННЫХ---------------------------------------------
	mds := tink.Sk.MarketDataServiceStream
	if err := mds.Send(marketDataStream_SubscribeCandles(figi, t.SubscriptionInterval_SUBSCRIPTION_INTERVAL_ONE_MINUTE)); err != nil {
		log.Fatalln(err)
	}
	if err := mds.Send(marketDataStream_SubscribeLastPrice(figiEtf)); err != nil {
		log.Fatalln(err)
	}
	//-------------------------------------------------------------------------------------------------------------------

	for {
		recv, err := mds.Recv()
		if err != nil {
			log.Println("ошибка во время стрима", err)
			time.Sleep(3 * time.Second) //todo решить проблему востановления связи с сервером
			break
		}
		fmt.Println(recv.GetCandle().GetTime().AsTime())
		if timeIntervalCheck(recv.GetCandle().GetTime().AsTime(), interval) { //todo написать функцию времени проверяющую совпадение времени обезличеной сделки переданой стримом и выбраным интервалом
			EMA.EMAstream(recv.GetCandle(), avarageCoifecent)
			EMAETF.EMAstreamETF(recv.GetLastPrice(), avarageCoifecent)
			ema.StreamTrade_strategy(tink.id, *EMA, Mapapa)
		}

	}

}
func timeIntervalCheck(time2 time.Time, interval string) bool {
	var i time.Duration
	switch interval {
	case "1min":
		i = time.Minute
	case "5min":
		i = 5 * time.Minute
	case "15min":
		i = 15 * time.Minute
	case "1h":
		i = time.Hour
	default:
		fmt.Println("Интервал свечей не выбран")
	}

	if time2.Minute()%int(i.Minutes()) == 0 {
		return true
	} else {
		return false
	}
}
