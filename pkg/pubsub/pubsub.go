package pubsub

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/pubsub/v1"
)

type Client struct {
	svc *pubsub.Service
}

func NewClient(ctx context.Context, filepath string) (*Client, error) {
	s, err := pubsub.NewService(ctx, option.WithCredentialsFile(filepath))
	if err != nil {
		return nil, err
	}

	return &Client{svc: s}, nil
}
