package sdk

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"log"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/internal/loggy"
	"tinkoff-invest-bot/internal/metrics"
)

type Instruments interface {
	// TradingSchedules The method of obtaining the trading schedule of trading platforms.
	TradingSchedules(exchange string, from, to *timestamp.Timestamp) ([]*t.TradingSchedule, error)
	// BondBy The method of obtaining a bond by its identifier.
	BondBy(filters t.InstrumentRequest) (*t.Bond, error)
	// Method of obtaining a list of bonds.
	Bonds(status t.InstrumentStatus) ([]*t.Bond, error)
	// Method of obtaining a coupon payment schedule for a bond.
	GetBondCoupons(figi string, from, to *timestamp.Timestamp) ([]*t.Coupon, error)
	// The method of obtaining a currency by its identifier.
	CurrencyBy(filters t.InstrumentRequest) (*t.Currency, error)
	// Method for getting a list of currencies.
	Currencies(status t.InstrumentStatus) ([]*t.Currency, error)
	// The method of obtaining an investment fund by its identifier.
	EtfBy(filters t.InstrumentRequest) (*t.Etf, error)
	// Method of obtaining a list of investment funds.
	Etfs(status t.InstrumentStatus) ([]*t.Etf, error)
	// The method of obtaining futures by its identifier.
	FutureBy(filters t.InstrumentRequest) (*t.Future, error)
	// Method for getting a list of futures.
	Futures(status t.InstrumentStatus) ([]*t.Future, error)
	// The method of obtaining a stock by its identifier.
	ShareBy(filters t.InstrumentRequest) (*t.Share, error)
	// Method of getting a list of shares.
	Shares(status t.InstrumentStatus) ([]*t.Share, error)
	// The method of obtaining the accumulated coupon income on the bond.
	GetAccruedInterests(figi string, from, to *timestamp.Timestamp) ([]*t.AccruedInterest, error)
	// The method of obtaining the amount of the guarantee for futures.
	GetFuturesMargin(figi string) (*t.GetFuturesMarginResponse, error)
	// The method of obtaining basic information about the tool.
	GetInstrumentBy(filters t.InstrumentRequest) (*t.Instrument, error)
	// A method for obtaining dividend payment events for an instrument.
	GetDividends(figi string, from, to *timestamp.Timestamp) ([]*t.Dividend, error)
	// The method of obtaining an asset by its identifier.
	GetAssetBy(assetID string) (*t.AssetFull, error)
	// Method for getting a list of assets.
	GetAssets() ([]*t.Asset, error)
	// The method of getting the favourite instruments.
	GetFavorites() ([]*t.FavoriteInstrument, error)
	// The method of editing the selected instruments.
	EditFavorites(newFavourites *t.EditFavoritesRequest) ([]*t.FavoriteInstrument, error)
}

type InstrumentsService struct {
	client t.InstrumentsServiceClient
}

func NewInstrumentsService() *InstrumentsService {
	conn, err := Connection()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := t.NewInstrumentsServiceClient(conn)
	return &InstrumentsService{client: client}
}

func (i InstrumentsService) TradingSchedules(exchange string, from, to *timestamp.Timestamp) ([]*t.TradingSchedule, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("TradingSchedules")
	res, err := i.client.TradingSchedules(ctx, &t.TradingSchedulesRequest{
		Exchange: exchange,
		From:     from,
		To:       to,
	})
	if err != nil {
		i.incrementApiCallErrors("TradingSchedules", err.Error())
		return nil, err
	}

	return res.Exchanges, nil
}

func (i InstrumentsService) BondBy(filters t.InstrumentRequest) (*t.Bond, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("BondBy")
	res, err := i.client.BondBy(ctx, &filters)
	if err != nil {
		i.incrementApiCallErrors("BondBy", err.Error())
		return nil, err
	}

	return res.Instrument, nil
}

func (i InstrumentsService) Bonds(status t.InstrumentStatus) ([]*t.Bond, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("Bonds")
	res, err := i.client.Bonds(ctx, &t.InstrumentsRequest{
		InstrumentStatus: status,
	})
	if err != nil {
		i.incrementApiCallErrors("Bonds", err.Error())
		return nil, err
	}

	return res.Instruments, nil
}

func (i InstrumentsService) GetBondCoupons(figi string, from, to *timestamp.Timestamp) ([]*t.Coupon, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("GetBoundCoupons")
	res, err := i.client.GetBondCoupons(ctx, &t.GetBondCouponsRequest{
		Figi: figi,
		From: from,
		To:   to,
	})
	if err != nil {
		i.incrementApiCallErrors("GetBoundCoupons", err.Error())
		return nil, err
	}

	return res.Events, nil
}

func (i InstrumentsService) CurrencyBy(filters t.InstrumentRequest) (*t.Currency, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("CurrencyBy")
	res, err := i.client.CurrencyBy(ctx, &filters)
	if err != nil {
		i.incrementApiCallErrors("CurrencyBy", err.Error())
		return nil, err
	}

	return res.Instrument, nil
}

func (i InstrumentsService) Currencies(status t.InstrumentStatus) ([]*t.Currency, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("Currencies")
	res, err := i.client.Currencies(ctx, &t.InstrumentsRequest{
		InstrumentStatus: status,
	})
	if err != nil {
		i.incrementApiCallErrors("Currencies", err.Error())
		return nil, err
	}

	return res.Instruments, nil
}

func (i InstrumentsService) EtfBy(filters t.InstrumentRequest) (*t.Etf, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("EtfBy")
	res, err := i.client.EtfBy(ctx, &filters)
	if err != nil {
		i.incrementApiCallErrors("EtfBy", err.Error())
		return nil, err
	}

	return res.Instrument, nil
}

