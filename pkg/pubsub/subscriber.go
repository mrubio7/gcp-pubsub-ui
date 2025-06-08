package pubsub

import (
	"context"
	"fmt"
	"pubsub-ui/pkg/config"

	"google.golang.org/api/pubsub/v1"
)

type Subscriber struct {
	client      *pubsub.ProjectsSubscriptionsService
	topic       string
	maxMessages int
}

func NewSubscriber(ctx context.Context, topic string) *Subscriber {
	c, err := getClient(ctx)
	if err != nil {
		return nil
	}
	pubsubService := pubsub.NewProjectsSubscriptionsService(c.svc)

	return &Subscriber{
		client:      pubsubService,
		topic:       fmt.Sprintf("%s/subscriptions/%s", config.AppConfig.GCP_BASE_URL, topic),
		maxMessages: config.AppConfig.SUSCRIPTION_MAX_MESSAGES,
	}
}
