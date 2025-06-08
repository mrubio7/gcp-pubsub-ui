package pubsub

import (
	"context"
	"sync"

	"google.golang.org/api/option"
	"google.golang.org/api/pubsub/v1"
)

type client struct {
	svc *pubsub.Service
}

var (
	instance *client
	once     sync.Once
	initErr  error
)

func getClient(ctx context.Context) (*client, error) {
	once.Do(func() {
		svc, err := pubsub.NewService(ctx, option.WithCredentialsFile("service_account.json"))
		if err != nil {
			initErr = err
			return
		}
		instance = &client{svc: svc}
	})

	return instance, initErr
}
