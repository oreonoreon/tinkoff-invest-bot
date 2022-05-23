package sdk

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/connect"
	"tinkoff-invest-bot/loggy"
	"tinkoff-invest-bot/metrics"
)

type OrdersInterface interface {
	// The method of submitting the order.
	PostOrder(order *t.PostOrderRequest) (*t.PostOrderResponse, error)
	// The method of cancellation of the trade order.
	CancelOrder(accountID string, orderID string) (*timestamp.Timestamp, error)
	// The method of obtaining the status of a trade order.
	GetOrderState(accountID string, orderID string) (*t.OrderState, error)
	// The method of getting a list of active orders for the account.
	GetOrders(accountID string) ([]*t.OrderState, error)
}

type OrdersService struct {
	client t.OrdersServiceClient
}

func NewOrdersService() *OrdersService {
	conn, err := connect.Connection()
	if err != nil {
		loggy.GetLogger().Sugar().Fatal(err.Error())
	}

	client := t.NewOrdersServiceClient(conn)
	return &OrdersService{client: client}
}

func (os OrdersService) PostOrder(order *t.PostOrderRequest) (*t.PostOrderResponse, error) {
	ctx, cancel := connect.NewContext()
	defer cancel()

	os.incrementRequestsCounter("PostOrder")
	res, err := os.client.PostOrder(ctx, order)
	if err != nil {
		os.incrementApiCallErrors("PostOrder", err.Error())
		return nil, err
	}

	return res, nil
}

func (os OrdersService) CancelOrder(accountID string, orderID string) (*timestamp.Timestamp, error) {
	ctx, cancel := connect.NewContext()
	defer cancel()

	os.incrementRequestsCounter("CancelOrder")
	res, err := os.client.CancelOrder(ctx, &t.CancelOrderRequest{
		AccountId: accountID,
		OrderId:   orderID,
	})
	if err != nil {
		os.incrementApiCallErrors("CancelOrder", err.Error())
		return nil, err
	}

	return res.Time, nil
}

func (os OrdersService) GetOrderState(accountID string, orderID string) (*t.OrderState, error) {
	ctx, cancel := connect.NewContext()
	defer cancel()

	os.incrementRequestsCounter("GetOrderState")
	res, err := os.client.GetOrderState(ctx, &t.GetOrderStateRequest{
		AccountId: accountID,
		OrderId:   orderID,
	})
	if err != nil {
		os.incrementApiCallErrors("GetOrderState", err.Error())
		return nil, err
	}

	return res, nil
}

func (os OrdersService) GetOrders(accountID string) ([]*t.OrderState, error) {
	ctx, cancel := connect.NewContext()
	defer cancel()

	os.incrementRequestsCounter("GetOrders")
	res, err := os.client.GetOrders(ctx, &t.GetOrdersRequest{
		AccountId: accountID,
	})
	if err != nil {
		os.incrementApiCallErrors("GetOrders", err.Error())
		return nil, err
	}

	return res.Orders, nil
}

func (os OrdersService) incrementRequestsCounter(method string) {
	metrics.ApiRequests.WithLabelValues(loggy.GetBotID(), "OrdersService", method).Inc()
}

func (os OrdersService) incrementApiCallErrors(method string, error string) {
	metrics.ApiCallErrors.WithLabelValues(loggy.GetBotID(), "OrdersService", method, error).Inc()
}