func (i InstrumentsService) Etfs(status t.InstrumentStatus) ([]*t.Etf, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("Etfs")
	res, err := i.client.Etfs(ctx, &t.InstrumentsRequest{
		InstrumentStatus: status,
	})
	if err != nil {
		i.incrementApiCallErrors("Etfs", err.Error())
		return nil, err
	}

	return res.Instruments, nil
}

func (i InstrumentsService) FutureBy(filters t.InstrumentRequest) (*t.Future, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("FutureBy")
	res, err := i.client.FutureBy(ctx, &filters)
	if err != nil {
		i.incrementApiCallErrors("FutureBy", err.Error())
		return nil, err
	}

	return res.Instrument, nil
}

func (i InstrumentsService) Futures(status t.InstrumentStatus) ([]*t.Future, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("Futures")
	res, err := i.client.Futures(ctx, &t.InstrumentsRequest{
		InstrumentStatus: status,
	})
	if err != nil {
		i.incrementApiCallErrors("Futures", err.Error())
		return nil, err
	}

	return res.Instruments, nil
}

func (i InstrumentsService) ShareBy(filters t.InstrumentRequest) (*t.Share, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("ShareBy")
	res, err := i.client.ShareBy(ctx, &filters)
	if err != nil {
		i.incrementApiCallErrors("ShareBy", err.Error())
		return nil, err
	}

	return res.Instrument, nil
}

func (i InstrumentsService) Shares(status t.InstrumentStatus) ([]*t.Share, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("Shares")
	res, err := i.client.Shares(ctx, &t.InstrumentsRequest{
		InstrumentStatus: status,
	})
	if err != nil {
		i.incrementApiCallErrors("Shares", err.Error())
		return nil, err
	}

	return res.Instruments, nil
}

func (i InstrumentsService) GetAccruedInterests(figi string, from, to *timestamp.Timestamp) ([]*t.AccruedInterest, error) {
	ctx, cancel := NewContext()
	defer cancel() //ctx, cancel :=createRequestContext() // defer cancel()

	i.incrementRequestsCounter("GetAccruedInterests")
	res, err := i.client.GetAccruedInterests(ctx, &t.GetAccruedInterestsRequest{
		Figi: figi,
		From: from,
		To:   to,
	})
	if err != nil {
		i.incrementApiCallErrors("GetAccruedInterests", err.Error())
		return nil, err
	}

	return res.AccruedInterests, nil
}

func (i InstrumentsService) GetFuturesMargin(figi string) (*t.GetFuturesMarginResponse, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("GetFuturesMargin")
	res, err := i.client.GetFuturesMargin(ctx, &t.GetFuturesMarginRequest{
		Figi: figi,
	})
	if err != nil {
		i.incrementApiCallErrors("GetFuturesMargin", err.Error())
		return nil, err
	}

	return res, nil
}

func (i InstrumentsService) GetInstrumentBy(filters t.InstrumentRequest) (*t.Instrument, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("GetInstrumentBy")
	res, err := i.client.GetInstrumentBy(ctx, &filters)
	if err != nil {
		i.incrementApiCallErrors("GetInstrumentBy", err.Error())
		return nil, err
	}

	return res.Instrument, nil
}

func (i InstrumentsService) GetDividends(figi string, from, to *timestamp.Timestamp) ([]*t.Dividend, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("GetDividends")
	res, err := i.client.GetDividends(ctx, &t.GetDividendsRequest{
		Figi: figi,
		From: from,
		To:   to,
	})
	if err != nil {
		i.incrementApiCallErrors("GetDividends", err.Error())
		return nil, err
	}

	return res.Dividends, nil
}

func (i InstrumentsService) GetAssetBy(assetID string) (*t.AssetFull, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("GetAssetBy")
	res, err := i.client.GetAssetBy(ctx, &t.AssetRequest{
		Id: assetID,
	})
	if err != nil {
		i.incrementApiCallErrors("GetAssetBy", err.Error())
		return nil, err
	}

	return res.Asset, nil
}

func (i InstrumentsService) GetAssets() ([]*t.Asset, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("GetAssets")
	res, err := i.client.GetAssets(ctx, &t.AssetsRequest{})
	if err != nil {
		i.incrementApiCallErrors("GetAssets", err.Error())
		return nil, err
	}

	return res.Assets, nil
}

func (i InstrumentsService) GetFavorites() ([]*t.FavoriteInstrument, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("GetFavourites")
	res, err := i.client.GetFavorites(ctx, &t.GetFavoritesRequest{})
	if err != nil {
		i.incrementApiCallErrors("GetFavourites", err.Error())
		return nil, err
	}

	return res.FavoriteInstruments, nil
}

func (i InstrumentsService) EditFavorites(newFavourites *t.EditFavoritesRequest) ([]*t.FavoriteInstrument, error) {
	ctx, cancel := NewContext()
	defer cancel()

	i.incrementRequestsCounter("EditFavorites")
	res, err := i.client.EditFavorites(ctx, newFavourites)
	if err != nil {
		i.incrementApiCallErrors("EditFavorites", err.Error())
		return nil, err
	}

	return res.FavoriteInstruments, nil
}

func (i InstrumentsService) incrementRequestsCounter(method string) {
	metrics.ApiRequests.WithLabelValues(loggy.GetBotID(), "InstrumentsService", method).Inc()
}

func (i InstrumentsService) incrementApiCallErrors(method string, error string) {
	metrics.ApiCallErrors.WithLabelValues(loggy.GetBotID(), "InstrumentsService", method, error).Inc()
}
