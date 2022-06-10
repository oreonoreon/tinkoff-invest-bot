package sdk

import (
	t "tinkoff-invest-bot/Tinkoff/investapi"
	"tinkoff-invest-bot/loggy"
)

type OrdersStreamInterface interface {
	// Recv listens for incoming messages and block until first one is received.
	Recv() (*t.TradesStreamResponse, error)
}

type OrdersStream struct {
	client t.OrdersStreamServiceClient
	stream t.OrdersStreamService_TradesStreamClient
}

func NewOrdersStream(request *t.TradesStreamRequest) *OrdersStream {
	conn, err := Connection()
	if err != nil {
		loggy.GetLogger().Sugar().Fatal(err.Error())
	}

	client := t.NewOrdersStreamServiceClient(conn)
	ctx, cancel := NewContext()
	defer cancel()

	stream, err := client.TradesStream(ctx, request)
	if err != nil {
		loggy.GetLogger().Sugar().Fatal(err.Error())
	}

	return &OrdersStream{client: client, stream: stream}
}

func (os OrdersStream) Recv() (*t.TradesStreamResponse, error) {
	return os.stream.Recv()
}
