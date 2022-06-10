package sdk

type Services struct {
	InstrumentsService      InstrumentsService
	MarketDataService       MarketDataService
	MarketDataServiceStream MarketDataStream
	OperationsService       OperationsService
	OrdersService           OrdersService
	SandboxService          SandboxService
	StopOrdersService       StopOrdersService
	UsersService            UsersService
}

func NewServices() *Services {
	return &Services{
		InstrumentsService:      *NewInstrumentsService(),
		MarketDataService:       *NewMarketDataService(),
		MarketDataServiceStream: *NewMarketDataStream(),
		OperationsService:       *NewOperationsService(),
		OrdersService:           *NewOrdersService(),
		StopOrdersService:       *NewStopOrdersService(),
		SandboxService:          *NewSandboxService(),
		UsersService:            *NewUsersService(),
	}
}
