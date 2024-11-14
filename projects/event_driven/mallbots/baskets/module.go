package baskets

import (
	"context"
	"mallbots/internal/modules"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mainModule modules.MainModule) (err error) {
	// setup Driven adapters
	// domainDispatcher := ddd.NewEventDispatcher()
	// baskets := repository.NewBasketRepository("baskets.baskets", mainModule.DB())
	// client, err := grpc.NewClient(ctx, mainModule.Logger(), mainModule.Config().RPCConfig.Address())
	// if err != nil {
	// 	return err
	// }
	// stores := grpc.NewStoreRepository(client)
	// products := grpc.NewProductRepository(client)
	// orders := grpc.NewOrderRepository(client)
	return
}
