package eventBus

type IEventBus interface {
	Publish(integrationEvent *IntegrationEvent)
}
