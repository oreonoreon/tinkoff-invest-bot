package trade

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/View"
	"tinkoff-invest-bot/sdk"
)

func convertInterval(interval string) t.CandleInterval {
	switch interval {
	case "1min":
		return t.CandleInterval_CANDLE_INTERVAL_1_MIN
	case "5min":
		return t.CandleInterval_CANDLE_INTERVAL_5_MIN
	case "15min":
		return t.CandleInterval_CANDLE_INTERVAL_15_MIN
	default:
		View.ShowInfo("Не выбран интервал свечей. Бот автоматически выберет интервал равный 15 минутам.")
		time.Sleep(5 * time.Second)
		return t.CandleInterval_CANDLE_INTERVAL_15_MIN
	}
}

func tradingHours(Sk *sdk.Services) ([]*t.TradingSchedule, error) {
	return Sk.InstrumentsService.TradingSchedules("", timestamppb.Now(), timestamppb.New(time.Now().Add(24*time.Hour)))
}

// GTHfaExchange -----------------------GET Trading Hours of all Exchanges-------------------
func GTHfaExchange(Sk *sdk.Services) {
	if tradeHours, err := tradingHours(Sk); err != nil {
		panic(err) //todo заменить панику да и вообще обработай уже нормально ошибки и научись уже пользоваться логированием
	} else {
		for _, value := range tradeHours {

			View.ShowInfo(fmt.Sprintf("Время открытия %v биржы, %v", value.Exchange, value.Days[0].StartTime.AsTime()))
			View.ShowInfo(fmt.Sprintf("Время закрытия %v биржы, %v", value.Exchange, value.Days[0].EndTime.AsTime()))
			View.ShowInfo(fmt.Sprintf("Торговый ли день для %v биржы, %v", value.Exchange, value.Days[0].IsTradingDay))

		}
	}
}
