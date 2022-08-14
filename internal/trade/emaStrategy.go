package trade

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/internal/loggy"
)

func (tink TinkBot) sellinstr() {

	saccount, err := tink.Sk.SandboxService.GetSandboxAccounts()
	if err != nil {
		log.Fatalln("GetSandboxAccounts error", err)
	}
	postorderResp, err := tink.Sk.SandboxService.PostSandboxOrder(&t.PostOrderRequest{
		Figi:      "BBG006L8G4H1", //TODO всавить уже функцию выбора инструмента ато задрало вставляти фиги Яндекса
		Quantity:  1,
		Direction: t.OrderDirection_ORDER_DIRECTION_SELL,
		AccountId: saccount[0].Id,
		OrderType: t.OrderType_ORDER_TYPE_MARKET,
		OrderId:   uuid.New().String(),
	})
	if err != nil {
		loggy.GetLogger().Sugar().Warn("postorderResp ", err)
	}
	fmt.Println(postorderResp)
	//accountBalance += price
	//sold = true
}
func (tink TinkBot) buyinstr() {

	saccount, err := tink.Sk.SandboxService.GetSandboxAccounts()
	if err != nil {
		log.Fatalf("GetSandboxAccounts", err)
	}
	postorderResp, err := tink.Sk.SandboxService.PostSandboxOrder(&t.PostOrderRequest{
		Figi:      "BBG006L8G4H1", //TODO всавить уже функцию выбора инструмента ато задрало вставляти фиги Яндекса
		Quantity:  1,
		Direction: t.OrderDirection_ORDER_DIRECTION_BUY,
		AccountId: saccount[0].Id,
		OrderType: t.OrderType_ORDER_TYPE_MARKET,
		OrderId:   uuid.New().String(),
	})
	if err != nil {
		loggy.GetLogger().Sugar().Warn("postorderResp ", err)
	}
	fmt.Println(postorderResp)
	fmt.Println(tink.Sk.SandboxService.GetSandboxOrders(saccount[0].Id))
	//accountBalance -= price
	//sold = false
}
