package grpc

import (
	"context"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(ctx context.Context, logger zerolog.Logger, endpoint string) (conn *grpc.ClientConn, err error) {
	conn, err = grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	// handle connection close
	defer func() {
		if err != nil {
			logger.Error().Err(err).Msg("basket: could not initialize the grpc client")
			if err = conn.Close(); err != nil {
				logger.Error().Err(err).Msg("basket: could not close the grpc client connection")
			}
			return
		}
		go func() {
			<-ctx.Done()
			if err = conn.Close(); err != nil {
				logger.Error().Err(err).Msg("basket: could not close the grpc client connection")
			}
		}()
	}()

	return
}
