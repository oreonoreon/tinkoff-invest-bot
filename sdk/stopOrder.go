package sdk

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/loggy"
	"tinkoff-invest-bot/metrics"
)

type StopOrdersInterface interface {
	// The method of placing a stop order.
	PostStopOrder(stopOrder *t.PostStopOrderRequest) (string, error)
	// Method for getting a list of active stop orders on the account.
	GetStopOrders(accountID string) ([]*t.StopOrder, error)
	// The method of canceling the stop order.
	CancelStopOrder(accountID string, stopOrderID string) (*timestamp.Timestamp, error)
}

type StopOrdersService struct {
	client t.StopOrdersServiceClient
}

func NewStopOrdersService() *StopOrdersService {
	conn, err := Connection()
	if err != nil {
		loggy.GetLogger().Sugar().Fatal(err.Error())
	}

	client := t.NewStopOrdersServiceClient(conn)
	return &StopOrdersService{client: client}
}

func (sos StopOrdersService) PostStopOrder(stopOrder *t.PostStopOrderRequest) (string, error) {
	ctx, cancel := NewContext()
	defer cancel()

	sos.incrementRequestsCounter("PostStopOrder")
	res, err := sos.client.PostStopOrder(ctx, stopOrder)
	if err != nil {
		sos.incrementApiCallErrors("PostStopOrder", err.Error())
		return "", err
	}

	return res.StopOrderId, nil
}

func (sos StopOrdersService) GetStopOrders(accountID string) ([]*t.StopOrder, error) {
	ctx, cancel := NewContext()
	defer cancel()

	sos.incrementRequestsCounter("GetStopOrders")
	res, err := sos.client.GetStopOrders(ctx, &t.GetStopOrdersRequest{
		AccountId: accountID,
	})
	if err != nil {
		sos.incrementApiCallErrors("GetStopOrders", err.Error())
		return nil, err
	}

	return res.StopOrders, nil
}

func (sos StopOrdersService) CancelStopOrder(accountID string, stopOrderID string) (*timestamp.Timestamp, error) {
	ctx, cancel := NewContext()
	defer cancel()

	sos.incrementRequestsCounter("CancelStopOrder")
	res, err := sos.client.CancelStopOrder(ctx, &t.CancelStopOrderRequest{
		AccountId:   accountID,
		StopOrderId: stopOrderID,
	})
	if err != nil {
		sos.incrementApiCallErrors("CancelStopOrder", err.Error())
		return nil, err
	}

	return res.Time, nil
}

func (sos StopOrdersService) incrementRequestsCounter(method string) {
	metrics.ApiRequests.WithLabelValues(loggy.GetBotID(), "StopOrdersService", method).Inc()
}

func (sos StopOrdersService) incrementApiCallErrors(method string, error string) {
	metrics.ApiCallErrors.WithLabelValues(loggy.GetBotID(), "StopOrdersService", method, error).Inc()
}
