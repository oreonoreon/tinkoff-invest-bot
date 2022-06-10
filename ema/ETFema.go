package ema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"tinkoff-invest-bot/View"
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
type EtfEma struct {
	ema   float64
	price float64
	time  time.Time
}

func NewEtfEma() *EtfEma {
	return new(EtfEma)
}

var SliceEtfEma []EtfEma

func GetETFema(symbol string, interval string, avarageCoifecent float64, test ...bool) {
	etfEma := NewEtfEma()
	etfcandles, err := getETFcandleFromOutterService(symbol, interval)
	//----------------ОБРАБОТАЕМ ОШИБКУ ПОЛУЧАЕМУЮ ВОВРЕМЯ UNMARSHALL, ОНА ПРОИСХОДИТ ИЗ-ЗА ТОГО ЧТО CLOUDFLARE НАС БЛОЧИТ-------------------------------------------------------------
	if err != nil {
		for n := 0; n < 10; n++ {
			View.ShowInfo("Жду 10 сек и пробую подключиться к сторонему noname-API ещё раз, что бы получить данные о ETF")
			time.Sleep(10 * time.Second)
			etfcandles, err = getETFcandleFromOutterService(symbol, interval)
			if err != nil {
				View.ShowInfo("Неудалось подключиться или нас заблочил стороний noname-API")
				continue
			} else {
				View.ShowInfo("Конекшен установлен с стороним noname-API")
				break
			}
		}
	}
	//---------------------РАСЧИТАЕМ ETF-EMA-------------------------------
	if etfEma.ema, err = strconv.ParseFloat(etfcandles[len(etfcandles)-1].Close, 64); err != nil {
		panic(err)
	}
	a := 2 / (avarageCoifecent + 1)
	for i := len(etfcandles) - 1; i >= 0; i-- {
		//-------расспарсим время со сторонего api так как там UTC-4 и хз какой формат. Приведём его нормальному формату UTC0-------
		if dur, err := time.Parse("2006-01-02 15:04:05", etfcandles[i].Datetime); err != nil {
			panic(err)
		} else {
			etfEma.time = dur.Add(4 * time.Hour)
			//ввыведем время и EMA в это время
			if etfEma.price, err = strconv.ParseFloat(etfcandles[i].Close, 64); err != nil {
				panic(err)
			} else {
				etfEma.ema = a*etfEma.price + (1-a)*etfEma.ema
			}
		}
		//---------------добавим все EtfEma структуры в слайс SLiceEtfEma------------------------------------------------
		SliceEtfEma = append(SliceEtfEma, *etfEma)
		//------------------------РЕШИЛ ИСПОЛЬЗОВАТЬ !МАР! ВМЕСТО СЛАЙСА ETF---------------------
		YaZaebalsya(*etfEma) //TODO если с мапой всё будет работать убрать слайс SliceEtfEma
	}
	//---------------------------------ВЫВЕДЕМ etf-ema ЕСЛИ РЕЖИМ ТЕСТА---------------------------------------------
	if len(test) != 0 {
		for _, value := range SliceEtfEma {
			str := fmt.Sprintf("EMA тикера %v равна %v, при цене закрытия свечи %v.\n Время закртия свечи %v", symbol, value.ema, value.price, value.time)
			View.ShowInfo(str)
		}
	}
}
func getETFcandleFromOutterService(symbol string, interval string) ([]struct {
	Close    string `json:"close"`
	Datetime string `json:"datetime"`
	High     string `json:"high"`
	Low      string `json:"low"`
	Open     string `json:"open"`
	Volume   string `json:"volume"`
}, error) {

	u := url.Values{}
	u.Add("symbol", symbol)     //"SPY"
	u.Add("interval", interval) //"15min"
	u.Add("outputsize", "1000") //1000
	u.Add("format", "json")
	u.Add("apikey", "afe17de2b1154b96b3d4d1d681883d8e")
	url := "https://api.twelvedata.com/" + "time_series?" + u.Encode()
	//url := "https://api.twelvedata.com/ema?symbol=SPY&interval=15min&series_type=close&format=json&outputsize=1000&time_period=20&apikey=afe17de2b1154b96b3d4d1d681883d8e"
	//url := "https://api.twelvedata.com/time_series?symbol=SPY&interval=15min&outputsize=1000&format=json&apikey=afe17de2b1154b96b3d4d1d681883d8e"
	req, err := http.NewRequest("GET", url, nil)

	/*
		req.Header.Add("X-RapidAPI-Host", "twelve-data1.p.rapidapi.com")
		req.Header.Add("X-RapidAPI-Key", "SIGN-UP-FOR-KEY")
	*/
	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	EtfCandles := new(EtfCandle)
	err = json.Unmarshal(body, EtfCandles)
	if err != nil {
		fmt.Println("ВО ВРЕМЯ UNMARSHAL ERROR", err)
		fmt.Println(string(body))
	}
	return EtfCandles.Values, err
}

func NewMapapa() map[time.Time]EtfEma {
	return make(map[time.Time]EtfEma)
}

var Mapapa = NewMapapa()

func YaZaebalsya(etfema EtfEma) {

	if value, ok := Mapapa[etfema.time]; !ok {
		value = etfema
		Mapapa[etfema.time] = value
	}
}
