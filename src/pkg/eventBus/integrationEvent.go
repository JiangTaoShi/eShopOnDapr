package eventBus

import (
	"time"
)

type IntegrationEvent struct {
	Id           string
	CreationDate time.Time
}
