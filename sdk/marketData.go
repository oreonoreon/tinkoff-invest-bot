package sdk

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/loggy"
	"tinkoff-invest-bot/metrics"
)

type MarketDataInterface interface {
	// The method of requesting historical candlesticks by instrument.
	GetCandles(figi string, from, to *timestamp.Timestamp, interval t.CandleInterval) ([]*t.HistoricCandle, error)
	// The method of requesting the latest prices for instruments.
	GetLastPrices(figi []string) ([]*t.LastPrice, error)
	// The method of obtaining a glass by instrument.
	GetOrderBook(figi string, depth int) (*t.GetOrderBookResponse, error)
	// The method of requesting the status of trading on instruments.
	GetTradingStatus(figi string) (*t.GetTradingStatusResponse, error)
	// The method of requesting the latest depersonalized transactions on the instrument.
	GetLastTrades(figi string, from, to *timestamp.Timestamp) ([]*t.Trade, error)
}

type MarketDataService struct {
	client t.MarketDataServiceClient
}

func NewMarketDataService() *MarketDataService {
	conn, err := Connection()
	if err != nil {
		loggy.GetLogger().Sugar().Fatal(err.Error())
	}

	client := t.NewMarketDataServiceClient(conn)
	return &MarketDataService{client: client}
}

func (mds MarketDataService) GetCandles(figi string, from, to *timestamp.Timestamp, interval t.CandleInterval) ([]*t.HistoricCandle, error) {
	ctx, cancel := NewContext()
	defer cancel()

	mds.incrementRequestsCounter("GetCandles")
	res, err := mds.client.GetCandles(ctx, &t.GetCandlesRequest{
		Figi:     figi,
		From:     from,
		To:       to,
		Interval: interval,
	})
	if err != nil {
		mds.incrementApiCallErrors("GetCandles", err.Error())
		return nil, err
	}

	return res.Candles, nil
}

func (mds MarketDataService) GetLastPrices(figi []string) ([]*t.LastPrice, error) {
	ctx, cancel := NewContext()
	defer cancel()

	mds.incrementRequestsCounter("GetLastPrices")
	res, err := mds.client.GetLastPrices(ctx, &t.GetLastPricesRequest{
		Figi: figi,
	})
	if err != nil {
		mds.incrementApiCallErrors("GetLastPrices", err.Error())
		return nil, err
	}

	return res.LastPrices, nil
}

func (mds MarketDataService) GetOrderBook(figi string, depth int) (*t.GetOrderBookResponse, error) {
	ctx, cancel := NewContext()
	defer cancel()

	mds.incrementRequestsCounter("GetOrderBook")
	res, err := mds.client.GetOrderBook(ctx, &t.GetOrderBookRequest{
		Figi:  figi,
		Depth: int32(depth),
	})
	if err != nil {
		mds.incrementApiCallErrors("GetOrderBook", err.Error())
		return nil, err
	}

	return res, nil
}

func (mds MarketDataService) GetTradingStatus(figi string) (*t.GetTradingStatusResponse, error) {
	ctx, cancel := NewContext()
	defer cancel()

	mds.incrementRequestsCounter("GetTradingStatus")
	res, err := mds.client.GetTradingStatus(ctx, &t.GetTradingStatusRequest{
		Figi: figi,
	})
	if err != nil {
		mds.incrementApiCallErrors("GetTradingStatus", err.Error())
		return nil, err
	}

	return res, nil
}

func (mds MarketDataService) GetLastTrades(figi string, from, to *timestamp.Timestamp) ([]*t.Trade, error) {
	ctx, cancel := NewContext()
	defer cancel()

	mds.incrementRequestsCounter("GetLastTrades")
	res, err := mds.client.GetLastTrades(ctx, &t.GetLastTradesRequest{
		Figi: figi,
		From: from,
		To:   to,
	})
	if err != nil {
		mds.incrementApiCallErrors("GetLastTrades", err.Error())
		return nil, err
	}

	return res.Trades, nil
}

func (mds MarketDataService) incrementRequestsCounter(method string) {
	metrics.ApiRequests.WithLabelValues(loggy.GetBotID(), "MarketDataService", method).Inc()
}

func (mds MarketDataService) incrementApiCallErrors(method string, error string) {
	metrics.ApiCallErrors.WithLabelValues(loggy.GetBotID(), "MarketDataService", method, error).Inc()
}
