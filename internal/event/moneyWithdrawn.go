package event

import (
	"time"

	"github.com/google/uuid"
)

type moneyWithdrawnData struct {
	amount int
}

func moneyWithdrawn(aggregateID uuid.UUID, amount int) Event {

	data := moneyWithdrawnData{
		amount: amount,
	}

	return Event{
		Id:               uuid.NewUUID(),
		AggregateID:      aggregateID,
		aggregateName:    AggregateName,
		correlationID:    nil,
		aggregateVersion: 1,
		eType:            "moneyWithdrawn",
		data:             data,
		createdAt:        time.Now(),
	}
}
