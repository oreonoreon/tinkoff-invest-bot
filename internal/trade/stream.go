package trade

import (
	"fmt"
	"log"
	"time"
	"tinkoff-invest-bot/internal/View"
	"tinkoff-invest-bot/internal/convertQuotation"
)

func (tink TinkBot) StreamDataLastPrice() {

	var figi string

	View.ShowInfo(tink.id, "Введите Фиги инструмента", "BBG000BDTBL9", "BBG000B9XRY4", "BBG006L8G4H1")
	//fmt.Scanln(&figi)
	f := <-tink.ChUser
	figi = fmt.Sprint(f)
	//var figi = "BBG000BDTBL9" //"BBG000B9XRY4"

	//--------------------------------ПОДПИСЫВАЕМСЯ НА ПОТОК РЫНОЧНЫХ ДАННЫХ---------------------------------------------
	mds := tink.Sk.MarketDataServiceStream
	err := mds.Send(marketDataStream_SubscribeLastPrice(figi))
	//-------------------------------------------------------------------------------------------------------------------
	if err != nil {
		log.Fatal(err)
	}

	for {
		recv, err := mds.Recv()
		if err != nil {
			log.Println("ошибка во время стрима", err)
			time.Sleep(20 * time.Second)
			break
		}
		//------------------------------последня цена инструмента--------------------------------------------------------
		View.ShowInfo(tink.id, convertQuotation.Convert(recv.GetLastPrice().GetPrice().GetUnits(), recv.GetLastPrice().GetPrice().GetNano()))
		//-----------------------------------------время-----------------------------------------------------------------
		View.ShowInfo(tink.id, recv.GetLastPrice().GetTime().AsTime())

	}

}
