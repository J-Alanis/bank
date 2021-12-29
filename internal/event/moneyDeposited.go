package event

import (
	"time"

	"github.com/google/uuid"
)

type moneyDepositedData struct {
	amount int
}

func moneyDeposited(aggregateID uuid.UUID, aggregateVersion int, correlationID uuid.UUID, amount int) Event {

	data := moneyDepositedData{
		amount: amount,
	}

	return Event{
		Id:               uuid.NewUUID(),
		AggregateID:      aggregateID,
		aggregateName:    AggregateName,
		correlationID:    correlationID,
		aggregateVersion: aggregateVersion + 1,
		eType:            "moneyDeposited",
		data:             data,
		createdAt:        time.Now(),
	}
}
