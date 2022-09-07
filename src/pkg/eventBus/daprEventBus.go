package eventBus

import (
	"context"
	dapr "github.com/dapr/go-sdk/client"
)

type DaprEventBus struct {
	IEventBus
}

const DAPR_PUBSUB_NAME string = "pubsub"

func (*DaprEventBus) Publish(integrationEvent *IntegrationEvent) {
	ctx := context.Background()
	client, _ := dapr.NewClient()
	client.PublishEvent(ctx, DAPR_PUBSUB_NAME, "topic-name", nil)
}
