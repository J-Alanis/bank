package event

import (
	"time"

	"github.com/google/uuid"
)

const (
	AggregateName = "account"
)

func NewAccountCreated() Event {

	return Event{
		Id:               uuid.NewUUID(),
		AggregateID:      uuid.NewUUID(),
		aggregateName:    AggregateName,
		correlationID:    nil,
		aggregateVersion: 1,
		eType:            "accountCreated",
		data:             nil,
		createdAt:        time.Now(),
	}
}
