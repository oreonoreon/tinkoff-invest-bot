package sdk

import (
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/internal/loggy"
	"tinkoff-invest-bot/internal/metrics"
)

type UsersServiceClient interface {
	// The method of receiving user accounts.
	GetAccounts() ([]*t.Account, error)
	// Calculation of margin indicators on the account.
	GetMarginAttributes(accountID string) (*t.GetMarginAttributesResponse, error)
	// Request for the user's tariff.
	GetUserTariff() (*t.GetUserTariffResponse, error)
	// The method of obtaining information about the user.
	GetInfo() (*t.GetInfoResponse, error)
}

type UsersService struct {
	client t.UsersServiceClient
}

func NewUsersService() *UsersService {
	conn, err := Connection()
	if err != nil {
		loggy.GetLogger().Sugar().Fatal(err.Error())
	}

	client := t.NewUsersServiceClient(conn)
	return &UsersService{client: client}
}

func (us UsersService) GetAccounts() ([]*t.Account, error) {
	ctx, cancel := NewContext()
	defer cancel()
	us.incrementRequestsCounter("GetAccounts")
	res, err := us.client.GetAccounts(ctx, &t.GetAccountsRequest{})
	if err != nil {
		us.incrementApiCallErrors("GetAccounts", err.Error())
		return nil, err
	}

	return res.Accounts, nil
}

func (us UsersService) GetMarginAttributes(accountID string) (*t.GetMarginAttributesResponse, error) {
	ctx, cancel := NewContext()
	defer cancel()

	us.incrementRequestsCounter("GetMarginAttributes")
	res, err := us.client.GetMarginAttributes(ctx, &t.GetMarginAttributesRequest{
		AccountId: accountID,
	})
	if err != nil {
		us.incrementApiCallErrors("GetMarginAttributes", err.Error())
		return nil, err
	}

	return res, nil
}

func (us UsersService) GetUserTariff() (*t.GetUserTariffResponse, error) {
	ctx, cancel := NewContext()
	defer cancel()
	us.incrementRequestsCounter("GetUserTariff")
	res, err := us.client.GetUserTariff(ctx, &t.GetUserTariffRequest{})
	if err != nil {
		us.incrementApiCallErrors("GetUserTariff", err.Error())
		return nil, err
	}

	return res, nil
}

func (us UsersService) GetInfo() (*t.GetInfoResponse, error) {
	ctx, cancel := NewContext()
	defer cancel()

	us.incrementRequestsCounter("GetInfo")
	res, err := us.client.GetInfo(ctx, &t.GetInfoRequest{})
	if err != nil {
		us.incrementApiCallErrors("GetInfo", err.Error())
		return nil, err
	}

	return res, nil
}

func (us UsersService) incrementRequestsCounter(method string) {
	metrics.ApiRequests.WithLabelValues(loggy.GetBotID(), "UsersService", method).Inc()
}

func (us UsersService) incrementApiCallErrors(method string, error string) {
	metrics.ApiCallErrors.WithLabelValues(loggy.GetBotID(), "UsersService", method, error).Inc()
}
