package event

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id               uuid.UUID
	AggregateID      uuid.UUID
	aggregateName    string
	correlationID    uuid.UUID
	aggregateVersion int
	eType            string
	data             interface{}
	createdAt        time.Time
}
