package sdk

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/internal/loggy"
	"tinkoff-invest-bot/internal/metrics"
)

type OperationsInterface interface {
	// Method for getting a list of account transactions.
	GetOperations(accountID string, from, to *timestamp.Timestamp, state t.OperationState, figi string) ([]*t.Operation, error)
	// The method of obtaining a portfolio by account.
	GetPortfolio(accountID string) (*t.PortfolioResponse, error)
	// Method for getting a list of account positions.
	GetPositions(accountID string) (*t.PositionsResponse, error)
	// The method of obtaining the available balance for withdrawal of funds.
	GetWithdrawLimits(accountID string) (*t.WithdrawLimitsResponse, error)
}

type OperationsService struct {
	client t.OperationsServiceClient
}

func NewOperationsService() *OperationsService {
	conn, err := Connection()
	if err != nil {
		loggy.GetLogger().Sugar().Fatal(err.Error())
	}

	client := t.NewOperationsServiceClient(conn)
	return &OperationsService{client: client}
}

func (os OperationsService) GetOperations(accountID string, from, to *timestamp.Timestamp, state t.OperationState, figi string) ([]*t.Operation, error) {
	ctx, cancel := NewContext()
	defer cancel()

	os.incrementRequestsCounter("GetOperations")
	res, err := os.client.GetOperations(ctx, &t.OperationsRequest{
		AccountId: accountID,
		From:      from,
		To:        to,
		State:     state,
		Figi:      figi,
	})
	if err != nil {
		os.incrementApiCallErrors("GetOperations", err.Error())
		return nil, err
	}

	return res.Operations, nil
}

func (os OperationsService) GetPortfolio(accountID string) (*t.PortfolioResponse, error) {
	ctx, cancel := NewContext()
	defer cancel()

	os.incrementRequestsCounter("GetPortfolio")
	res, err := os.client.GetPortfolio(ctx, &t.PortfolioRequest{
		AccountId: accountID,
	})
	if err != nil {
		os.incrementApiCallErrors("GetPortfolio", err.Error())
		return nil, err
	}

	return res, nil
}

func (os OperationsService) GetPositions(accountID string) (*t.PositionsResponse, error) {
	ctx, cancel := NewContext()
	defer cancel()

	os.incrementRequestsCounter("GetPositions")
	res, err := os.client.GetPositions(ctx, &t.PositionsRequest{
		AccountId: accountID,
	})
	if err != nil {
		os.incrementApiCallErrors("GetPositions", err.Error())
		return nil, err
	}

	return res, nil
}

func (os OperationsService) GetWithdrawLimits(accountID string) (*t.WithdrawLimitsResponse, error) {
	ctx, cancel := NewContext()
	defer cancel()

	os.incrementRequestsCounter("GetWithdrawLimits")
	res, err := os.client.GetWithdrawLimits(ctx, &t.WithdrawLimitsRequest{
		AccountId: accountID,
	})
	if err != nil {
		os.incrementApiCallErrors("GetWithdrawLimits", err.Error())
		return nil, err
	}

	return res, nil
}

func (os OperationsService) incrementRequestsCounter(method string) {
	metrics.ApiRequests.WithLabelValues(loggy.GetBotID(), "OperationsService", method).Inc()
}

func (os OperationsService) incrementApiCallErrors(method string, error string) {
	metrics.ApiCallErrors.WithLabelValues(loggy.GetBotID(), "OperationsService", method, error).Inc()
}
