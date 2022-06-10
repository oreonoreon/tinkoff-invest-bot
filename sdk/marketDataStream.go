package sdk

import (
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/loggy"
)

type MarketDataStreamInterface interface {
	// Recv listens for incoming messages and block until first one is received.
	Recv() (*t.MarketDataResponse, error)
	// Send puts t.MarketDataRequest into a stream.
	Send(request *t.MarketDataRequest) error
}

type MarketDataStream struct {
	client t.MarketDataStreamServiceClient
	stream t.MarketDataStreamService_MarketDataStreamClient
}

func NewMarketDataStream() *MarketDataStream {
	conn, err := Connection()
	if err != nil {
		loggy.GetLogger().Sugar().Fatal(err.Error())
	}

	client := t.NewMarketDataStreamServiceClient(conn)
	ctx := NewContextStream()
	//defer cancel()

	stream, err := client.MarketDataStream(ctx)
	if err != nil {
		loggy.GetLogger().Sugar().Fatal(err.Error())
	}

	return &MarketDataStream{client: client, stream: stream}
}

func (m MarketDataStream) Recv() (*t.MarketDataResponse, error) {
	return m.stream.Recv()
}

func (m MarketDataStream) Send(request *t.MarketDataRequest) error {
	return m.stream.Send(request)
}
