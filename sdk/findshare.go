package sdk

import (
	"fmt"
	"log"
	t "tinkoff-invest-bot/Tinkoff/investapi"
)

func (sk *Services) findS(tiker string) string {
	var shareFigi string
	shares, err := sk.InstrumentsService.Shares(t.InstrumentStatus_INSTRUMENT_STATUS_BASE)
	if err != nil {
		log.Println(err)
		return ""
	}
	for _, share := range shares {
		if share.Ticker == tiker {
			shareFigi = share.Figi
			break
		}
	}
	fmt.Println(shareFigi)
	return shareFigi
}
