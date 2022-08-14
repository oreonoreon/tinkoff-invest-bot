package ema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Str struct {
	Meta struct {
		Currency         string `json:"currency"`
		Exchange         string `json:"exchange"`
		ExchangeTimezone string `json:"exchange_timezone"`
		Indicator        struct {
			Name       string `json:"name"`
			SeriesType string `json:"series_type"`
			TimePeriod int    `json:"time_period"`
		} `json:"indicator"`
		Interval string `json:"interval"`
		Symbol   string `json:"symbol"`
		Type     string `json:"type"`
	} `json:"meta"`
	Status string `json:"status"`
	Values []struct {
		Datetime string `json:"datetime"`
		Ema      string `json:"ema"`
	} `json:"values"`
}
type EtfCandle struct {
	Meta struct {
		Currency         string `json:"currency"`
		Exchange         string `json:"exchange"`
		ExchangeTimezone string `json:"exchange_timezone"`
		Interval         string `json:"interval"`
		Symbol           string `json:"symbol"`
		Type             string `json:"type"`
	} `json:"meta"`
	Status string `json:"status"`
	Values []struct {
		Close    string `json:"close"`
		Datetime string `json:"datetime"`
		High     string `json:"high"`
		Low      string `json:"low"`
		Open     string `json:"open"`
		Volume   string `json:"volume"`
	} `json:"values"`
}

type Price_Ema struct {
	price float64
	ema   float64
}

//-------расспарсим время со сторонего api так как там UTC-4 и хз какой формат. Приведём его к нормальному формату UTC-0-------
func parseTime(t string) time.Time {
	if dur, err := time.Parse("2006-01-02 15:04:05", t); err != nil { // t=etfcandles.Values[i].Datetime
		panic(err)
	} else {
		return dur.Add(4 * time.Hour)
		//ввыведем время и EMA в это время

	}
}
func parsePrice(pr string) float64 {
	price, err := strconv.ParseFloat(pr, 64) //pr=etfcandles.Values[i].Close
	if err != nil {
		panic(err)
	}
	return price
}

//-----------------------------------------------------------------------------------------
func getStockCandleFromOutterService(symbol, interval, outputsize, end_date string) (*EtfCandle, error) {

	u := url.Values{}
	u.Add("symbol", symbol)         //"SPY"
	u.Add("interval", interval)     //"15min"
	u.Add("outputsize", outputsize) //1000
	u.Add("format", "json")
	u.Add("apikey", "afe17de2b1154b96b3d4d1d681883d8e")
	if end_date != "" {
		u.Add("end_date", end_date)
	}
	url := "https://api.twelvedata.com/" + "time_series?" + u.Encode()

	//url := "https://api.twelvedata.com/ema?symbol=SPY&interval=15min&series_type=close&format=json&outputsize=1000&time_period=20&apikey=afe17de2b1154b96b3d4d1d681883d8e"
	//url := "https://api.twelvedata.com/time_series?symbol=SPY&interval=15min&outputsize=1000&format=json&apikey=afe17de2b1154b96b3d4d1d681883d8e"
	req, err := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	EtfCandles := new(EtfCandle)
	err = json.Unmarshal(body, EtfCandles)
	if err != nil {
		fmt.Println("ВО ВРЕМЯ UNMARSHAL ERROR", err)
		fmt.Println(string(body))
	}
	//return EtfCandles.Values, err
	return EtfCandles, err
}

func emaFrom12API(symbol, interval, coifecent, outputsize, end_date string) (*Str, error) {

	u := url.Values{}
	u.Add("symbol", symbol)         //"SPY"
	u.Add("interval", interval)     //"15min"
	u.Add("outputsize", outputsize) //1000
	u.Add("end_date", end_date)
	u.Add("format", "json")
	u.Add("apikey", "afe17de2b1154b96b3d4d1d681883d8e")
	u.Add("time_period", coifecent)
	url := "https://api.twelvedata.com/" + "ema?" + u.Encode()

	//url := "https://api.twelvedata.com/ema?symbol=SPY&interval=15min&series_type=close&format=json&outputsize=10&time_period=20&apikey=afe17de2b1154b96b3d4d1d681883d8e"
	//url := "https://api.twelvedata.com/time_series?symbol=SPY,AAPL&interval=15min&outputsize=10&format=json&apikey=afe17de2b1154b96b3d4d1d681883d8e"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("request error", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("response error", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("read body error", err)
	}
	emas := new(Str)
	err = json.Unmarshal(body, emas)
	if err != nil {
		fmt.Println("ВО ВРЕМЯ UNMARSHAL ERROR", err)
		fmt.Println(string(body))
	}
	return emas, err
}

