package trade

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/internal/View"
	"tinkoff-invest-bot/internal/ema"
	"tinkoff-invest-bot/internal/sdk"
)

func convertInterval(interval string) t.CandleInterval {
	switch interval {
	case "1min":
		return t.CandleInterval_CANDLE_INTERVAL_1_MIN
	case "5min":
		return t.CandleInterval_CANDLE_INTERVAL_5_MIN
	case "15min":
		return t.CandleInterval_CANDLE_INTERVAL_15_MIN
	case "1h":
		return t.CandleInterval_CANDLE_INTERVAL_HOUR
	default:
		fmt.Println("Не выбран интервал свечей. Бот автоматически выберет интервал равный 15 минутам.")
		time.Sleep(1 * time.Second)
		return t.CandleInterval_CANDLE_INTERVAL_15_MIN
	}
}

func tradingHours(Sk *sdk.Services) ([]*t.TradingSchedule, error) {
	return Sk.InstrumentsService.TradingSchedules("", timestamppb.Now(), timestamppb.New(time.Now().Add(24*time.Hour)))
}

// GTHfaExchange -----------GET Trading Hours of all Exchanges-------------------
func (tink TinkBot) GTHfaExchange() {
	if tradeHours, err := tradingHours(tink.Sk); err != nil {
		panic(err) //todo заменить панику да и вообще обработай уже нормально ошибки и научись уже пользоваться логированием
	} else {
		for _, value := range tradeHours {

			View.ShowInfo(tink.id, fmt.Sprintf("Время открытия %v биржы, %v", value.Exchange, value.Days[0].StartTime.AsTime()))
			View.ShowInfo(tink.id, fmt.Sprintf("Время закрытия %v биржы, %v", value.Exchange, value.Days[0].EndTime.AsTime()))
			View.ShowInfo(tink.id, fmt.Sprintf("Торговый ли день для %v биржы, %v", value.Exchange, value.Days[0].IsTradingDay))
			// хотел распечатать расписание на 3 дня но если это конец недели (к примеру суббота) то в слайсе меньше ячеек чем я хочу вывести
			/*
				View.ShowInfo(fmt.Sprintf("Время открытия %v биржы, сегодня %v, завтра %v, послезавтра %v", value.Exchange, value.Days[0].StartTime.AsTime(), value.Days[1].StartTime.AsTime(), value.Days[2].StartTime.AsTime()))
				View.ShowInfo(fmt.Sprintf("Время закрытия %v биржы, сегодня %v, завтра %v, послезавтра %v", value.Exchange, value.Days[0].EndTime.AsTime(), value.Days[1].EndTime.AsTime(), value.Days[2].EndTime.AsTime()))
				View.ShowInfo(fmt.Sprintf("Торговый ли день для %v биржы, %v", value.Exchange, value.Days[0].IsTradingDay))
			*/
		}
	}
}

func (tink TinkBot) Restart() {
	var answer string
	View.ShowInfo(tink.id, "Do you want to Restart program or Exit?", "Restart", "Exit")
	answer = fmt.Sprint(<-tink.ChUser)
	switch answer {
	case "Exit", "exit":
		//work.Wg.Done()
		return
	default:
		ema.SliceEma = nil //обнулим перед рестартом
		View.ShowInfo(tink.id, "RESTART")
	}

}
func (tink TinkBot) sandBox() {
	sadboxaccount, err := tink.Sk.SandboxService.GetSandboxAccounts()
	if err != nil {
		log.Fatal("sadboxaccount", err)
	}
	accountId := sadboxaccount[0].Id
	sandboxPortfolio, err := tink.Sk.SandboxService.GetSandboxPortfolio(accountId)
	if err != nil {
		log.Fatal("sandboxPortfolio", err)
	}
	if sandboxPortfolio.TotalAmountCurrencies.Units == 0 {
		if _, err = tink.Sk.SandboxService.SandboxPayIn(accountId, &t.MoneyValue{
			Currency: "rub",
			Units:    10000,
			Nano:     0,
		}); err != nil {
			log.Fatalln("SandboxPayIn rub ", err)
		}
		//if _, err = tink.Sk.SandboxService.SandboxPayIn(accountId, &t.MoneyValue{
		//	Currency: "usd",
		//	Units:    1000,
		//	Nano:     50,
		//}); err != nil {
		//	log.Fatalln("SandboxPayIn usd", err)
		//}
	}
	View.ShowInfo(tink.id, sandboxPortfolio)
}
func (tink TinkBot) AccSandBox() {
	account, err := tink.Sk.SandboxService.GetSandboxAccounts()
	if err != nil {
		View.ShowInfo(tink.id, err)
	}
	portfolio, err := tink.Sk.SandboxService.GetSandboxPortfolio(account[0].Id)
	if err != nil {
		View.ShowInfo(tink.id, err)
	}
	View.ShowInfo(tink.id, portfolio)
}

func marketDataStream_SubscribeLastPrice(figi string) *t.MarketDataRequest {
	return &t.MarketDataRequest{
		Payload: &t.MarketDataRequest_SubscribeLastPriceRequest{
			SubscribeLastPriceRequest: &t.SubscribeLastPriceRequest{
				SubscriptionAction: t.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments: []*t.LastPriceInstrument{
					{
						Figi: figi, //"BBG000B9XRY4", //"BBG006L8G4H1",
					},
				},
			},
		},
	}
}

func marketDataStream_SubscribeCandles(figi string, subscriptionInterval t.SubscriptionInterval) *t.MarketDataRequest {
	return &t.MarketDataRequest{
		Payload: &t.MarketDataRequest_SubscribeCandlesRequest{
			SubscribeCandlesRequest: &t.SubscribeCandlesRequest{
				SubscriptionAction: t.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments: []*t.CandleInstrument{
					{
						Figi:     figi,                 // "BBG006L8G4H1", // "BBG000B9XRY4", "BBG000BDTBL9"
						Interval: subscriptionInterval, //t.SubscriptionInterval_SUBSCRIPTION_INTERVAL_FIVE_MINUTES,
					},
				},
			},
		},
	}
}

func marketDataStream_SubscribeInfo(figi string) *t.MarketDataRequest {
	return &t.MarketDataRequest{
		Payload: &t.MarketDataRequest_SubscribeInfoRequest{
			SubscribeInfoRequest: &t.SubscribeInfoRequest{
				SubscriptionAction: t.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments: []*t.InfoInstrument{
					{
						Figi: figi, //"BBG000B9XRY4",
					},
				},
			},
		},
	}
}
func Сonvert_Symbol_To_Figi(symbol string) string {
	return symbol //todo написать ковертор символа в фиги
}
