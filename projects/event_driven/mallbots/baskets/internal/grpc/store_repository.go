package grpc

// var _ domain.StoreRepository = (*StoreRepository)(nil)

// func NewStoreRepository(conn *grpc.ClientConn) StoreRepository {
// 	return StoreRepository{client: storespb.NewStoresServiceClient(conn)}
// }

// type StoreRepository struct {
// 	client storespb.StoresServiceClient
// }

// func (r *StoreRepository) Find(ctx context.Context, storeID string) (*domain.Store, error) {
// 	resp, err := r.client.GetStore(ctx, &storespb.GetStoreRequest{Id: storeID})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return r.storeToDomain(resp.Store), nil
// }

// func (r StoreRepository) storeToDomain(store *storespb.Store) *domain.Store {
// 	return &domain.Store{
// 		ID:       store.GetId(),
// 		Name:     store.GetName(),
// 		Location: store.GetLocation(),
// 	}
// }
