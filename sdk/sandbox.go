package sdk

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/loggy"
	"tinkoff-invest-bot/metrics"
)

type SandboxInterface interface {
	// The method of registering an account in the sandbox.
	OpenSandboxAccount() (string, error)
	// The method of getting accounts in the sandbox.
	GetSandboxAccounts() ([]*t.Account, error)
	// The method of closing an account in the sandbox.
	CloseSandboxAccount(accountID string) error
	// The method of placing a trade order in the sandbox.
	PostSandboxOrder(order *t.PostOrderRequest) (*t.PostOrderResponse, error)
	// Method for getting a list of active applications for an account in the sandbox.
	GetSandboxOrders(accountID string) ([]*t.OrderState, error)
	// Method for getting a list of active orders for an account in the sandbox.
	CancelSandboxOrder(accountID string, orderID string) (*timestamp.Timestamp, error)
	// The method of obtaining the order status in the sandbox.
	GetSandboxOrderState(accountID string, orderID string) (*t.OrderState, error)
	// The method of obtaining positions on the virtual sandbox account.
	GetSandboxPositions(accountID string) (*t.PositionsResponse, error)
	// The method of receiving operations in the sandbox by account number.
	GetSandboxOperations(filter *t.OperationsRequest) ([]*t.Operation, error)
	// The method of getting a portfolio in the sandbox.
	GetSandboxPortfolio(accountID string) (*t.PortfolioResponse, error)
	// The method of depositing funds in the sandbox.
	SandboxPayIn(accountID string, amount *t.MoneyValue) (*t.MoneyValue, error)
}

type SandboxService struct {
	client t.SandboxServiceClient
}

func NewSandboxService() *SandboxService {
	conn, err := Connection()
	if err != nil {
		loggy.GetLogger().Sugar().Fatal(err.Error())
	}

	client := t.NewSandboxServiceClient(conn)
	return &SandboxService{client: client}
}

func (ss SandboxService) OpenSandboxAccount() (string, error) {
	ctx, cancel := NewContext()
	defer cancel()

	ss.incrementRequestsCounter("OpenSandboxAccount")
	res, err := ss.client.OpenSandboxAccount(ctx, &t.OpenSandboxAccountRequest{})
	if err != nil {
		ss.incrementApiCallErrors("OpenSandboxAccount", err.Error())
		return "", err
	}

	return res.AccountId, nil
}

func (ss SandboxService) GetSandboxAccounts() ([]*t.Account, error) {
	ctx, cancel := NewContext()
	defer cancel()

	ss.incrementRequestsCounter("GetSandboxAccounts")
	res, err := ss.client.GetSandboxAccounts(ctx, &t.GetAccountsRequest{})
	if err != nil {
		ss.incrementApiCallErrors("GetSandboxAccounts", err.Error())
		return nil, err
	}

	return res.Accounts, nil
}

func (ss SandboxService) CloseSandboxAccount(accountID string) error {
	ctx, cancel := NewContext()
	defer cancel()

	ss.incrementRequestsCounter("CloseSandboxAccount")
	_, err := ss.client.CloseSandboxAccount(ctx, &t.CloseSandboxAccountRequest{
		AccountId: accountID,
	})
	if err != nil {
		ss.incrementApiCallErrors("CloseSandboxAccount", err.Error())
		return err
	}

	return nil
}

func (ss SandboxService) PostSandboxOrder(order *t.PostOrderRequest) (*t.PostOrderResponse, error) {
	ctx, cancel := NewContext()
	defer cancel()

	ss.incrementRequestsCounter("PostSandboxOrder")
	res, err := ss.client.PostSandboxOrder(ctx, order)
	if err != nil {
		ss.incrementApiCallErrors("PostSandboxOrder", err.Error())
		return nil, err
	}

	return res, nil
}