func ShowEma(symbol, interval, coifecent, outputsize, end_date string) map[time.Time]Price_Ema {
	ema, err := emaFrom12API(symbol, interval, coifecent, outputsize, end_date)
	//----------------ОБРАБОТАЕМ ОШИБКУ ПОЛУЧАЕМУЮ ВОВРЕМЯ UNMARSHALL, ОНА ПРОИСХОДИТ ИЗ-ЗА ТОГО ЧТО CLOUDFLARE НАС БЛОЧИТ-------------------------------------------------------------
	if err != nil {
		log.Println(err)
		for n := 0; n <= 10; n++ {
			fmt.Println("Жду 10 сек и пробую подключиться к сторонему noname-API ещё раз, что бы получить данные о ETF")
			time.Sleep(10 * time.Second)
			ema, err = emaFrom12API(symbol, interval, coifecent, outputsize, end_date)
			if err != nil {
				fmt.Println("Неудалось подключиться или нас заблочил стороний noname-API")
				if n == 10 {
					log.Fatalln("Неудалось подключиться или нас заблочил стороний noname-API", err)
				}
				continue
			} else {
				fmt.Println("Конекшен установлен с стороним noname-API")
				break
			}
		}
	}
	//-------------------------------------------------------------------------------------------------------------------
	stock, err := getStockCandleFromOutterService(symbol, interval, outputsize, end_date)
	//----------------ОБРАБОТАЕМ ОШИБКУ ПОЛУЧАЕМУЮ ВОВРЕМЯ UNMARSHALL, ОНА ПРОИСХОДИТ ИЗ-ЗА ТОГО ЧТО CLOUDFLARE НАС БЛОЧИТ-------------------------------------------------------------
	if err != nil {
		log.Println(err)
		for n := 0; n <= 10; n++ {
			fmt.Println("Жду 10 сек и пробую подключиться к сторонему noname-API ещё раз, что бы получить данные цены Инструмента")
			time.Sleep(10 * time.Second)
			stock, err = getStockCandleFromOutterService(symbol, interval, outputsize, end_date)
			if err != nil {
				fmt.Println("Неудалось подключиться или нас заблочил стороний noname-API")
				if n == 10 {
					log.Fatalln("Неудалось подключиться или нас заблочил стороний noname-API", err)
				}
				continue
			} else {
				fmt.Println("Конекшен установлен с стороним noname-API")
				break
			}
		}
	}

	timeMap := make(map[time.Time]Price_Ema)
	for i := 0; i < len(stock.Values); i++ {
		if stock.Values[i].Datetime == ema.Values[i].Datetime {
			newPriceEma := Price_Ema{parsePrice(stock.Values[i].Close), parsePrice(ema.Values[i].Ema)}
			t := parseTime(stock.Values[i].Datetime)
			if _, ok := timeMap[t]; !ok { //t=stock.Values[i].Datetime
				timeMap[t] = newPriceEma //t=stock.Values[i].Datetime
			}
		} else {
			log.Fatalln("При попытке соединить цену и ema инструмента или etf-а, интервалы не совпали в слайсе не совпали")
		}

	}
	return timeMap
}

// Mapapa_Ema_Price Функция что бы получить ema и price из map[time.Time]Price_Ema, но с учётом только если там одно значение (когда outputsize = 1)
func Mapapa_Ema_Price(timemap map[time.Time]Price_Ema) (float64, float64) {
	for _, v := range timemap {
		return v.ema, v.price
	}
	return 0, 0
}

//---------------------------------------------------------------------------

func StockSlice(symbol, interval, coefficient, outputsize, end_date string) []Ema {
	//ema:=NewEma()
	stock, err := getStockCandleFromOutterService(symbol, interval, outputsize, end_date)
	//----------------ОБРАБОТАЕМ ОШИБКУ ПОЛУЧАЕМУЮ ВОВРЕМЯ UNMARSHALL, ОНА ПРОИСХОДИТ ИЗ-ЗА ТОГО ЧТО CLOUDFLARE НАС БЛОЧИТ-------------------------------------------------------------
	if err != nil {
		log.Println(err)
		for n := 0; n <= 10; n++ {
			fmt.Println("Жду 10 сек и пробую подключиться к сторонему noname-API ещё раз, что бы получить данные цены Инструмента")
			time.Sleep(10 * time.Second)
			stock, err = getStockCandleFromOutterService(symbol, interval, outputsize, end_date)
			if err != nil {
				fmt.Println("Неудалось подключиться или нас заблочил стороний noname-API")
				if n == 10 {
					log.Fatalln("Неудалось подключиться или нас заблочил стороний noname-API", err)
				}
				continue
			} else {
				fmt.Println("Конекшен установлен с стороним noname-API")
				break
			}
		}
	}
	ema, err := emaFrom12API(symbol, interval, coefficient, outputsize, end_date)
	//----------------ОБРАБОТАЕМ ОШИБКУ ПОЛУЧАЕМУЮ ВОВРЕМЯ UNMARSHALL, ОНА ПРОИСХОДИТ ИЗ-ЗА ТОГО ЧТО CLOUDFLARE НАС БЛОЧИТ-------------------------------------------------------------
	if err != nil {
		log.Println(err)
		for n := 0; n <= 10; n++ {
			fmt.Println("Жду 10 сек и пробую подключиться к сторонему noname-API ещё раз, что бы получить данные о ETF")
			time.Sleep(10 * time.Second)
			ema, err = emaFrom12API(symbol, interval, coefficient, outputsize, end_date)
			if err != nil {
				fmt.Println("Неудалось подключиться или нас заблочил стороний noname-API")
				if n == 10 {
					log.Fatalln("Неудалось подключиться или нас заблочил стороний noname-API", err)
				}
				continue
			} else {
				fmt.Println("Конекшен установлен с стороним noname-API")
				break
			}
		}
	}

	stockSlice := make([]Ema, 0)
	for i := len(stock.Values) - 1; i >= 0; i-- {
		stockSlice = append(stockSlice, Ema{parsePrice(ema.Values[i].Ema), parsePrice(stock.Values[i].Close), parseTime(stock.Values[i].Datetime)})
		fmt.Println(parseTime(stock.Values[i].Datetime))
	}
	return stockSlice
}