func (ss SandboxService) GetSandboxOrders(accountID string) ([]*t.OrderState, error) {
	ctx, cancel := NewContext()
	defer cancel()

	ss.incrementRequestsCounter("GetSandboxOrders")
	res, err := ss.client.GetSandboxOrders(ctx, &t.GetOrdersRequest{
		AccountId: accountID,
	})
	if err != nil {
		ss.incrementApiCallErrors("GetSandboxOrders", err.Error())
		return nil, err
	}

	return res.Orders, nil
}

func (ss SandboxService) CancelSandboxOrder(accountID string, orderID string) (*timestamp.Timestamp, error) {
	ctx, cancel := NewContext()
	defer cancel()

	ss.incrementRequestsCounter("CancelSandboxOrder")
	res, err := ss.client.CancelSandboxOrder(ctx, &t.CancelOrderRequest{
		AccountId: accountID,
		OrderId:   orderID,
	})
	if err != nil {
		ss.incrementApiCallErrors("CancelSandboxOrder", err.Error())
		return nil, err
	}

	return res.Time, nil
}

func (ss SandboxService) GetSandboxOrderState(accountID string, orderID string) (*t.OrderState, error) {
	ctx, cancel := NewContext()
	defer cancel()

	ss.incrementRequestsCounter("GetSandboxOrderState")
	res, err := ss.client.GetSandboxOrderState(ctx, &t.GetOrderStateRequest{
		AccountId: accountID,
		OrderId:   orderID,
	})
	if err != nil {
		ss.incrementApiCallErrors("GetSandboxOrderState", err.Error())
		return nil, err
	}

	return res, nil
}

func (ss SandboxService) GetSandboxPositions(accountID string) (*t.PositionsResponse, error) {
	ctx, cancel := NewContext()
	defer cancel()

	ss.incrementRequestsCounter("GetSandboxPositions")
	res, err := ss.client.GetSandboxPositions(ctx, &t.PositionsRequest{
		AccountId: accountID,
	})
	if err != nil {
		ss.incrementApiCallErrors("GetSandboxPositions", err.Error())
		return nil, err
	}

	return res, nil
}

func (ss SandboxService) GetSandboxOperations(filter *t.OperationsRequest) ([]*t.Operation, error) {
	ctx, cancel := NewContext()
	defer cancel()

	ss.incrementRequestsCounter("GetSandboxOperations")
	res, err := ss.client.GetSandboxOperations(ctx, filter)
	if err != nil {
		ss.incrementApiCallErrors("GetSandboxOperations", err.Error())
		return nil, err
	}

	return res.Operations, nil
}

func (ss SandboxService) GetSandboxPortfolio(accountID string) (*t.PortfolioResponse, error) {
	ctx, cancel := NewContext()
	defer cancel()

	ss.incrementRequestsCounter("GetSandboxPortfolio")
	res, err := ss.client.GetSandboxPortfolio(ctx, &t.PortfolioRequest{
		AccountId: accountID,
	})
	if err != nil {
		ss.incrementApiCallErrors("GetSandboxPortfolio", err.Error())
		return nil, err
	}

	return res, nil
}

func (ss SandboxService) SandboxPayIn(accountID string, amount *t.MoneyValue) (*t.MoneyValue, error) {
	ctx, cancel := NewContext()
	defer cancel()

	ss.incrementRequestsCounter("SandboxPayIn")
	res, err := ss.client.SandboxPayIn(ctx, &t.SandboxPayInRequest{
		AccountId: accountID,
		Amount:    amount,
	})
	if err != nil {
		ss.incrementApiCallErrors("SandboxPayIn", err.Error())
		return nil, err
	}

	return res.Balance, nil
}

func (ss SandboxService) incrementRequestsCounter(method string) {
	metrics.ApiRequests.WithLabelValues(loggy.GetBotID(), "SandboxService", method).Inc()
}

func (ss SandboxService) incrementApiCallErrors(method string, error string) {
	metrics.ApiCallErrors.WithLabelValues(loggy.GetBotID(), "SandboxService", method, error).Inc()
}
